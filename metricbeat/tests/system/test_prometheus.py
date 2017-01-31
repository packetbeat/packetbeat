import os
import metricbeat
import unittest
from nose.plugins.attrib import attr

PROMETHEUS_FIELDS = metricbeat.COMMON_FIELDS + ["prometheus"]


class Test(metricbeat.BaseTest):

    @unittest.skipUnless(metricbeat.INTEGRATION_TESTS, "integration test")
    def test_stats(self):
        """
        prometheus stats test
        """
        self.render_config_template(modules=[{
            "name": "prometheus",
            "metricsets": ["stats"],
            "hosts": self.get_hosts(),
            "period": "5s"
        }])
        proc = self.start_beat()
        self.wait_until(lambda: self.output_lines() > 0)
        proc.check_kill_and_wait()

        # Ensure no errors or warnings exist in the log.
        log = self.get_log()
        self.assertNotRegexpMatches(log.replace("WARN EXPERIMENTAL", ""), "ERR|WARN")

        output = self.read_output_json()
        self.assertEqual(len(output), 1)
        evt = output[0]

        self.assertItemsEqual(self.de_dot(PROMETHEUS_FIELDS), evt.keys(), evt)

        self.assert_fields_are_documented(evt)

    def get_hosts(self):
        return ["http://" + os.getenv('PROMETHEUS_HOST', 'localhost') + ':' +
                os.getenv('PROMETHEUS_PORT', '9090')]
