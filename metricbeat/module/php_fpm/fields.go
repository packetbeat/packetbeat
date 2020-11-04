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

package php_fpm

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "php_fpm", Asset); err != nil {
		panic(err)
	}
}

// AssetPhpFpm returns asset data.
// This is the base64 encoded gzipped contents of module/php_fpm.
func AssetPhpFpm() string {
	return "eJzMWV1v67gRfc+vGOQlDuoIu68pUKC43e3Nw26Ne5OnopDH1NgiQpFacmTH++uLoShZluV8NN5i9XCBiObwzOGZL907eKb9PdRlna/r6gqANRu6h+vF10X+8+KX6yuAgoLyumbt7D387QoAYPF1cffz4hcI5LfkITByE6Ai9loFUM4YUkwFrL2ruh9nVwChdJ5z5exab+5hjSbQFYAnQxjoHjYovyFmbTfhHv59HYK5nsN1yVxf/+cKYK3JFOE+YrgDixUNscvD+1oMedfU6c0EfHmWad8SlLOM2gbgknofuESGHXkCt5LVkTPJ5d5ajRsChcZk6dUQ6RFa50z/cgruK5Bb2M6ZD2KW33S4a+8UhRBxZAPLY7xDzPLv0UKH+5n2O+eL0dor6OV5LClaBLeOyI6R/Ml4OrI55uxYt+9hMlnNK7S4IX85Ur8zslZzKPYWK63AeXC2oAptkU0iUc5aUmIuTKIYU/0ODF96kzE6CEJNSq+1in/qwFqFbLRriqwDSFSKaqYxFx1G4+zmZOkNkEl/TbUiLwrUVrlK2w14+q2hwEkdQymkJFdi6AH9dcLuriQLOCAW9GEDaJY/PVVuO5Sb0YHJwm8NNTTmpmMhLl6eA9V4T5YHXAw00bJQ4pZgRWRBW80amYo5rBoG63jC6p64dziDBwlvHWCLpiFx3jp79zt5J1zwvtaSLvdQEXbHoTETVoUm3KI2uDLUXUYKJAqAvnfF7GHVhP0c0BayzVNctW7C6sBAb5tdaz4eaemFO1Fk8A110BNMS6KpsgpfclVqU3iyS5jV3m11QRFDB1ihhRJtYQg030ohbEwBJZlxkMnzTFS35KXLMW6XwaO8qL2ryfMe1s4YtwsHKa1RceRxwmAn5VZuAbYaASE49UwMs8cvC0kYa20IVhiouO0obAJoW5LXPM4S8gR3yOKqRI+KybdxLq9b8ych38laSGvx5FHhFxd4hS+6aqqBwPsY1zaCjufKQk22mL7eYUwEbRWB8Ci5IDB6EfoZ94au5Ybsxd0L+vee/XSTY3+G6KfLQB8GlyoCi9RfXKgC6MJcXhgHQbTB2ethuUMtvadkgsTMMnky0xllZ/LeIf2IPdmPsSG8TTHbZsAU8ms0BlbEO0mrXJ56l1KKtnmo0VOeMC5jSHbZZrQUTwhtCZq22fV9qfOQfLxMzcLynIRRsd7+P+jvSlECeWAQ7trGTdsYyQL6x+WE2VlwQFuy4KQErxspLCKddMI8WvEUGsOw0/EC5C7BExaw/GF5e44CdoynRemC7YeA/Eti+hCNr+XMP+hWTrPlGNTH8t+wJuaeUJUX7mF+7YGyrii0d9zJ3OhKc4QZ25d0/vxc21ZXwF5TiC2A+AWVk74i4YfZzvnnAM4aKbzjvl2euoKbFE83MVJvuv775nY69QbjdnlXkSbT7wQv75mtjnkB7Moe0AupJnamshLZoRdFVEzczDJtygWmcZtctriGl2ecEdLyKJALufLrIEsoZ4t3ye8YjkCeRFMgjxfeQazsavsiYe8sjlFp/cwQ26n5s1Ps/zK09m7oY3G0TqDROBqOkcve6Wy8q9Ibj62L7Ee93qvUC/GLh3/0HwtOSB1c+ehSz4/Tb57Y1vzjM2H2UBiaw7fGWm03cyBWt9NAprR3RnlvQjlW3RDPQH1nYZxGZItDW6bxB4g3odgzIfl+UJPp7iKIBuP7CIx0AK9gyYumVealMHX2pKmstPKuI2v2I1TamLSEicTbfoJKTVCXrWfVHtzOQkHrOH07O6W2zouKuHQfidSSuc66+XZi96citvOhtQuzf/70OIfFv74/tlEDs2mfX3Ow8foD3jXeZM7rjbaj9u0ibj19e4Cd5rIbIv0eAntpWz/gmCR1siyz4YbLj99cqJ0NlK1csc9Wex6NcJ/yM0GDFlon0F6YsQ2K/suVfug2mzAKpzeuMZDPTj4/f8o3MQmzxddF/venx6/50/efvokHHu5Ar+NQEIhvYbZ2/r1OtSdeqvZUqG0yGf9HgYpPwjMYuGs0c1U3E0BP2rM3UX5ZPEFNXpFlmdPit0wMh+ylnA1NFb8D8k0ANDvcB/hBnBhm6PhZkCUfPsRxLdbdFSlsQnuGQqMag9331MJZ6sfc/rDB3Cgpn8lX2uJ0CTpio6LK+f2lMn+FL4CVayxLyLS2/xhmku1PkvPfAAAA//9lCOQa"
}
