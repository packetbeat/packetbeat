package actions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/stretchr/testify/assert"
)

const gceMetadataV1 = `{
  "instance": {
    "attributes": {},
    "cpuPlatform": "Intel Haswell",
    "description": "",
    "disks": [
      {
        "deviceName": "test-gce-dev",
        "index": 0,
        "mode": "READ_WRITE",
        "type": "PERSISTENT"
      }
    ],
    "hostname": "test-gce-dev.c.test-dev.internal",
    "id": 3910564293633576924,
    "image": "",
    "licenses": [
      {
        "id": "1000000"
      }
    ],
    "machineType": "projects/111111111111/machineTypes/f1-micro",
    "maintenanceEvent": "NONE",
    "networkInterfaces": [
      {
        "accessConfigs": [
          {
            "externalIp": "10.10.10.10",
            "type": "ONE_TO_ONE_NAT"
          }
        ],
        "forwardedIps": [],
        "ip": "10.10.0.2",
        "ipAliases": [],
        "mac": "44:00:00:00:00:01",
        "network": "projects/111111111111/networks/default"
      }
    ],
    "scheduling": {
      "automaticRestart": "TRUE",
      "onHostMaintenance": "MIGRATE",
      "preemptible": "FALSE"
    },
    "serviceAccounts": {
      "111111111111-compute@developer.gserviceaccount.com": {
        "aliases": [
          "default"
        ],
        "email": "111111111111-compute@developer.gserviceaccount.com",
        "scopes": [
          "https://www.googleapis.com/auth/devstorage.read_only",
          "https://www.googleapis.com/auth/logging.write",
          "https://www.googleapis.com/auth/monitoring.write",
          "https://www.googleapis.com/auth/servicecontrol",
          "https://www.googleapis.com/auth/service.management.readonly"
        ]
      },
      "default": {
        "aliases": [
          "default"
        ],
        "email": "111111111111-compute@developer.gserviceaccount.com",
        "scopes": [
          "https://www.googleapis.com/auth/devstorage.read_only",
          "https://www.googleapis.com/auth/logging.write",
          "https://www.googleapis.com/auth/monitoring.write",
          "https://www.googleapis.com/auth/servicecontrol",
          "https://www.googleapis.com/auth/service.management.readonly"
        ]
      }
    },
    "tags": [],
    "virtualClock": {
      "driftToken": "0"
    },
    "zone": "projects/111111111111/zones/us-east1-b"
  },
  "project": {
    "attributes": {
      "sshKeys": "developer:ssh-rsa 222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222 google-ssh {\"userName\":\"foo@bar.com\",\"expireOn\":\"2016-10-06T20:20:41+0000\"}\ndev:ecdsa-sha2-nistp256 4444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444= google-ssh {\"userName\":\"foo@bar.com\",\"expireOn\":\"2016-10-06T20:20:40+0000\"}\ndev:ssh-rsa 444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444 dev"
    },
    "numericProjectId": 111111111111,
    "projectId": "test-dev"
  }
}`

func initGCETestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/computeMetadata/v1/?recursive=true&alt=json" {
			w.Write([]byte(gceMetadataV1))
			return
		}

		http.Error(w, "not found", http.StatusNotFound)
	}))
}

func TestRetrieveGCEMetadata(t *testing.T) {
	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, []string{"*"})
	}

	server := initGCETestServer()
	defer server.Close()

	config, err := common.NewConfigFrom(map[string]interface{}{
		"host": server.Listener.Addr().String(),
	})
	if err != nil {
		t.Fatal(err)
	}

	p, err := newCloudMetadata(*config)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := p.Run(common.MapStr{})
	if err != nil {
		t.Fatal(err)
	}

	expected := common.MapStr{
		"meta": common.MapStr{
			"cloud": common.MapStr{
				"provider":          "gce",
				"instance_id":       "3910564293633576924",
				"machine_type":      "projects/111111111111/machineTypes/f1-micro",
				"availability_zone": "projects/111111111111/zones/us-east1-b",
				"project_id":        "test-dev",
			},
		},
	}
	assert.Equal(t, expected, actual)
}
