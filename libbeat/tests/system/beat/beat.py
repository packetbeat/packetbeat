import subprocess

import jinja2
import unittest
import os
import shutil
import json
import signal
import sys
import time
import yaml
from datetime import datetime, timedelta

from .compose import ComposeMixin


BEAT_REQUIRED_FIELDS = ["@timestamp",
                        "beat.name", "beat.hostname", "beat.version"]

INTEGRATION_TESTS = os.environ.get('INTEGRATION_TESTS', False)


class TimeoutError(Exception):
    pass


class Proc(object):
    """
    Slim wrapper on subprocess.Popen that redirects
    both stdout and stderr to a file on disk and makes
    sure to stop the process and close the output file when
    the object gets collected.
    """

    def __init__(self, args, outputfile):
        self.args = args
        self.output = open(outputfile, "ab")
        self.stdin_read, self.stdin_write = os.pipe()

    def start(self):

        if sys.platform.startswith("win"):
            self.proc = subprocess.Popen(
                self.args,
                stdin=self.stdin_read,
                stdout=self.output,
                stderr=subprocess.STDOUT,
                bufsize=0,
                creationflags=subprocess.CREATE_NEW_PROCESS_GROUP)
        else:
            self.proc = subprocess.Popen(
                self.args,
                stdin=self.stdin_read,
                stdout=self.output,
                stderr=subprocess.STDOUT,
                bufsize=0,
            )
        return self.proc

    def kill(self):
        if sys.platform.startswith("win"):
            # proc.terminate on Windows does not initiate a graceful shutdown
            # through the processes signal handlers it just kills it hard. So
            # this sends a SIGBREAK. You cannot sends a SIGINT (CTRL_C_EVENT)
            # to a process group in Windows, otherwise Ctrl+C would be
            # sent.
            self.proc.send_signal(signal.CTRL_BREAK_EVENT)
        else:
            self.proc.terminate()

    def wait(self):
        try:
            return self.proc.wait()
        finally:
            self.output.close()

    def check_wait(self, exit_code=0):
        actual_exit_code = self.wait()
        assert actual_exit_code == exit_code, "Expected exit code to be %d, but it was %d" % (
            exit_code, actual_exit_code)
        return actual_exit_code

    def kill_and_wait(self):
        self.kill()
        os.close(self.stdin_write)
        return self.wait()

    def check_kill_and_wait(self, exit_code=0):
        self.kill()
        os.close(self.stdin_write)
        return self.check_wait(exit_code=exit_code)

    def __del__(self):
        # Ensure the process is stopped.
        try:
            self.proc.terminate()
            self.proc.kill()
        except:
            pass
        # Ensure the output is closed.
        try:
            self.output.close()
        except:
            pass


class TestCase(unittest.TestCase, ComposeMixin):

    @classmethod
    def setUpClass(self):

        # Path to test binary
        if not hasattr(self, 'beat_name'):
            self.beat_name = "beat"

        if not hasattr(self, 'beat_path'):
            self.beat_path = "."

        # Path to test binary
        if not hasattr(self, 'test_binary'):
            self.test_binary = os.path.abspath(self.beat_path + "/" + self.beat_name + ".test")

        # Create build path
        build_dir = self.beat_path + "/build"
        self.build_path = build_dir + "/system-tests/"

        # Start the containers needed to run these tests
        self.compose_up()

    @classmethod
    def tearDownClass(self):
        self.compose_down()

    def run_beat(self,
                 cmd=None,
                 config=None,
                 output=None,
                 logging_args=["-e", "-v", "-d", "*"],
                 extra_args=[],
                 exit_code=None):
        """
        Executes beat.
        Waits for the process to finish before returning to
        the caller.
        """
        proc = self.start_beat(cmd=cmd, config=config, output=output,
                               logging_args=logging_args,
                               extra_args=extra_args)
        if exit_code != None:
            return proc.check_wait(exit_code)

        return proc.wait()

    def start_beat(self,
                   cmd=None,
                   config=None,
                   output=None,
                   logging_args=["-e", "-v", "-d", "*"],
                   extra_args=[]):
        """
        Starts beat and returns the process handle. The
        caller is responsible for stopping / waiting for the
        Proc instance.
        """

        # Init defaults
        if cmd is None:
            cmd = self.test_binary

        if config is None:
            config = self.beat_name + ".yml"

        if output is None:
            output = self.beat_name + ".log"

        args = [cmd,
                "-systemTest",
                "-test.coverprofile",
                os.path.join(self.working_dir, "coverage.cov"),
                "-path.home", os.path.normpath(self.working_dir),
                "-c", os.path.join(self.working_dir, config),
                ]

        if logging_args:
            args.extend(logging_args)

        if extra_args:
            args.extend(extra_args)

        proc = Proc(args, os.path.join(self.working_dir, output))
        proc.start()
        return proc

    def render_config_template(self, template_name=None,
                               output=None, **kargs):

        # Init defaults
        if template_name is None:
            template_name = self.beat_name

        template_path = "./tests/system/config/" + template_name + ".yml.j2"

        if output is None:
            output = self.beat_name + ".yml"

        template = self.template_env.get_template(template_path)

        kargs["beat"] = self
        output_str = template.render(**kargs)

        output_path = os.path.join(self.working_dir, output)
        with open(output_path, "wb") as f:
            os.chmod(output_path, 0600)
            f.write(output_str.encode('utf8'))

    # Returns output as JSON object with flattened fields (. notation)
    def read_output(self,
                    output_file=None,
                    required_fields=None):

        # Init defaults
        if output_file is None:
            output_file = "output/" + self.beat_name

        jsons = []
        with open(os.path.join(self.working_dir, output_file), "r") as f:
            for line in f:
                if len(line) == 0 or line[len(line) - 1] != "\n":
                    # hit EOF
                    break

                try:
                    jsons.append(self.flatten_object(json.loads(
                        line, object_pairs_hook=self.json_raise_on_duplicates), []))
                except:
                    print("Fail to load the json {}".format(line))
                    raise

        self.all_have_fields(jsons, required_fields or BEAT_REQUIRED_FIELDS)
        return jsons

    # Returns output as JSON object
    def read_output_json(self, output_file=None):

        # Init defaults
        if output_file is None:
            output_file = "output/" + self.beat_name

        jsons = []
        with open(os.path.join(self.working_dir, output_file), "r") as f:
            for line in f:
                if len(line) == 0 or line[len(line) - 1] != "\n":
                    # hit EOF
                    break

                event = json.loads(line, object_pairs_hook=self.json_raise_on_duplicates)
                del event['@metadata']
                jsons.append(event)
        return jsons

    def json_raise_on_duplicates(self, ordered_pairs):
        """Reject duplicate keys. To be used as a custom hook in JSON unmarshaling
           to error out in case of any duplicates in the keys."""
        d = {}
        for k, v in ordered_pairs:
            if k in d:
                raise ValueError("duplicate key: %r" % (k,))
            else:
                d[k] = v
        return d

    def copy_files(self, files, source_dir="files/"):
        for file_ in files:
            shutil.copy(os.path.join(source_dir, file_),
                        self.working_dir)

    def setUp(self):

        self.template_env = jinja2.Environment(
            loader=jinja2.FileSystemLoader(self.beat_path)
        )

        # create working dir
        self.working_dir = os.path.abspath(os.path.join(
            self.build_path + "run", self.id()))
        if os.path.exists(self.working_dir):
            shutil.rmtree(self.working_dir)
        os.makedirs(self.working_dir)

        fields_yml = os.path.join(self.beat_path, "fields.yml")
        # Only add it if it exists
        if os.path.isfile(fields_yml):
            shutil.copyfile(fields_yml, os.path.join(self.working_dir, "fields.yml"))

        try:
            # update the last_run link
            if os.path.islink(self.build_path + "last_run"):
                os.unlink(self.build_path + "last_run")
            os.symlink(self.build_path + "run/{}".format(self.id()),
                       self.build_path + "last_run")
        except:
            # symlink is best effort and can fail when
            # running tests in parallel
            pass

    def wait_until(self, cond, max_timeout=10, poll_interval=0.1, name="cond"):
        """
        Waits until the cond function returns true,
        or until the max_timeout is reached. Calls the cond
        function every poll_interval seconds.

        If the max_timeout is reached before cond() returns
        true, an exception is raised.
        """
        start = datetime.now()
        while not cond():
            if datetime.now() - start > timedelta(seconds=max_timeout):
                raise TimeoutError("Timeout waiting for '{}' to be true. ".format(name) +
                                   "Waited {} seconds.".format(max_timeout))
            time.sleep(poll_interval)

    def get_log(self, logfile=None):
        """
        Returns the log as a string.
        """
        if logfile is None:
            logfile = self.beat_name + ".log"

        with open(os.path.join(self.working_dir, logfile), 'r') as f:
            data = f.read()

        return data

    def wait_log_contains(self, msg, logfile=None,
                          max_timeout=10, poll_interval=0.1,
                          name="log_contains"):
        self.wait_until(
            cond=lambda: self.log_contains(msg, logfile),
            max_timeout=max_timeout,
            poll_interval=poll_interval,
            name=name)

    def log_contains(self, msg, logfile=None):
        """
        Returns true if the give logfile contains the given message.
        Note that the msg must be present in a single line.
        """

        return self.log_contains_count(msg, logfile) > 0

    def log_contains_count(self, msg, logfile=None):
        """
        Returns the number of appearances of the given string in the log file
        """

        counter = 0

        # Init defaults
        if logfile is None:
            logfile = self.beat_name + ".log"

        try:
            with open(os.path.join(self.working_dir, logfile), "r") as f:
                for line in f:
                    if line.find(msg) >= 0:
                        counter = counter + 1
        except IOError:
            counter = -1

        return counter

    def output_lines(self, output_file=None):
        """ Count number of lines in a file."""
        if output_file is None:
            output_file = "output/" + self.beat_name

        try:
            with open(os.path.join(self.working_dir, output_file), "r") as f:
                return sum([1 for line in f])
        except IOError:
            return 0

    def output_has(self, lines, output_file=None):
        """
        Returns true if the output has a given number of lines.
        """

        # Init defaults
        if output_file is None:
            output_file = "output/" + self.beat_name

        try:
            with open(os.path.join(self.working_dir, output_file), "r") as f:
                return len([1 for line in f]) == lines
        except IOError:
            return False

    def output_has_message(self, message, output_file=None):
        """
        Returns true if the output has the given message field.
        """
        try:
            return any(line for line in self.read_output(output_file=output_file, required_fields=["message"])
                       if line.get("message") == message)
        except (IOError, TypeError):
            return False

    def all_have_fields(self, objs, fields):
        """
        Checks that the given list of output objects have
        all the given fields.
        Raises Exception if not true.
        """
        for field in fields:
            if not all([field in o for o in objs]):
                raise Exception("Not all objects have a '{}' field"
                                .format(field))

    def all_have_only_fields(self, objs, fields):
        """
        Checks if the given list of output objects have all
        and only the given fields.
        Raises Exception if not true.
        """
        self.all_have_fields(objs, fields)
        self.all_fields_are_expected(objs, fields)

    def all_fields_are_expected(self, objs, expected_fields,
                                dict_fields=[]):
        """
        Checks that all fields in the objects are from the
        given list of expected fields.
        """
        for o in objs:
            for key in o.keys():
                known = key in dict_fields or key in expected_fields
                ismeta = key.startswith('@metadata.')
                if not(known or ismeta):
                    raise Exception("Unexpected key '{}' found"
                                    .format(key))

    def load_fields(self, fields_doc=None):
        """
        Returns a list of fields to expect in the output dictionaries
        and a second list that contains the fields that have a
        dictionary type.

        Reads these lists from the fields documentation.
        """

        if fields_doc is None:
            fields_doc = self.beat_path + "/_meta/fields.generated.yml"

        def extract_fields(doc_list, name):
            fields = []
            dictfields = []

            if doc_list is None:
                return fields, dictfields

            for field in doc_list:

                # Skip fields without name entry
                if "name" not in field:
                    continue

                # Chain together names
                if name != "":
                    newName = name + "." + field["name"]
                else:
                    newName = field["name"]

                if field.get("type") == "group":
                    subfields, subdictfields = extract_fields(field["fields"], newName)
                    fields.extend(subfields)
                    dictfields.extend(subdictfields)
                else:
                    fields.append(newName)
                    if field.get("type") in ["object", "geo_point"]:
                        dictfields.append(newName)
            return fields, dictfields

        # Not all beats have a fields.generated.yml. Fall back to fields.yml
        if not os.path.isfile(fields_doc):
            fields_doc = self.beat_path + "/_meta/fields.yml"

        # TODO: Make fields_doc path more generic to work with beat-generator
        with open(fields_doc, "r") as f:
            path = os.path.abspath(os.path.dirname(__file__) + "../../../../_meta/fields.generated.yml")
            if not os.path.isfile(path):
                path = os.path.abspath(os.path.dirname(__file__) + "../../../../_meta/fields.common.yml")
            with open(path) as f2:
                content = f2.read()

            #content = "fields:\n"
            content += f.read()
            doc = yaml.load(content)

            fields = []
            dictfields = []

            for item in doc:
                subfields, subdictfields = extract_fields(item["fields"], "")
                fields.extend(subfields)
                dictfields.extend(subdictfields)
            return fields, dictfields

    def flatten_object(self, obj, dict_fields, prefix=""):
        result = {}
        for key, value in obj.items():
            if isinstance(value, dict) and prefix + key not in dict_fields:
                new_prefix = prefix + key + "."
                result.update(self.flatten_object(value, dict_fields,
                                                  new_prefix))
            else:
                result[prefix + key] = value
        return result

    def copy_files(self, files, source_dir="files/", target_dir=""):
        if target_dir:
            target_dir = os.path.join(self.working_dir, target_dir)
        else:
            target_dir = self.working_dir
        for file_ in files:
            shutil.copy(os.path.join(source_dir, file_),
                        target_dir)

    def output_count(self, pred, output_file=None):
        """
        Returns true if the output line count predicate returns true
        """

        # Init defaults
        if output_file is None:
            output_file = "output/" + self.beat_name

        try:
            with open(os.path.join(self.working_dir, output_file), "r") as f:
                return pred(len([1 for line in f]))
        except IOError:
            return False

    def get_elasticsearch_url(self):
        """
        Returns an elasticsearch.Elasticsearch instance built from the
        env variables like the integration tests.
        """
        return "http://{host}:{port}".format(
            host=os.getenv("ES_HOST", "localhost"),
            port=os.getenv("ES_PORT", "9200"),
        )

    def get_kibana_url(self):
        """
        Returns kibana host URL
        """
        return "http://{host}:{port}".format(
            host=os.getenv("KIBANA_HOST", "localhost"),
            port=os.getenv("KIBANA_PORT", "5601"),
        )
