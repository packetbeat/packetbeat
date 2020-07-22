// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

var crowdstrikeFalcon = (function () {
    var processor = require("processor");

    var convertUnderscore = function (text) {
        return text.split(/(?=[A-Z])/).join('_').toLowerCase();
    };

    var convertToMSEpoch = function (evt, field) {
        var timestamp = evt.Get(field);
        if (timestamp) {
            if (timestamp < 100000000000) { // check if we have a seconds timestamp, this is roughly 1973 in MS
                evt.Put(field, timestamp * 1000);
            }
            (new processor.Timestamp({
                field: field,
                target_field: field,
                timezone: "UTC",
                layouts: ["UNIX_MS"]
            })).Run(evt);
        }
    };

    var decodeJson = new processor.DecodeJSONFields({
        fields: ["message"],
        target: "crowdstrike",
        process_array: true,
        max_depth: 8
    });

    var dropFields = function (evt) {
        evt.Delete("message");
        evt.Delete("host.name");
    };

    var setFields = function (evt) {
        evt.Put("agent.name", "falcon");
    };

    var convertFields = new processor.Convert({
        fields: [
            // DetectionSummaryEvent
            {
                from: "crowdstrike.event.LocalIP",
                to: "source.ip",
                type: "ip"
            },
            {
                from: "crowdstrike.event.ProcessId",
                to: "process.pid"
            },
            {
                from: "crowdstrike.event.ParentImageFileName",
                to: "process.parent.executable"
            },
            {
                from: "crowdstrike.event.ParentCommandLine",
                to: "process.parent.command_line"
            },
            // UserActivityAuditEvent and AuthActivityAuditEvent
            {
                from: "crowdstrike.event.UserIp",
                to: "source.ip",
                type: "ip"
            },
            // FirewallRuleIP4Matched
            {
                from: "crowdstrike.event.Ipv",
                to: "network.type",
            },
        ],
        mode: "copy",
        ignore_missing: true,
        fail_on_error: false
    });

    var addTimestamp = new processor.Convert({
        fields: [{
            from: "crowdstrike.metadata.eventCreationTime",
            to: "@timestamp",
        }],
        mode: "copy",
        ignore_missing: false,
        fail_on_error: true
    });

    var normalizeEpochMS = function (evt) {
        convertToMSEpoch(evt, "crowdstrike.event.ProcessStartTime")
        convertToMSEpoch(evt, "crowdstrike.event.ProcessEndTime")
        convertToMSEpoch(evt, "crowdstrike.event.IncidentStartTime")
        convertToMSEpoch(evt, "crowdstrike.event.IncidentEndTime")
        convertToMSEpoch(evt, "crowdstrike.event.StartTimestamp")
        convertToMSEpoch(evt, "crowdstrike.event.EndTimestamp")
        convertToMSEpoch(evt, "crowdstrike.event.UTCTimestamp")
        convertToMSEpoch(evt, "crowdstrike.metadata.eventCreationTime")
    };

    var normalizeProcess = function (evt) {
        var commandLine = evt.Get("crowdstrike.event.CommandLine")
        if (commandLine && commandLine !== "") {
            var args = commandLine.split(' ')
            var executable = args[0]

            evt.Put("process.command_line", commandLine)
            evt.Put("process.args", args)
            evt.Put("process.executable", executable)
        }
    }

    var processEvent = function (evt) {
        var eventType = evt.Get("crowdstrike.metadata.eventType")
        var outcome = evt.Get("crowdstrike.event.Success")

        evt.Put("event.kind", "event")

        if (outcome === true) {
            evt.Put("event.outcome", "success")
        } else if (outcome === false) {
            evt.Put("event.outcome", "failure")
        } else {
            evt.Put("event.outcome", "unknown")
        }

        switch (eventType) {
            case "DetectionSummaryEvent":
                var tactic = evt.Get("crowdstrike.event.Tactic").toLowerCase()
                var technique = evt.Get("crowdstrike.event.Technique").toLowerCase()
                evt.Put("threat.technique.name", technique)
                evt.Put("threat.tactic.name", tactic)

                evt.Put("event.action", evt.Get("crowdstrike.event.PatternDispositionDescription"))
                evt.Put("event.kind", "alert")
                evt.Put("event.type", ["info"])
                evt.Put("event.category", ["malware"])
                evt.Put("event.url", evt.Get("crowdstrike.event.FalconHostLink"))
                evt.Put("event.dataset", "crowdstrike.falcon_endpoint")

                evt.Put("event.severity", evt.Get("crowdstrike.event.Severity"))
                evt.Put("message", evt.Get("crowdstrike.event.DetectDescription"))
                evt.Put("process.name", evt.Get("crowdstrike.event.FileName"))

                normalizeProcess(evt);

                evt.Put("user.name", evt.Get("crowdstrike.event.UserName"))
                evt.Put("user.domain", evt.Get("crowdstrike.event.MachineDomain"))
                evt.Put("agent.id", evt.Get("crowdstrike.event.SensorId"))
                evt.Put("host.name", evt.Get("crowdstrike.event.ComputerName"))
                evt.Put("agent.type", "falcon")
                evt.Put("file.hash.sha256", evt.Get("crowdstrike.event.SHA256String"))
                evt.Put("file.hash.md5", evt.Get("crowdstrike.event.MD5String"))
                evt.Put("rule.name", evt.Get("crowdstrike.event.DetectName"))
                evt.Put("rule.description", evt.Get("crowdstrike.event.DetectDescription"))

                break;

            case "IncidentSummaryEvent":
                evt.Put("event.kind", "alert")
                evt.Put("event.type", ["info"])
                evt.Put("event.category", ["malware"])
                evt.Put("event.action", "incident")
                evt.Put("event.url", evt.Get("crowdstrike.event.FalconHostLink"))
                evt.Put("event.dataset", "crowdstrike.falcon_endpoint")

                evt.Put("message", "Incident score " + evt.Get("crowdstrike.event.FineScore"))

                break;

            case "UserActivityAuditEvent":
                var userid = evt.Get("crowdstrike.event.UserId")
                evt.Put("user.name", userid)
                if (userid.split('@').length == 2) {
                    evt.Put("user.email", userid)
                }

                evt.Put("message", evt.Get("crowdstrike.event.OperationName"))
                evt.Put("event.action", convertUnderscore(eventType))
                evt.Put("event.type", ["change"])
                evt.Put("event.category", ["iam"])
                evt.Put("event.dataset", "crowdstrike.falcon_audit")

                break;

            case "FirewallMatchEvent":
                evt.Put("message", "Firewall Rule '" + evt.Get("crowdstrike.event.RuleName") + "' triggered")

                evt.Put("event.category", ["network"])
                evt.Put("event.type", ["connection", "start"])
                evt.Put("event.outcome", ["unknown"])
                evt.Put("event.action", convertUnderscore(eventType))
                evt.Put("event.code", evt.Get("crowdstrike.event.EventType"))
                evt.Put("event.dataset", "crowdstrike.falcon_endpoint")
                evt.Put("process.pid", evt.Get("crowdstrike.event.PID"))

                normalizeProcess(evt);

                evt.Put("rule.id", evt.Get("crowdstrike.event.RuleId"))
                evt.Put("rule.name", evt.Get("crowdstrike.event.RuleName"))
                evt.Put("rule.ruleset", evt.Get("crowdstrike.event.RuleGroupName"))
                evt.Put("rule.description", evt.Get("crowdstrike.event.RuleDescription"))
                evt.Put("rule.category", evt.Get("crowdstrike.event.RuleFamilyID"))

                evt.Put("host.name", evt.Get("crowdstrike.event.HostName"))

                var localAddress = evt.Get("crowdstrike.event.LocalAddress");
                var localPort = evt.Get("crowdstrike.event.LocalPort");
                var remoteAddress = evt.Get("crowdstrike.event.RemoteAddress");
                var remotePort = evt.Get("crowdstrike.event.RemotePort");
                if (evt.Get("crowdstrike.event.ConnectionDirection") === "1") {
                    evt.Put("network.direction", "inbound")
                    evt.Put("source.ip", remoteAddress)
                    evt.Put("source.port", remotePort)
                    evt.Put("destination.ip", localAddress)
                    evt.Put("destination.port", localPort)
                } else {
                    evt.Put("network.direction", "outbound")
                    evt.Put("destination.ip", remoteAddress)
                    evt.Put("destination.port", remotePort)
                    evt.Put("source.ip", localAddress)
                    evt.Put("source.port", localPort)
                }
                break;

            case "AuthActivityAuditEvent":
                var userid = evt.Get("crowdstrike.event.UserId")
                evt.Put("user.name", userid)
                if (userid.split('@').length == 2) {
                    evt.Put("user.email", userid)
                }

                evt.Put("message", evt.Get("crowdstrike.event.ServiceName"))
                evt.Put("event.action", convertUnderscore(evt.Get("crowdstrike.event.OperationName")))
                evt.Put("event.type", ["change"])
                evt.Put("event.category", ["authentication"])
                evt.Put("event.dataset", "crowdstrike.falcon_audit")

                break;

            case "RemoteResponseSessionStartEvent":
            case "RemoteResponseSessionEndEvent":
                var username = evt.Get("crowdstrike.event.UserName")
                evt.Put("user.name", username)
                if (username.split('@').length == 2) {
                    evt.Put("user.email", username)
                }

                evt.Put("host.name", evt.Get("crowdstrike.event.HostnameField"))
                evt.Put("event.action", convertUnderscore(eventType))
                evt.Put("event.dataset", "crowdstrike.falcon_audit")

                if (eventType == "RemoteResponseSessionStartEvent") {
                    evt.Put("event.type", ["start"])
                    evt.Put("message", "Remote response session started")
                } else {
                    evt.Put("event.type", ["end"])
                    evt.Put("message", "Remote response session ended")
                }

                break;

            default:
                break;
        }
    }

    var pipeline = new processor.Chain()
        .Add(decodeJson)
        .Add(normalizeEpochMS)
        .Add(dropFields)
        .Add(addTimestamp)
        .Add(convertFields)
        .Add(processEvent)
        .Add(setFields)
        .Build();

    return {
        process: pipeline.Run,
    };
})();

function process(evt) {
    crowdstrikeFalcon.process(evt);
}
