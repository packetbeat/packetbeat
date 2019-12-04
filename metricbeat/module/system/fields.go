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

package system

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/system.
func AssetSystem() string {
	return "eJzsXf9vGzey/z1/BeGHQ517thLn2l7PPzwgTa6AgbY24gR3wMODQu2OJJ655JbkSlH/+gcOud+50q60kpWi/qF3saXhZ4bD4cxwOLwmT7C5JXqjDSQvCDHMcLglF4/4i4sXhMSgI8VSw6S4Jf/zghBC3B+JNtRkmiRgFIv0FeHsCci7h0+EipgkkEi1IZmmC7giZkkNoQpIJDmHyEBM5komxCyByBQUNUwsPIrJC0L0UiozjaSYs8UtMSqDF4Qo4EA13JIFfUHInAGP9S0CuiaCJlBhw/6YTWo/q2SW+t8EWLE/n93XPpNICkOZ0ITLiHJPLedv4j9fHbc6diQVFL8Mjb4FQQXFtaVTgWLl6RGQuVSEEs3EggOOR+ScUJJk3DD8XkWC+U9daPlPk4kqIyyu/TpnhUuxaPxhCzf2x0J/Z1GJLJmBKlHVPvlf5AFUBMLQBeggoEyDmqSRCcLSEeUQT+dc0uYH5lIl1NyS1NEfBv7jEvIv0gUK2rJjWAJEpyAMYQKBEZ3SCDp4q3FgWPSkxxGtBUcTmQlzIDCvL+co3CdQAvgQLkYU8E4JD0AnWATnJ2EpCJfr61QxqZjZkFTJCLQG3Yebk0l6X5Qs5mcoc0TVA/jpFLkHILmmzJyhLAWxwMilFCRm+ullPz5OaSOG4VO/nZ+QNagVi6xrZl26JRUxt/9YUhWvrTfHhAGlstTsXI/qt9OJfjTUWs7N1zQvFu9+HD733OyB3ADl5zczTBAmVpJnwlC1cSZgtsE4Z8WUySjHb6yXjAP+drlJrUi0VK3B1lTX5CXNElS+BUo1aX3h7YoyTmcciBR8YzfPT4J96SXIU9rF8xVQEcul2UGhXJRmrWjScmUjZn1YdGbDvDEnysVm+UQhdZIq0N77whmQ2kzch6W4Fnb9cPY7NMNEUlkZmqwZ52RJV2ADVPqFJVlCVpRnuGg+37x+/RfyVzfcZ6TdIlaOU6NLuQIab4ihT1Y/mPZUmTCS0ChCtXO2ZdUmGsBiofyhQ1NyL9opAn3VIruRGYmocJNWFXmRvFkooAaU/YVwciM/SUXgC01SDleEzcnfWmSdStmvU0O+f/0XC+3K6pVTLp/2mERpNsml+dlpzwzIzQ+dk/PHCmH/WEHi1xt+/VGina/Ia/3TLw9w+Kd3O453a6Q5U0FaXxA0cWzjjnoXc0DFubv/l7VCXU7Jr6Vn1Ms/sZ7UWYpgaJr6bBkZutGfJyMH7fbnyVL/Lf9M8e+x758nJ6Nv/l8Vm/t6AOfJ5NfqBpybNPt4AVd5IkRDnAu5zNlgcB3gveExfGxl976Wk+lzPtP9Ok5Bz/Aw8awP4Z77KGT/HfG5ke+7yf159lCVidVTJl80RTHk+MGSqJw/2H+Su/uijKxnDV7+M/yMwv43OJ9PsFlL1Tw48PnjW6JjejN8upE9O+QuZQPFKJ+6zXMAvJ4QvtF+hLzcjXxcMk0SuiFCGjIDqxwrFrttnHJeCr1F0+fodzCkgMYTPPAYcfGgp1TxMOwgVmXsDFmV0VlkNXyecb7ZgW+tmIGjA8RR9kSIEpxtTP8TtdwVDH1pD/BIBmHUYZN7QX5mIvvijrhYcyjS8AM1REYqTwkPe1LOvKYJQrXOEisZ/BTR7Hf0Q7+7edNrBp9fQBaHATGOjHJiPcXUorpbbKhWdt85otonjNuYIJIi1n5782YFV2yviX02iG7N7nQWjw0wjDGWdh+8e3W/G6CN3iY42wp+y0CbSQJqAXqagppqiILYQxHmDvDNo3pc5n5ITXBMPCUnjhN3YrsGBeS3DDKIiZG4GGJYsZ2xjWfLqchp+cIxj81Ybb5OOlEleqZ1C32Fzz0m6LQzMy4nOCOegS27zQhs/Fjut4Xv28Lc9N3DW9pOhqiNL8ZlhK5A0QVUY5q5VA0tC86IkdYDtQELxEPW/wlnxanYMafFsXS6eWksmpEmJl/xdLWYWh/lOKyg93PJhBPvSztNFnVPC9CPE7ThR+YDxyAcxMIsj8LEKZf5uIrk0hcw7fSyDlciN4JjxCpT1d16iUzdvbofdz5mmd6Mx81DOHMfZ8o6iesli5Z1Fro3xcsZFfGaxWZJMsM4+53aYVEI5adeTsh793FNTabcR2QUZTZwcTVzZcmjJhGXGqe+XsWYiwSEUTLdHJJMKtNW/jpkm+bwBBHNiU5nzIya+ivQWsJ2ytpwSxjPfxRU4vU4r6w0qWEryLUnlZIXIfu3r//xfWuW54xD7eYr2StrWJJp1S6XfxqjhLlg+kQ5BUwQ4qlORd5G2pA/E6liK8bBxhl4NpXveJMgdLdIpwMTnIOSmNWS2lvy+VUMq1f2rzefg4jsuEeAYmk0ocAX820YBCbcp6lkHZm+vbEgYWtpkXZLNmE0qK1HTBtY+kTIGLTVFrtG8TftzHkFkoJn1fbtWm3RTceWWkVeCmAfoaHcTyQ1N8cV2W2XWKbhtIljO+BAeM+/uzVAb6tTKFRR2w3moH3Mq5SjVNnKKptYfhJGFwsFC1ochVHOnclpXG4pv3rw7Z19D0N+rZsfj4bMZdaMjGvL54Bl/TFg9jr0zQ0VCOK2aX14ZtuMg8b5KbkmsYx0KxsQkDrZboG3imIX+hbOUPRAvBDRAjbWQBOgXSzPBhBX6g6AIXN8OoTO7F0i0JRnGmX6sh3ycEnjQ8yHDfEsjTyGPXDBX9xcDDXC9k9MLKZzGhmpbm1oN8wQ/1yBX4SXnGpDEiYyA+E1fPHdOSH9zmPtMDgXN2eF9iYAN4wbSxCfSycCukBiVtQk9CstbLPzXFMRVpgxOHo27erQqgN5wg+FOTrupeGWdXZdwQ5y7xyJVorC9xsbIT1xsrAD9zWHe3f/qNOFG5/sFtsL1gmjWheflZV96FH5OS9iIVdx5ZKjsQSNhVdMRDyLiw9HUrgqj9kmdycjGi1BEyra/tcsm89BaXKpoYhVvWhoZDLKJw035OzDsV4T63jbz19vI3mL1MqOgBBj3aiV3C4vfqu3HFwR5AQeqWeoIs+KDt4ZosAbQ+0y+8wqEYgIyAzMGvzNd6/SWNVQzdX4GQo2RbA/zU+SGFIQsc4t7/2jy5MlUgGJwVDG9RVJ0QySaAnRUxEjV3T4c4dKkOePoby4w0v+zuA5COVRxjGQn1E7LRVZFGVizODqZ9qdDGAOqqQZHBojjdI+5PYAid4//tuSZJpQorOkaZXyiWWCRpjPz+f1XpB/MRHLtb7y34ff2qvNi1YWc+W/3neuOmwO6WN3yE7b03PmAgceraWzq2p3TZvmptsQpQrm7MstufhfZOv/mj5PPb9hNwukUvoS1n1g2rBIu3OY8hDP4qg1Nc1VLJTB3J2OeOZgumSmryo9l61Fb2QY3ucyU+VJ6SC4MjOTtHWDuwfmGqYo94yQFEJIrR3MzG4ETBwPABO7x1dAY7rEIrCDYaDsC4KkTrAHAjT2gxNx2yAgRbKsHnR/bVY767cIcy6W2QJCM3lyw22BeH18bpsdvHbSYi08IWXuv8KQP0VLpeTnbsN/qXj+TBDKbchup6hkZwyPZgcDB62Jeh1QTa86y63ISHvSOKpThg17K5ENatRq/2OKcRjJUZBZZjDCD+nTQM50pqyr/7yMyRWoSCYJG7w0YpjTjJvQCVxvHg5Y3+/d8K5KaS7VMPB2X5lU/ZQm8NCG0UJWmfqQ71Pht8PKVyF1uSFklzBbsPogqlgJyvmMRk+jDP0ud8gqosECyyTTeB9Rp5zZ/zPHtoBrmlbhFXc5waylqiIanrL1NCo5W/+b6q3U2uMG+d/xJvG8cSx5ugupYJYDs/iYHW+C73M71UYBpywoaV6x0yC6Ok5VwoTnRKggAra7uNmFU9ETjFpYWoLxtHsK7HhIVIGkp2CYmIBSUh1HLI60vzvvEDGx6DFXp8KkQcS7ETExiZW0xvooiJiIZIL1jH7uyop3P2wPiR0ToMzMQm4HWD1lYZpQvqab9mb52sZa76laW4dfxOTHx/dkBhHNNPisp3XdFKRSmfJgs7sPQbE1uw4LB+1HnkZlP/K/sZsRjamhV9V3d66qDxo1XgMi4+5HlDPalGVKzbLgexL4asIW7qJA8VJSe0TsMzRgC+xzuORlhqQbpbkXKhOCicVFuHwj7XhjaDf77W/24T49YMA9R1zsP2L7q31GjJKYMzHyHM8zzomNSaiIry15F8QbaWddGRdiOdxX/qTVLl4TOLqiapEleCamIaWK+lUfLDpjCyEVTOlMruCWvHn97Q9BljMNao+l5Jpi7reOovW+02pdQiYW05gp7C/QPHDqMzqIVXB0OfsPtBIj7pfTAzUAxIopKezMkRVVjM64z3oEtcD1ibcmNNSQgVZ64JCfFMCPj++v3OmcM7L3j+TfYZNRb8lPxssmvnv4dK1TiNicRdU0Ylq28xmaKOxsqkYG5ZG7s7OBDkemapG3dVtrgnWt8XA7PxLaotW+BevysJqJCJz2eHvRJevdrXzJMyfHG02mvCdTzAVyWtR0ZWmMu+WdqbhQmiWMU+WPGoPD/sWOUgiyOkDMdMrppvShjExzk513mWo3FAoLt6NB4lclYVjVArM65arjWnlgolVWV5a1WSkyQxQVXSkhvDnwun0JsyniLR0NyYntQrjTYROw04lj4nUlMFund4s8rfUIXV4u0cVtp3cIOotpnT9UkQsRG7/ZoaHj3kWwnKRn8ZQ7KRm6H+3a73btV8+UNi41IO++52OsqriXdIsKKK2f7VDLov8AmsVWZx/BkEf2O0wayzDAkIyiLGXuJCyh9j/uM5cf3v7ycjur52eZx+NPL6l6LiXEseMQM5nuvDMZjgP2KI/8iXEoPiOV95DyPIPbtDR4dXK5cqYrrvQ80J3SFbzZvct72WObDJmCOGhTaFwRrctAI/3eOwFnCTMTLeeDz4b7KoicGzdKXkGwA3rhUwRJ1mKlCu2ICjIDEi2tsxE3/RxqCBUb3JV2iWJJW6HeWKKwpI8ligptKwrskzoDomje/FpJaTrCw9DC23tJ/pLfDxUejy77ErmRsCcWdv9Al5vqJ1cImkD9qev8x3+ruHiqoMx9tlyMJdWekF6yFEsmWgSFFNdWHJ4yClBDbQCUXy3kRrMwNJptZaPIjrxSDwETr0137zFUsZok8fat40YTqrWMGCaJ1swsXdsTK+awZ3+HMRF2XhHfGEJzqnfvXarC9x3MqSM15Dt/aSxIlc62HPKQ2nmxWR5PSJZ6Xobq9ajZJMT/WmczF2V8o909Ztc2YZDIcLRTCK2d0SHDjvy7JRalWSkLoqMlxBkH95g+xRai7qYV1U9FpYhfR0Gab913cvsshVGSc2/Z1rLIaBZDKX1F3v30iAbkw8cwUft3baiIHZi8gS3fkDllqiTl7UyqpLUXTArKebha3d0Sc5cOiqAqv3KQT2NRH78GtliaCfnwsQIjSFcB5T5Ca4DSYHTlUcVg/Bn0R0nZxL4+AShkf0knb7NEyYKtQFjfk8lt1Vb9qjuCBo30WK+kqYF37/NsTFN7tgLoMBd7QQgvAvvzsI/Z6KQWMidbmYzmeuInLFhYRQYXtGxhFcfBufDvaiQsUjLv64olSXJNFCwyTpXdFTtJOZF8o3M7YSTqsgItMxWBJnopMx6jXwJF5dkAmfyWSUOPL5KPjQtznYJxC5ny8AUUhJSbSVpdoyoT+fqUAvzaJJdUkxjmzLl93VKuKkfX9bmQ9DBUO7bs3gqs3VmA8tlCPAT3SRmwBq9YSIinavA6idZaT+VOY02sk0q2PB8s9taxW5Jp5oXi3O+8uOsNsUrPFsuqN7pVvMqc8Xot1mW3fDvWK9N7LFRlJioTGGqdgzAwuS3FAjReYTNMZDLTfs11EmaiEaLUF7F7nDkstZ5icretHYxji6msk/WmBsvrVpRrNDq1BWMXRd3EdBs3u7RRFMBpur2Yu826WSppDIf45EKwuqK7ZnXm7ph6bOQSmWSBN5/zn7yMeu2Oda1tz4t1zBI2XkBfljTDZjz4esV8q12qmDur1bUZcvkApgjuhX3Nf1Pi4uhbaJGfznt+unacl0wQQYWs9TH1K62Yjx0ORmie+sVMNNqSBe4VN/koKG+tF6hoyn/+9KaLn2f2pv357Gm8xuoDdRVFr/VOsCagGj/v0PdBa9wV6IzDa8FLFXwBHAskEhl3VWeH8eXP0Z4M4aU7sH05BGoKKpxhIdvLhvKfWvnQ4ZpVcNmyngXbUhCg0RI/2tCwLds307tVbOvRLBlmPf1lNp8WdpWhfxrQIxnQ4YYygWSCJ2idB8OkzwrddbI4gPFqnx9/uDfbdKa/ypb7gxlO6JfzYXoJRVaw2vhlbM7dedc5cl2mXtwmU29cYrnNq2Vt3N69fdJoCS9dmBLIV+NBT8Vzz3Tf/cFKb04Zz46fT6mf9frIBRlaFu1S3LHfZWNOX5J1q6y2/FGAl/X7M6zX52YblpD3jokypUCYuqHA/jN48d71v/RrqJPemGurENa52pVWmY3djNvCqglliyNxqLDO3hTlmSSvcE2h4WRvz5OMbYD0+pxMUHOx4YR2UrxszToaq4FG6elc/RVfIno0t+Xp/P2Wpgh2ui+dVIdL5uyNiZw35LPNQHRnCfcxHE/n6bo0521M38XSnpooPUtTUbMRdpP5+O6h6JzZfqhgCKPnahqqNqHJccBG7EiO7Wc9UUxfg53wwmrKqWUwdklpb0+jkNZ5Go3mRHYfVg33L1zC0nWgnFIhwy0cegtgRF15K6TYJDLTpQfqOqVJQXzHTA5Um2sFEQjDN9e42i5//vCpW0CcaVO7iJqkc00u9TKB5OXVUGNUE56N0k8svJ8Yh+sZjZ7K4vRSOD9/+FSwuwdXKOsT8/NgNwgceOw5WjJQVEVLFlE+daKanpdprKaNi0gsh+29p6IdQcVOONvXfXI7irj0+jylVUZkveXWSbIuz/3klvfy/XosadF9uGouaiuvO8Btrsi9JPUMZrNbUmGDGpTRHtqRYJOv8+L40T8c6bi9dhCJ/x980anbFI9rc1K6gCn2iTt5kYy1EbS4XVGPTo1iiwUoiO0ntiXAEPpAffiPVNOvgG8EuoNxcvGL/dSF+6cmS6tCory74pMBrr833+AdFiO3hb+uPTq2isDLNTGr3u7oqVF62pl2OULhGfbQs//F6jNZadXP3KU81wW5d8lllY+u3oDHZkRmlSDtUFa2XcjtxcoptkV/9GZXiKJCpxTPXYqmxS+viJDded9xHVel9dSOfDZS+7XRdU/OCS0EGZTXsNKZNU3PhtfH4thjz9nLBKxYZOjsjHb8Xyrp2IgKIY27qxBxyhKIe3GacznjTyxkwweUy/zIZVTt8/lnlczYVTJ7FMm4asJz0djmQ6PO8KCtmYNSLtvnXgzyb7TOUKncY9zdAUz3Yc0gMTF5mqrLUgB3r+7zTohSYJm/lbarkLPs78841mFDebXe1x67F7Y5izadDRcnOksSWn+7jRkOt+TB+5eP7Q8M7czoSdR6AlcuRoMmvq9g+JW5o7/21utFt0qr2AI2wmW6hbdy48QxNhKSSrd0L7AhWFjc/23+3kAs0UEoNAdIjyGSnPAwNGbMpqsVMI7uICy/y2TGxp8hR3YQkhjo+CKJ8eGWMApyZ77RZAVqQzLB2RNw7+ow426l27CUKnwbgAmiZeLv0lFONDOZN6nMkIRufBAbZi0TT0Kum8Hl4dyVjFWujSzBtVG1QRePxTfeZzOKwcrafWVDMo+obaIVZQe9edz4/tH7q++SFU2KLnduq+taktS0bucdMC5mAZnZXLup6IGAwwrCm8fe/TbtXDi6dQBhCWxENLWwZVhP90LxzhciWuLEEb8izGH58PbuPaFK0Y27VxlnIqbChHtVx0w/5cdnIy2j6nsmLmfrBtky/jE3eByhMknV9wW3YcIYenyRINl4t0jmlPHRtrLK+I7u7vFxfenJX4MIjtTLliQUm/YoukYUzt7qIEoML8bVnEpaBYlXlWYpeYxpeHLz+s231zb8ySFsg2fX5xEcEo/PO9geokslK7wRZsfdgbawTzJ6AnPI1vTx3YOnoksQbukf5ve7NsxdhrO7nbN/3GPS9f2uNs5BEHOasFYvl74I7OcOGRzfzZ6wcNOs1q+L1u43/3gzeT15M7mxXsmb169vbl+///GH27c//vP97Q/f/e3729ubYar3M77fffdAaBwr3wuMFc12qCB3D6tv7WB3D6vviw/14S2VKvzmcmB9FPy9ebMPfDvUDkwKEmngDAT+AYGMLHHP3UlE7hnoL3MbOgxwkApgf//++s3NzfXNzd+v//b9RKwn/i+TSLZeD9yB+eHjB6IgkioO9DYDD5TcPeQvGsuZodhFZcUoUbACpdtnk3cPhEv51JnQaogBDI+nKc/0VA56KKB8D2lf9rGT/HwOkU9kptfOxY0ltgi9hI8/v3+ZO7xeFnbSXAWIFICvWTepcjoDXnt54goJWGr/fYPx1cVcysmMqslCcioWE6kWkwsr34vqL1pJ6aKJvaURgwGVMJF3KrfkSSQT8F0BqSCQzCCOISaRTDeF405Nqw0AfmFpTHr76lWazTiLdDafsy+Io7cuT/Flm/Fc/n9acv5Ds5xN1/6hmBPUQK9uxBdS7kDc/WjH+M99bAXg26HvCWLAUxHbUYz9QsdPldc5SI30VhzwZd/HZ+ALRBkedx0iD7zeP1glwt8aPvC+D4bMM86nA1Sh7gN3p88f8e8k8PdDs+dy7tro5v4zK3Pmvhn9QR50u2XY3v1W36IeC+E86uYk7GpQioUK7YAvd6xcJ6/A33uczlhgKMNudNXeu9pAIIk/IpZiCHR+wtGriUbrg2uHtBHY/nOzo+NCt0BGfnZ4tqmFkvkZ+FXZDrPM6pTvybo6GSwfca1EU/d4ye8wIe+kUqBTbIxiZN4PQgPmnV9Zi/lKb/QrAeYVS1ffvjJROk0gmZD7jra83cfw4eZ8B3dK3T27pOfhvFTpkm6vw+qe6Z5oEXH5yjy2oMJhIbYqn09tt3y3ctBlQ8ZmILcnu+Xez64cAZ+Fts3ONOGBth4B0+Fn80cGWOapKsMOkmbEpYbpmnZe7T0K2gZCayOmJZJp8H2+Om7DkvOAXQAJoS5crXjUTejT+z/IJmQZecZNKIvPcRPaPruk5yZ0ahPehXrL/ylWR9poFTvY1//sSHyu3xlqPp/qB/IRwYEJc9/+cJIMfgA6/2rjz0ykmZnmH0oY58w3BRtmID4ugdw/5rxi/9SS1OTF/wcAAP//y3p6mg=="
}
