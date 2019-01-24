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

package amqp

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "amqp", asset.ModuleFieldsPri, AssetAmqp); err != nil {
		panic(err)
	}
}

// AssetAmqp returns asset data.
// This is the base64 encoded gzipped contents of protos/amqp.
func AssetAmqp() string {
	return "eJy0WE1v3DYQvftXDHJqAMtJi1zqQ4EgSQEf3C/kXsyKo9XU1JAmR/vRX19QWmrXXdpdR6pPu5T53tNw5nFmK3ig/S1g9+ivAJTV0i28+Xj/+29vrgAMxTqwV3ZyC2kRoqeaG66BNiQKDZM18eYKDp9urwAAKhDsaEJNf7r3dAvr4Pq8crrhdFMgb/dV7QxNj/J262R9svhE3U8nD2DUOiBBQgJ1gAIUggvXELljiyEttqq+zEg77HwKxof3H66eEam00zORD7TfumAu0/mVdgq08xZZWNagLY0qb845a4sxVmxmhOVnZJtoOtLWmRGxwDQ+XpTq7nOBh3Z1i7I+P+hXxfAX7AhcM8bugPgCWZUo5jF+OUANe29KSdOguF7PRXiMkTfn9CvnLKFcRn/XQCS9BuNAnEIdCPX46u8ee+pLATB9wJVdiDvQY09RATPsf/PTrrb9cm8/KZAjMjzHjb26ypAlXYj9BHAkhW1LAr30kUxBgLhqi3zuFt9EnjI9UthQgC1bO2RBoOidmGRq6fFYdiUXcRL7jkKleGEl3xkS5YYpQOPCgJ5BrmGDlg1sWVuW8VEfQroZUjII2VIekuUNhf2/FLzWYL5OQahSUa2FDKCYTFxNN1XmA8V10exixDVVtevl/Hxep0f6bkUhmdEBNcIhLEOOXMO25bodD21F8DcFl2L6BEdoa/eVodpiIDNuLJn0dJDLCs+wMX3BZwsquF5Z1tUD7ed56f0YqAyYNpfLB+uHxavHOIpD9dDOU62A9YO4rSWzpo5E45Dw+SjLsqyr0f5/ZR1JzDGZDsVdOxGqExRoi/oEzPcry7Elk/6zK2jmphpdap7oz6P3ObF74OZ54+Omos7reZbMYRsQC2RDts7syHKxDpDA2fzisYxP7W7Tuqil+qCD7cyN850YrlEHftSDt48l02KEFZGAD7Rh10e7h4l1zBWOT8BcABSnLQWoLZOUlHe9VfZz+4SPx0qaEF+qJAzrfii5M1q3+otqvYz112EVLaAxnD9m5KH5GgMTXZevyHgDn1CSG7vmCdYGQ4rpoKLYJKMYVBdmZvbxfCfA7IWlauo6Moxzu5g/Dr3TBDddkuWLRkl0gb75/u7+S4Z7vndOU9W7YSJ6XgtJ7QzLef/w7XoyZCECLaGhcEFqjgt/zrgGR6pxOAZNjfVLHVRXmpJfNzw5qTyFyHEIwnffv00mcbryw9uCAh/YBS64+uVtR37jDHUN71Np/ljMwBDIYkIpDaWveuGP3ttUcOnyPME9Gn1p9D6M+24mtTGB4nCLjz9NqCuOSZ7DoGmZxuqIN/14M3x7oRVeMsb5snoxvsodRcXOL/PKE1yJar6LZZa9H5uEAksfKcwO46c02KeeOKEBl7or9H45Gjw5toHtnwAAAP//OIaPDg=="
}
