// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package vsphere

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded zlib format compressed contents of module/vsphere.
func AssetVsphere() string {
	return "eJzsWF1v2jAUfedXXPW99J2HSVWnbZVGW421r8gkN8GrPyL7mor9+skOySA4MBpne1keIpQ45xzb536Ya3jF7Qw2tlqjwQkAcRI4g6vNIjy5mgDkaDPDK+JazeDDBABg9xakzp3wnxkUyCzOoGQTgIKjyO0sDL0GxSTuU/iLtpUfbLSrdk8iLIdA+2A5I2ZJt3BxyF7Y3asIyOE8mqsrY19KxiqWcdpOC4M4XW0J7cG4RpjQquy8OKHNX58MIgRA0AXQGn8rnnYGF9pIRjM4pj+SSZqYSKrzu0dML9RZzJPqfLaYjySzyigq0mZMYL4shGbdAWfEPqHJUBEr0Ytthba8EHh7ZFf1x1HhhfXSompfcfumTX6hS7lAu7WEMqBMo6RrbWmaaafiy3T5Xj44uULjl8ZD2xOs/mfcQe+a7lduydMyIYKFPEeg69HAdZXKvwvSxvvh/uYR7rQiowWwsjRYMsIc7lXl6ObRUeUIHis0zANZ7yRYYKZVHlfo7+nW52PrVI8bZzTI0kb2N0YhTDwwV2UIFyiMlkOjPCgVjFBlTeaUqUTfbjBsJ5M+LkJK4hK9KGCBGHSziZG5AFcguRDchq3tcZ8lRi6h+xcB70z6bCu+HCnkX+Y9093IscP9ZX4q2N8Mp7Q9QONsj9w6m/RQX9dC/4WxA/M7nX3UBvrkO6QD7Hx/efOXNnd+aWpJvOOoXN1syPXPlB3R3dOzX/J5B3WfdmeQZLx1x/gHxKGnTscbOuoztBKlNmN1n/MA7ulj0OfjdidurB4+kbxxTkJDxSmkN21el4krxEMNC4ewbUXihhwTkmVrrgadVnuRLs9aoS/m3XkOTVwdwAMyfxspVzbNSCeXj5WjX+ptgHm9D/3pWif0WJd0192rEhb1yet/0RivaNxuGBdsJS6qHKVDS6PVD13AZ08wOFMHrSFCR5RaZ4ckJS/9sh4UvmTrGiyYXux+HRyuNXOWtFzWlSKqUK9+4NE/W/XD5YB0dheIIUL8V2v1rwAAAP//rR7c0A=="
}
