import os
import metricbeat
import unittest

NATS_FIELDS = metricbeat.COMMON_FIELDS + ["nats"]


class Test(metricbeat.BaseTest):

    COMPOSE_SERVICES = ['nats']

    @unittest.skipUnless(metricbeat.INTEGRATION_TESTS, "integration test")
    def test_varz(self):
        """
        nats varz test
        """
        self.render_config_template(modules=[{
            "name": "nats",
            "metricsets": ["varz"],
            "hosts": self.get_hosts(),
            "period": "5s",
            "varz_metrics_path": "/varz"
        }])
        print(self.get_hosts())
        proc = self.start_beat()
        self.wait_until(lambda: self.output_lines() > 0)
        proc.check_kill_and_wait()
        self.assert_no_logged_warnings()

        output = self.read_output_json()
        self.assertEqual(len(output), 1)
        evt = output[0]

        self.assertItemsEqual(self.de_dot(NATS_FIELDS), evt.keys(), evt)

        self.assert_fields_are_documented(evt)

    def get_hosts(self):
        return ["nats:8222"]
