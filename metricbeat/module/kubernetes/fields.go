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

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

<<<<<<< HEAD
// Asset returns asset data
func Asset() string {
	return "eJzsXU9v6zgOv79PQfTUB2R6WuyhhwVmO/uwxfszRV/fzGGxKBSbSTS1JY8kp5P99AvJtuzYku0kSpq21qm1E/IniqJIkVJ+gifcXMNTPkfBUKH8AKCoSvAaLj7bhxcfAGKUkaCZopxdwz8+AADUH4AUlaCR/rbABInEa1iSDwASlaJsKa/hPxdSJhczuFgplV38V79bcaEeI84WdHkNC5JI/ACwoJjE8tow+AkYSbEFTze1yTQHwfOsfOKAp9stW3CREv0YCItBKqKoVDSSwBeQ8VhCShhZYgzzTYPPVUmhiaaJiGRUolijsG9coHqAteT3890tFAQboqzatkir1obWhCfwzxyluooSikxtfaTC+YSbZy7i1rsetLrdGHoQc8qWoFZYMZK9KARKnosIw+G4LyhjDE7abQAynx8Tg498B0bEs/AAwJCFyyjJpUIxM0xlRiKcWel87MW1RjEPD+vfDw930CHd0VCeexQ04Wy5G+cHrkgCLE/nKPT0HqWcCVHIos2VzNNAMEoBSChJz0DmqcZT/E9RAmWQ0khwiRFn8TiAISVVjZFFuKfQ5nn0hG5QfP4HRu1XxcPHQLBhRaXiS0FSKIDIjp2OOFOEssPsdL0s1PSCmGmpiFCPiqZuqxAT1X4xIKDvmiB0CFppZLmTUVsWIzjd3P2AXJIlOgTh63YTivlu520foD6qW53kwkV4mPgQgyYT1u5vl41DvZttQL7NdmOVTkv9hgssRc8Ic5qQDlrCuBaLD/Qg4JFgC6XAeIChhcVjvMo6RmIblYxIgvHjIuHE98HCxbuGDEWETLkVa+duaAETCaRBVttH7fWoYqHhMQJJEh4RReYJ6u/19jehKVWvssMxLijDuOiBZm+e1sbwUj/xCgXoAnJmvoux2xVJ+LKtK3ubpi98qVfYBd/RJJE1oYnGfBSzNN+o/edfNeB9REaOtZGO7SpEJCMRVRvtkripW7tafvLtS6fQ5PGS0Sbv7UvFGPbxQqHaErj4hljhu57wNvXDl7Iilqjnibc7NayFwH7HIxQqzWgMII9ehgdkVMMBqAKSYspF23D49WAy1ODQQKcQT+FQn5dACjF4u3v+3uXXRgd2dDA9KgCvwccc0+0D3MxSLfyeZlNIQh5nYTqXmXL//Xv/PKkAP3PxRNlSdvZw4E3J4/eimyBRjZNLRpa4IHmi/HriQT4C0Te716bZgIePXTvJH1ycCI/h5UVlZw/najE+Whtazd9FXHHPuYIFTVBupMJ05xDjfbg8bik1nfD3Hom5JVT63y8XkZ0g0vjhiDEq9rjeznLuvMP/sMJmNtbQs0ltVBDxJMFI2TdqRRQQgbBEhoKoIn1cZDckiJwx2uovZZLGxtH53E5mww65A38I7JFzr3RvNJWCCwiMuIil8bjqbJCiKRbPMiIUjfKEiEIIsCISeBTlQmyNfYXQfFORNHOg7KpaX5ZkQYVUjyUr5knh7p4reagA6n4aHlDz0M/aWtVwssnRAWkWA3jq6Fp2nBl//rYXxNeCVKkMGFsffEnXyBwSiXi2eVTcBaLOGRLJWQh094bSWHBWETdZENk8bDIbqPRzTFGRmCgySvMHxqOgBERKHlFjaJ6pWvWOSd9ccs/K3d03a4cEalAdXfbOgRH2fmseGAaUs37JN1NdnmTqXrUM30hqx7yfpym4CMvYkNTL+vOKRqvS6j4TWS86bve8rPl4XKOQtDXzDgL1W0FwSyD9BTg5bbM4gP0PRv/MEWiMTNEFRQGKN4A4Cg5snh2TxWNC2VNAMPdfQGAmUGo0ZTWUzyBQtubJGuNHB8Zj2YWKp0sufRaCZDS85vx8dwvrbe3pGa4nygKqjeatKY5gHNZ4sIbx6GF6vPlaUd5B9GEn7I/bXwZ4NzdrD3HgGyU6Zmtwqs6ZqnM8LXR1zjetb6+7MGfK07nalKdrtXB5uikR0wI8JWLcwKdETE8ihqHSehPMXou/3rTy3WOEdG22an207IayEFwce1G+/8vHx+7WvO0BeRCEyZQqdT5j8uAcE7sTPWU9izZSmp+mhOeOAppynXXrCOc9pDlrH8BXVNkGdYpq2BrVedTB1nh8tbDWp8mZdwdnH7tNU+0BHqmu2b8mDDMYYgIjZziM3SIZM9Nht62U29R4vLuvGjBy5YD3LMYRawvsYuzeoQjdK5ANVnlTYIfsYWc8fpVb2FNEWrQpIq3baxqQVxeRvouc0ZlkSTqwzvSAyS7Hl9/bkWW9sNrzJLJ9oGTcWeXAWbIpIdSCfa7zajq4FXSy7X16631sDW5NGn+XWwnE7288gVhI5bmTRvRHEG88w1wIRKA0ZZFGIpL+b7AKISNL/HSsPGaBaXRO9e74WPz51HoX9Zwvitq7Kq/r5Vtc1F0+LBVRebg8V7Yi0m+o3R1od6LPDbTdMYzgsjx7M4NnQpX5Q6FIKSP9B4KRxP5U3JzzBEm7MHMkyhqhYeKW71Y1qCKiZypQpnC5paZ7gin4eNIHPYc3mmAOGr/fixGCS4vqxhT760G7EUSuvnCe/ZNET3yxmMG/hDCbcnd5kszA/lm+7w6tbtrhKEefcqYZpVmCCuNZLYkbwhhX9zkzLLiYwa+/fv1MkwTjj2X3rw4OvYdmSeH8+ULOQ23g9v1hhkvBsWfYq9sZT4FI2Js03fy2pdQXnA/gygRG2hBcw9+v/hYCucUyUp592IfhHUvqJy0KLQbR5/n0dnHIbdpJBKVnXUQsg0nlagBfHnc9bFXQ5EtHxJglfJMeeDS24dPUBIM4NWFPW3x24uzwqB1gR36rb7nfi33BxbXm14qVJTQi4x2uvXBUXPa5Fi9GSUVPqclBDskvjaGyOf2SY436UmYYHbIhEApjnUL1jFtjW4+dDlaD1whgWew8shkcVMGnC+jMz/yEjH7644qDvGdz7KQZUsClEjnOiov0teubsyfGn5l/3uRMRiuM834lPSj6MSi3+PQZw5AudWN7ccCN9W2mju2edqaam5n9TmxVzHI039pismUzpzvw05D5S3lK33x7y2PvN31Z5CXa/pInd8EJBBs7s0t/LNVsjk3GHQfkOgNyVDgmSdEuvDvfih7qvrmh83i8m6iR3d45ma24VI/H4ahJ+9juuAjvxrhcLPercTjibmYLZrmdeV9tZ94hiylbXl1d7buLGRLdYX5H6Q30+KAhsVpuLryzLtp2ZIahwueSYJm1OuP4uQnUG0AfMXBt8vdH0CHOHx+QV9q6JstGqhkKuC/+caVix8bUL4Wr34KEQ6Wtx67Y+Nz8QNaxhFZeqmOuqig5wXxj6hhqcCatJ3iSOOJju8FJ5thn20JJcZEnyabiNijN5tqKizwJZ9Yqiudv17aQeg2b+0or79gNXeq2wvoOK3v7FlxixqPVR1Mq872E1Vb+E1jaLYlYFdrL2B55etZ6b2fnlsr7hAgvYHU7+5d9ACtwtf059jg3LB2tfx7xvIbbDnID7HkMczW4I4BZm2uOkIQyt8V5lEZdXQib6yp/gcMMb52y8pra6YalTgtwivDN3bA0Xa7kIDeVjb+jKs/pHqHtNt0jdOg9QhWaNU/yNFQatiB2hkHgbwUwryMyXexStulil+liF/cHpotdDuu06/crXFBOcHvKp5G/H3i631kswfw/AAD//7QUBBw="
=======
// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
>>>>>>> add specific asset name to allow multiple assets per package
}
