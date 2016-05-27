from topbeat import BaseTest

"""
Contains tests for filtering.
"""


class Test(BaseTest):
    def test_dropfields(self):
        """
        Check drop_fields filtering action
        """
        self.render_config_template(
            system_stats=False,
            process_stats=True,
            filesystem_stats=False,
            drop_fields={"fields": ["proc"]},
        )
        topbeat = self.start_beat()
        self.wait_until(
            lambda: self.output_count(lambda x: x >= 1),
            max_timeout=15)

        topbeat.kill_and_wait()

        output = self.read_output(
            required_fields=["@timestamp", "type"],
        )[0]

        for key in [
            "proc.cpu.start_time",
            "proc.cpu.total.pct",
            "proc.name",
            "proc.state",
            "proc.pid",
        ]:
            assert key not in output

    def test_dropfields_with_condition(self):
        """
        Check drop_fields action works when a condition is associated.
        """
        self.render_config_template(
            system_stats=False,
            process_stats=True,
            filesystem_stats=False,
            drop_fields={
                "fields": ["proc.memory"],
                "condition": "range.proc.cpu.total.pct.lt: 0.5",
            },
        )
        topbeat = self.start_beat()
        self.wait_until(
            lambda: self.output_count(lambda x: x >= 1),
            max_timeout=15)

        topbeat.kill_and_wait()

        output = self.read_output(
            required_fields=["@timestamp", "type"],
        )

        for event in output:
            if float(event["proc.cpu.total.pct"]) < 0.5:
                assert "proc.memory.size" not in event
            else:
                assert "proc.memory.size" in event

    def test_dropevent_with_condition(self):
        """
        Check drop_event action works when a condition is associated.
        """
        self.render_config_template(
            system_stats=False,
            process_stats=True,
            filesystem_stats=False,
            drop_event={
                "condition": "range.proc.cpu.total.pct.lt: 0.001",
            },
        )
        topbeat = self.start_beat()
        self.wait_until(
            lambda: self.output_count(lambda x: x >= 1),
            max_timeout=15)

        topbeat.kill_and_wait()

        output = self.read_output(
            required_fields=["@timestamp", "type"],
        )
        for event in output:
            assert float(event["proc.cpu.total.pct"]) >= 0.001

    def test_include_fields(self):
        """
        Check include_fields filtering action
        """
        self.render_config_template(
            system_stats=False,
            process_stats=True,
            filesystem_stats=False,
            include_fields={"fields": ["proc.cpu", "proc.memory"]},
        )
        topbeat = self.start_beat()
        self.wait_until(
            lambda: self.output_count(lambda x: x >= 1),
            max_timeout=15)

        topbeat.kill_and_wait()

        output = self.read_output(
            required_fields=["@timestamp", "type"],
        )[0]
        print(output)

        for key in [
            "proc.cpu.start_time",
            "proc.cpu.total.pct",
            "proc.memory.size",
            "proc.memory.rss.bytes",
            "proc.memory.rss.pct"
        ]:
            assert key in output

        for key in [
            "proc.name",
            "proc.pid",
        ]:
            assert key not in output

    def test_multiple_actions(self):
        """
        Check the result when configuring two actions: include_fields
        and drop_fields.
        """
        self.render_config_template(
            system_stats=False,
            process_stats=True,
            filesystem_stats=False,
            include_fields={"fields": ["proc"]},
            drop_fields={"fields": ["proc.memory"]},
        )
        topbeat = self.start_beat()
        self.wait_until(
            lambda: self.output_count(lambda x: x >= 1),
            max_timeout=15)

        topbeat.kill_and_wait()

        output = self.read_output(
            required_fields=["@timestamp", "type"],
        )[0]

        for key in [
            "proc.cpu.start_time",
            "proc.cpu.total.pct",
            "proc.name",
            "proc.pid",
        ]:
            assert key in output

        for key in [
            "proc.memory.size",
            "proc.memory.rss.bytes",
            "proc.memory.rss.pct"
        ]:
            assert key not in output

    def test_contradictory_multiple_actions(self):
        """
        Check the behaviour of a contradictory multiple actions
        """
        self.render_config_template(
            system_stats=False,
            process_stats=True,
            filesystem_stats=False,
            include_fields={"fields": ["proc.memory.size", "proc.memory.rss.pct"]},
            drop_fields={"fields": ["proc.memory.size", "proc.memory.rss.pct"]},
        )
        topbeat = self.start_beat()
        self.wait_until(
            lambda: self.output_count(lambda x: x >= 1),
            max_timeout=15)
        topbeat.kill_and_wait()

        output = self.read_output(
            required_fields=["@timestamp", "type"],
        )[0]

        for key in [
            "proc.memory.size",
            "proc.memory.rss",
            "proc.cpu.start_time",
            "proc.cpu.total.pct",
            "proc.name",
            "proc.pid",
            "proc.memory.rss.pct"
        ]:
            assert key not in output
