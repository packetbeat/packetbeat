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

package auditd

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "auditd", asset.ModuleFieldsPri, AssetAuditd); err != nil {
		panic(err)
	}
}

// AssetAuditd returns asset data.
// This is the base64 encoded gzipped contents of module/auditd.
func AssetAuditd() string {
	return "eJzMXVuPHLdyftevIPwiG7DWJ8dBHvRwAB07ARZxEiO2ggRBMOKS1T3UdpMtkj2zc359ULx0s28zXbNyED8svLPDj7eqry4sUu/YM1zeM95L5eUbxrzyDbxnH/LvEpywqvPK6Pfs9yM4YNwC80dglYJGOlaDBss9SPZ0CZ9HLNYa2Tfw8IalL75/84axd0zzFt6z3oF9wxhj/tLBe1Zb03fh9/xd/P/85QAYPll+fzbCMO6AzpSujG05fv6QvlqCj/BKpg8y+jNczsaOn2718PjzwwwKf94Jhk0f3kzmXakG3MV5aPdN/p+G7/+hKzDv5vXLMEeMazHFLGe9thLbkOFba4sxX47VBdmawK3+xlXZWJd7gVcEhffDgCMmbxR36ZOO+2NUuIegSA/Dd1tVWx578baHCeROxF1YsBMMqgqEVyfYB1u5nbijHu0D3ovr+AnkPsh6H2LY4Z2LuhNyXFQCuNsJHleAAFztRS72jACPPw8t7/bx5V8GnXus2CcLzjQnOCjpPjHlmAPPvAl9MKWDZRNGV6ruY/f4iWafcpefBrCzahr8qudKM85a3nVK18xUqFSRcsJAXYA/Aks9yyk/fAsP9UNQbfbuL8wa47+7zuIFC6wt7goTTPpbW9c1PriJTUAFAuwoyIQOSo642UMhc4QuKD1EfSGA1/uxo5pQFp8APmcRyvoQuin5hLTLhD4WzHKzo2Em0Cjdv+wjl9+PwH77x1+wAVMStFf+ghwQPGThjb2uzIODzPa5UFwI02vPXP/UKo/eeGUs4z2SlFciTGnWhTUNyUvDIb11w6SK5hlRmpYrTcH8Pa9GARxRmLEBYO4HNnCC5kYX8MLbDsMY96f93QbgeW+Ce6iNvbx2ThkHZyVM23HrW9DePZRBUWeNABeFdi5dkx5+jV9k3HurnnoPLo57LXgS51v2NnX7cDb2Wen6IJUFHP7lisFdmbLorQXtWYJhA8xkjs70VsDtKf4Wvsf8kXvmraprsCCD8sAJtN+eL87pzfZezYatHJp6hMV2jDtnhArB7Fnh76zX6oU5I57BT+YhwXmlR7W6Opmfxy8zLqXFvft/O7MhFbAvOHfwpQctYDKyxuj6useFEpObMt23T2AzO4ZlYNzhmFWtx8TCM1gNzQP7bdolS+1dyEw4b3A9sTnrlfY//jl7b7E541oywTXyV2NOkHh4nI5zI1XeXOf5hEJj9vjzOHZvGGeNqZV+YB+aJs7OMQtN2IzxzwNSRgl+5JGfYrbF8RbYiTc9TAdswfWNvzLekQx7EVjDWFZx1WzpcQQcTBVKA0hmOkhu77cJ5wcE+W4Wibq+bfnAHEsTueG5Ik/MKPZKbD8QbJbwkDRYUagw+FKtlkMo2VcVQ9/m+k0JyFKQgDas/kQ1s5kwVtVK82YmCfjf488P7NFHYdDGM3Hkuo5KwlQ1Tj9+HmIKro0/gs0+wTIP4UAYLe+YbBTy1PjmBC+dErxpLpP5DLLMEzvF1fqewYuAzge35YyR1TCzI3f4P5J9cv2nuYE2T59B+P2yU25XuRP+GGI0mwDZE+DvXKDs953RmUR2ChOO43WS9KH8Y17hc5DxI7Bvwni/wdHHOBG92u8TnX8/AcI1fJc45bulLNwp9t988/WkasDKSEdzpjhciz2Nf30Ch0hZNAOlcsc6sJWxLcgH9tH1KJ+4+YMgwEu5cdFda9FmBH5BgCga8AKiDwy+acgX1tvtCx1+US7wb2iysNwLKdzI6WojSe59aJBM6dyvhxMFScJJoVVPWl6ZXkvUnx9GnEJ9D9RQZyEuCEINZlZB6PHLKsyeEOUmyoRB7sJY5iWuAiCDMHPWYHNafx4vEsUABTWJQrQUlQLLvnUdCMWbeLrBjG4u3806wp/U6T8rLVFn4iwGZyXqqoUKLHqMcr5Gy7zB3jXKWf652o2nNbsQz0cljqEVElAerjBWLgbbEhUav8+qhte4yoyHwa+sNHn+Qa+VZvwk3JTjJPd8H8WhI4HfZpU1beGrteAcr3M8e8VfHH30XQPX4CvVeLCs42gimVSuM06tpEVapRfO6B66C+3W+ZOLhYtyPZmTcy05q1Ps04ApJWmUMdHcGg+zEDS7WcoxYbQG4VEGcV9mPQrVHWlMHSTFVEzYS+dNAmAOGkC/aq45JE2UvY0OWlyg5P3OIEF7q8CRhjyEoqlxdvxGCfL8aaFISD+UbhzYE3ZjmWgUhrtK51UayGvehVN/o1EiXBi2CS71DXDX0Zbf4ZhT4mhJgtyKI1U2oalCO+VB+N4m5loA14IKPO4ot3Ufsm4xSAou3AmYuzgMU+Y0wD/fRwPYbp0Geq1INBCz1LJsN8QXhQwSCTA0XWMUD7bFIJS0wKlNVPWSS2yvNWpoZ01teYs2aNZfbbn2xpLUs+NtKixxjHedNSfsY/Tw545ryEESbcXQqKCuLb0x85KEm+I4902GcCTSsXI55TLfHU9KQnt/YX0SSdq2TNVhr6Bik6xqSrN4XjQG8jGnZ0RIEs+nVrgN+xLsv/8X8/Ay14mKt6ohLVLhFVjjjTBLEiDRTRKWf/nwE+NNbazyx3bL3HUVTfDBhhWtjD1zKzHYtiAurAV/NAtT6qEloU95MiSMoxPqJnmPORv/6VUBCv+71zX/8+ua//iq5kfjPNV7xnXM7ajOV9MZSzIejRFIy+DPxj6zovUQzlERk/eIzdbtG7yQ7VugDWzHxBjdDApNItekfOj1VErXQanVQmYbqt88XcfkOi+8p7s2Z3slBe/4k2oUjfMxoHkZ26qFP6y5rV/BC1P/KSUb2YrvpOH8DjT6GSSnUsM5cHtK20cA5sCjUszVF79zeOLiuTH1oVEtTfRiF9HBeutYwmFfeuiBFY524UjQXQhjL2t+luDdgZbcyI72UIVRCAgby53KRAHNfoPobcDClskxCSl3DB1OasEdGs6HjqSduLd5Gh06jaFG4eo0TCPvkaLOQohO9kqSWambutpBPCF46+JhzUayzDQ0zEbGw70ZzBPXGmgucmoS98/oyIIgWcfrBcECx1CLFlSGQ9TUMolKjL7nXsipfSf8C5VuTqiPwmh08JjzdmXDiLHqsF+b8aoDFEnSKufCjiDL4WzFFXqzwJ+c0pI6+PAfPzEJQoUj4RAygfxBglaLXpB17bBa5Ey9qZjVdT74khK135qWcXaaOyOomRpI08laGXzjlH8qz8gXfThVw+mAXzAkKlN1CEfXrKqiQV05EDGNPICujBWKtuCo57gEsTGgNWXOc9/PTTUu8Ul0/T1rPFrsn379yISxC0fAor5SoIfaL5aKvwqAohKB5k5erTOYlhnMrb+URNdiWBI0RPHoQYJfC8wqTkrgJy8oy/QkCB00s9HP74is1SkZh+sbpZ9z2tqBlgtpdP3TZ5L36bouNIq0eJVs+X//6d2P/0Ml8bmnuJphE2tXOm5kf/wREUW+0oHI1Vy4o4miG7W3jp3ABp5dV3phWpJomHGF4xFyo/Rq5g2Zm8inJY8myr7GptlMnUhaP9SvhjTAiFEeRQ272XBH1f9Y4xdaXh9+KJkkrY9yh9jo4Ll7nlZJDZquKIjh7GsYpdJHsOqmDzv1jIhUlRpHupr7+I5kd2Id4KUxXE4YN6V55tHOV08NoIQ7oJ70I1dnGSzajvb8xBslD4nA7pHsadNh/sS8X5p+oZPzkXYvVAV//PU/h6zDujdDDFtUJ0ZGWg9aWi5CIpQCC/6I7OmzhcIm7PHneHI7Z9A77FSqx7pqpBTRtj7+GnLOteUtqyyvgxs21iisyC4tWRuoIl3aK/3pNWo7tVTPIERIm7FMyC6QiC1D7aQ09EypIXNn4aRM7+J1prVA1zia9zhI8lhzO2d3emVH/mTrAE8Ry0UmKrdRMILSJZV7JoVtyj3fFKyO2zFlv9tIlEFPthQRaa1+pAHSyVoDuh5q2IedrxpOi6E60EPueO2UuCdKJ2cfPy71iFht40Cgw5ErDVfS2SEKqC0xbRbd/3T1dyNXfKBlyoOs70zCoXi2tJImtNq8DWUZpmIttMZe0H/857/O45aQdrnHbI9Zl6QGEoSSId816+OuKB1nsCtKx9URR06tjEN8bMaFB5uTMDu8eKTer5XhQZ/ldC3H03tVkdSy90xpD7biYiNtIlqSXuawaVpwO8M8GkPizvHMF1vGY7jkughcpJWzN9xlck0n7rHt9eS+2iD4VPsZ0jrX0s1OkaTiWpZserJy5sofvKIddG6criAWK7AGSxUy1KTIIDVZC6ZRSYiMlRNnu1iLylnxqZTMuCVprWSKTSOJ3ptpJNWDQyq7Yz8lhzZYtvISfAyscTYl6uCDhSc36F01pn7rpq3LoyeiG4aqmL2vGc3OwP0doWS4PcJtDT5c5CmzPldClpaTQqHbxWVn7mlFd9Na3th+Ix0QOr/T6sywiP7P+OTJDOfJGBodJwOYPQhsD3zuLSjRdgdPLD0PEa+pQuNcwLwitI0RREN1jpFSKGtP520FyIQyRqO11wjGshBrWuVEjzxRuOFj3TJpjXm8XxjqztHXyAm+qz6NovrDZR1GqEqY2Y/JGTWJSssz6p1UKqHifePf3cEbqWlwNNfTUCEWo5LdQHQpeAsgyExnvZrZQCPfcefOZEqNwlkZG47aA4axErm7XvJd2A1ixeqZWviQS+3u2Y7s82fWLoV2eil4EnidaH5LmRkqC3WunzwUkTeN2hebQFySqGzbec2iNGTfQUu6pYPtet6ov+26pRNSWrSDInqlCflINDuNw3sY60eilphJTpx51U0JJ/3U8ZbJ9LWxmqfPr+GayUXc0bMnZpMLw7RWj4sb9Ypk2STmXViiGPhUXDXEmphZuJMQ1g7JlKbF1ErfCqlPJLuf9I61XBw3ziLLhNouzLbl3dV0XMh7EU/Vh2PvJAebZRivzsxflwnsgZiZHnLdDml9JdXt4AuxpmX6TsZ8w/5PTNCXp4unrcOY/HbsS8+HlwJKoHFJqEeD84KUkBYp7tlNjDTdeA7Wf6fngYxMLddBQr5dqoMSSE4+mUZuJp9CvECypmWKYadFldQy63SE+vjrhqcRTDTxjvfCQi9frFJG+IaYloyPpHzpwfmh0CU/5RfwNupdHPFUORaJbztfvOu4bWmX1HKbdCGneBR1zip/vM6gKB1aTrPm80watp8/wZCTGWt3VohkHnOum2eMgV3oVSR7T1qfPh+IieKUHd50yCxwR7sNV6SwmARtQq6PJ6AQ/W3evmuUI2fS5r5UTNUj0nYlJjWzknlhT3blqaddH05pJtc/ZS+jd+F4pXTtG6NrxwbPeMLMJL0rmZlgqahEOgkdliT6B2WfekvSrI///sg6o3QQ0FB2uJ4Xim7+HRcLStF0b126VvCDVC7cqN0s470vwzIV0p1ZluiuUn2puZ1cKbXqaIhNPqV1RvPiGtPAwnfm2XcWrXpyMeKYwE/lRgigqvWHM4kOd3l6cqvGVAPp5aMoHDGQcRBeEZje6RiGTLvWHd70VyLhr71AtHUhWrX13bc+QqSuWl7DjQsgjbzn1D2L+eLk/daZ+4G4evklzMG+XD90SIXulB7Ox1B2F3Uihd5n7jJU1TfMWKbNekb5fp/iVkI5BdMkYzBEQBIa8BvlrrHKfoY7rR/ffNevrGXdGtKKi18Us668oDeJbwiolp/nVwact3148mLZzd29TB56WeLOrspvIw+PR/ZavVzrMc0pArNv8evfM9Wd/j78/Ifvc0Zn7QW68V1TwiRp75uW/eU3ht6U/RXWaNLR+LbcB82MlSHAaNIDbT5taEZkFgSoIZtUPDqKEcqAdAYL8XqeN8h0UQDiK3TSiBBQpmcU42vxyg3kparwcom1RdyXn1XIrmB+EiM8sGQstvmktGh6CQfLz4c03PzK/YAzeeV+umZnbrXS9XTNpnu0sWwoHLn18t+K+Stwn9/ySX3H1SgePZw9uTj+IzTJ8wpg+TSNaxn+NtxelXCCxnQhSMc/Snjqx1qZrredcekdsskzuDWYdDQ5J5vVieI0Q5P8z+KEAUKldH6MVhh9Aq1CIk9pJrgDdjF9Kl0bowHQVonjuIG9ixFXRg8BkdLsF1M7z90Rd/hR1+A8+1cjYfl+cFnTiN4xaH/QxKcLpk8sjvVkac8i6uLxbOUvX7cn5S/zTizUyuiv2k2EXMzG9Nrby0E5c6AWh5a9/RRx2ONv/xaqRBcPnJuJ0znIH5hDCG/2zQdDTOV7CUHoG+7DLw9v/jcAAP//Dsvwtg=="
}
