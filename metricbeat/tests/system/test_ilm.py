from metricbeat import BaseTest
import os
from elasticsearch import Elasticsearch, TransportError
from nose.plugins.attrib import attr
import unittest
import re
import datetime

INTEGRATION_TESTS = os.environ.get('INTEGRATION_TESTS', False)


class Test(BaseTest):

    COMPOSE_SERVICES = ['elasticsearch']

    @unittest.skipUnless(INTEGRATION_TESTS, "integration test")
    @attr('integration')
    def test_ilm_enabled(self):
        """
        Test ilm sends data to correct index
        """

        self.render_config_template(
            modules=[{
                "name": "system",
                "metricsets": ["cpu"],
                "period": "1s"
            }],
            elasticsearch={
                "host": self.get_elasticsearch_url(),
                "ilm": True,
            },
        )

        proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("metricbeat start running."))
        self.wait_until(lambda: self.log_contains("Overwriting setup.template for ILM"))
        self.wait_until(lambda: self.log_contains("PublishEvents: 1 events have been published"))
        proc.check_kill_and_wait()

        es = Elasticsearch([self.get_elasticsearch_url()])

        # Extracts metricbeat version from log to make it cross version compatible
        log = self.get_log()
        r = re.findall('"version": "(.*)"', log)
        beat_name = "metricbeat-" + r[0]

        # Check if template is loaded with settings
        template = es.transport.perform_request('GET', '/_template/' + beat_name)
        assert template[beat_name]["settings"]["index"]["lifecycle"]["name"] == "beats-default-policy"
        assert template[beat_name]["settings"]["index"]["lifecycle"]["rollover_alias"] == beat_name

        # Make sure the correct index + alias was created
        alias = es.transport.perform_request('GET', '/_alias/' + beat_name)
        d = datetime.datetime.now()
        now = d.strftime("%Y.%m.%d")
        index_name = beat_name + "-" + now + "-000001"
        assert index_name in alias
        assert alias[index_name]["aliases"][beat_name]["is_write_index"] == True

        # Asserts that data is actually written to the ILM indices
        self.wait_until(lambda: es.transport.perform_request(
            'GET', '/' + index_name + '/_search')["hits"]["total"] > 0)

        data = es.transport.perform_request('GET', '/' + index_name + '/_search')
        assert data["hits"]["total"] > 0
