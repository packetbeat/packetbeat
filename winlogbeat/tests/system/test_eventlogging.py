import sys
import time
import unittest
from winlogbeat import WriteReadTest

if sys.platform.startswith("win"):
    import win32security

"""
Contains tests for reading from the Event Logging API (pre MS Vista).
"""


@unittest.skipUnless(sys.platform.startswith("win"), "requires Windows")
class Test(WriteReadTest):
    @classmethod
    def setUpClass(self):
        self.api = "eventlogging"
        super(WriteReadTest, self).setUpClass()

    def test_read_one_event(self):
        """
        eventlogging - Read one classic event
        """
        msg = "Hello world!"
        self.write_event_log(msg)
        evts = self.read_events()
        self.assertTrue(len(evts), 1)
        self.assert_common_fields(evts[0], msg=msg)

    def test_read_unknown_event_id(self):
        """
        eventlogging - Read unknown event ID
        """
        msg = "Unknown event ID"
        event_id = 1111
        self.write_event_log(msg, eventID=event_id)
        evts = self.read_events()
        self.assertTrue(len(evts), 1)
        self.assert_common_fields(evts[0], eventID=event_id)
        self.assertEqual(evts[0]["message_error"].lower(),
                         ("The system cannot find message text for message "
                          "number 1111 in the message file for "
                          "C:\\Windows\\system32\\EventCreate.exe.").lower())

    def test_read_unknown_sid(self):
        """
        eventlogging - Read event with unknown SID
        """
        # Fake SID that was made up.
        accountIdentifier = "S-1-5-21-3623811015-3361044348-30300820-1013"
        sid = win32security.ConvertStringSidToSid(accountIdentifier)

        msg = "Unknown SID " + accountIdentifier
        self.write_event_log(msg, sid=sid)
        evts = self.read_events()
        self.assertTrue(len(evts), 1)
        self.assert_common_fields(evts[0], msg=msg, sid=accountIdentifier)

    def test_fields_under_root(self):
        """
        eventlogging - Add tags and custom fields under root
        """
        msg = "Add tags and fields under root"
        self.write_event_log(msg)
        evts = self.read_events(config={
            "tags": ["global"],
            "fields": {"global": "field", "env": "prod", "level": "overwrite"},
            "fields_under_root": True,
            "event_logs": [
                {
                    "name": self.providerName,
                    "api": self.api,
                    "tags": ["local"],
                    "fields_under_root": True,
                    "fields": {"local": "field", "env": "dev"}
                }
            ]
        })
        self.assertTrue(len(evts), 1)
        self.assert_common_fields(evts[0], msg=msg, level="overwrite", extra={
            "global": "field",
            "env": "dev",
            "local": "field",
            "tags": ["global", "local"],
        })

    def test_fields_not_under_root(self):
        """
        eventlogging - Add custom fields (not under root)
        """
        msg = "Add fields (not under root)"
        self.write_event_log(msg)
        evts = self.read_events(config={
            "fields": {"global": "field", "env": "prod", "level": "overwrite"},
            "event_logs": [
                {
                    "name": self.providerName,
                    "api": self.api,
                    "fields": {"local": "field", "env": "dev", "num": 1}
                }
            ]
        })
        self.assertTrue(len(evts), 1)
        self.assert_common_fields(evts[0], msg=msg, extra={
            "fields.global": "field",
            "fields.env": "dev",
            "fields.level": "overwrite",
            "fields.local": "field",
            "fields.num": 1,
        })
        self.assertTrue("tags" not in evts[0])

    def test_ignore_older(self):
        """
        eventlogging - Query by time (ignore_older than 4s)
        """
        self.write_event_log(">=4 seconds old", eventID=20)
        time.sleep(4)
        self.write_event_log("~0 seconds old", eventID=10)
        evts = self.read_events(config={
            "event_logs": [
                {
                    "name": self.providerName,
                    "api": self.api,
                    "ignore_older": "2s"
                }
            ]
        }, expected_events=1)
        self.assertTrue(len(evts), 1)
        self.assertEqual(evts[0]["event_id"], 10)

    def test_unknown_eventlog_config(self):
        """
        eventlogging - Unknown config parameter
        """
        self.render_config_template(
            event_logs=[
                {
                    "name": self.providerName,
                    "api": self.api,
                    "event_id": "10, 12",
                    "level": "info",
                    "provider": ["me"],
                    "include_xml": True,
                }
            ]
        )
        self.start_beat(extra_args=["-configtest"]).check_wait(exit_code=1)
        assert self.log_contains("4 errors: Invalid event log key")
