import os
import metricbeat
import unittest

PROMETHEUS_FIELDS = metricbeat.COMMON_FIELDS + ["prometheus"]


class Test(metricbeat.BaseTest):

    COMPOSE_SERVICES = ['prometheus']

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
        self.assert_no_logged_warnings()

        output = self.read_output_json()
        self.assertEqual(len(output), 1)
        evt = output[0]

        self.assertItemsEqual(self.de_dot(PROMETHEUS_FIELDS), evt.keys(), evt)

        self.assert_fields_are_documented(evt)

    def get_hosts(self):
        return [self.compose_host()]
