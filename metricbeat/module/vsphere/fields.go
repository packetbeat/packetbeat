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
<<<<<<< HEAD
	return "eJzsl8+OmzAQxu95itHedx+AQ6Vqq20vaSttt9fIgSG4MQyyh6zo01e2IeKP002KaS/1ASnYfN8PzHxD7uGIbQInUxeocQPAkhUmcHd6dmfuNgAZmlTLmiVVCbzbAAB0s1BS1ih7mUaFwmACB7EByCWqzCRu6T1UosShhR3c1naxpqbuzgRcxkJDsUywMExnubDkRdluKiAyvo9+TDGGKPY4muhJjti+ks4mc7/hseNDzzTX7Q1zY/XjWT5JhaY1jCXMhHvPVNQildw+MLFQD/uW0QQJFFWH2+y/WUVwikA5cIHBjbEjJ10KTmBuP+PMNWJUzCeNGJ2yMZhFpXwxmK1DWaccZDSpUJjtckViuuAa1hp1ihVfS9stP8/OgqEgw0syYXL9v46DT2T4chKkdeM3pyx+xnyBHr++gKxgO1Ed2vociOfrU+AKY1fY8XxdWb9hW2JJeq1i3Tpxax+SfrtKO7i1cjkS3jpxvBSuQn4lfdzZX2GwPyraz14WxrLn7yCpuRGqFGkhq0XfLxeVbk8tG3sPcnqfS4NrIjgys4eVsrJvJpMsXyujv/ttgK3fh8txTRHfsanplxq1YFkd4Nl/yv1vGus1jfcnIZXYq5s6x6FBw6v1D8rhozVYnNSO1VXoiqg+HaK0vPiPddT4oj1X9wrGhx32weWsaWOYyp3vFEFC2v/A2T8Bf3K3IM4enTEEjP9qr/4VAAD//9YITTs="
=======
	return "eJzUXEuP2zgSvvevKMxlEyDx3PuwwGyCnQTYTgZxT18Dmipb3KZILVmy4fz6BSnqYVmSZZtyd/vQCPyo+updLJbyEZ5xfw9bm6do8A6ABEm8h9+2S//Ob3cACVpuRE5Cq3v45x0AQPgUMp0U0v3MoERm8R427A5gLVAm9t5/9SMolmGbhXvRPndfNrrIwzs9XA4JtYkljJglXZPrJzlINnzUQ+RQjurVhdGGwlnOuKD9Ym0QF6s9oT34XgVMarXpfDCCzb3+bRDBEwS9BkqxQbzofHmtTcboHo7ZH8EkTUxGxfnoKMYHWlhMouL822IyE8ycUy9Iy5nE5Odaatb9wgmwf6HhqIht0IGtgdZ8wfMdgJ2XP+4FvrYOWi/aZ9zvtEnO9FIh0e4tYeapLHqZptrSgutC9avpfFt+K7IVGqcaR9qOcHX/7Pegi8T9j7Dk2DIpvQs5Hp7dAAah81j+uyRtnD98/f07fNKKjJbANhuDG0aYwFeVF/T794LyguB7joY5QtZ5EiyRa5X0I3R/4+nnc+2pjm4/R4MsbmT/YOTDxBEWauPDBdZGZ9dGuUcqGaHiVebMYoH+Y4venCxzceFTksjQgQLmGYOujNgjCwgFmZBSWG/aAe+zxKiI6P1LT+9E+qwrfjZTyD89DIi7zeYO96eHsWDfGUFxe4DKsx3l2rNJX+vXJdCXcGzP+ULPPmoDXfK9pgP84pK3UKW+ajRVc4tqK4xWGSpaXNci5kXZJmTpr5i9zKe//nbKevjyq98dHd9g22iMy2ZvCmffD8dj7LvhU3ybkJg5DyQHZW4gHTRo5smDNX0LWtXdyAAWYZ8XraaVbeImqse0HfA2tCp1t8oLY1CR3MMKXR7jWtki8/04aFPBR0UO3dnpzAuX4FZwrHIaK5PQ7FlNEBB7RuuSMtdZLpEQmILlp+VX90bGVFKmlTzdW8GZhBLotALuJZspT38RmxQtQSAPWyYLBMaNttY7umNu/UHDmanudqcDj97oVUZQdQyU5znfLjWlRNhnQMZTsD1972SHil/Nh+A7ToSqLuxx8Mc/6S+LzOcdh8+r3Dl3WdENc4K0QDs38bzPR59hps1cA5UHT/x6dHPNUWLhm2d8cjW6uKfNL9Xpe+Boi7TT5vln5GbgW0l2rPQHznO3IWo6knlakEB9QgNS4VgxlexEQumCDFM2c5lvniLhUhIwgl0qeFqenHbMQostJIVx3YhDLhSh2TK5gMdUuIqSG7SoyPpPa9TV2bcS5/wAONKDQY5ie1slVDxfhwZmq1RVfLRt7mpWLX5Ztqp01quNi+XLGX9Gsgd+Pk8UBk5tMceDsIJWu968uGo2k0ChMdrcUm07QSmUXM9XYUB7I022oV6mVR9qMUGWbYsagzoNYVZIEpxZuoHpa16Xx04Dd2bbH2M9z/Itvc5u/COs0yAmRuc5JjcwfLsSVOqsuJ8FdWaj1xXqKpC3i/ZReLGvIR5TBO0aHCkhRSYpDSx8b1xPKVwZPzHMbTAWOYmBo8ll0zCv+0D2EFcYnPgMeS7Ml77imO1SZysMFUxCxngq1NBIs+Pwd10E3eF/PaxfIbXH9SOwwunuo/GzxKR9P9BWydjAn3GO1oqV7PenldYSmTpPU19VIrjvVncpUorGdfVOOAKtqsmcBa7VWmwKgwmQhtzorUgQyDX0vQ17PbH3v1vEjtRj1M6mYV8gZRYSJOS+La+hl1cxwtpiZJugF2LX+kN2gs5AoivwKaEnCO5e3xzp6rjk9xScjKqU18/6hqwCB5Y5DjcYC7mp6A5XKKZCm2eQ41Q1erCcq5T8icrVkaqMBAjfxiDE3Z953OddyeEdLjaLDxWKd0tiKmEmef8BPgtLRqwKwuSpTJi5NvR+KFe/uSg5qgKvMmAuQdmslFhdGI651nKO8vUj0AfH4OzL7RPX2NH8qbyAPL4dhgh2+hezgvur4hyNF19x9IlDWBLcfui9Qm6jSzBjqu+6/nXA89ekErPuil8bX59RYMQwbfr9cp+UfaL8cHRb7HRRhUUl/QeglFEZS08PrrMJUrsQ+wDMAmeSF2WLttrD5x/LY4VBp3AIPmDT2LKVzA4Ea1st7MJ00kh/Ei+vc+IG3ixesynQUu8Q95DH9Rr+03EKioGChBS/ygzXCaLeK6papME5blcwfwK5hVz+NmtWsW6VQW6hrcMcUqqt0U+0/HGOp4QccyvpQ5YJHjM90cTyotyILRtYkTgp6wQ5nYyuvXUxMGxiYWFjmArN15GwZdujtPpoU+aOxWkTZX3iTxO9pPUWJG+JW/8qaGKFtENUx/u0Z2lix/zg7w2oojqG/8N61GBzxo+WZ6dLvmJSaq1mlt2KX/X5MHCExIgtGqeDM2N7mmB6iyZF1neee7ulocmP7yoNvS99ZyekhBXWi27H3sNWeosgyJZNTkPxf4UwPtFGaDOqlciblY9GQ5XBe6JrcHGTNJjCOeDTw3VVhOssN2iH1pVOSj0xjhppG4aVIRsh23upTw+XBtM8MzInhWrNyUz7pP92L14iTqleYPITbilGbAE3HqOdCbE2Q/mz8KuTk6kRRIOUzn9ywh+9RFcj1y4MioFrSs/N/Ym/pNiO25Eduah8w4AYHkIVG96T1BFTRJdreBJRbWBZ3voMw5j7OZn0xZ6TGeMc+TmZP7ZMSLaSeJJ5e1F4eHgTY11Yr8PMJs5W8+BAJhLW8FhYlAXx+Io9WBOPp1nvhvHRtnfuI4DlhSWd/SxrRi9EvfovHj2NX77584q09skzDsXqtW2b9/1PFvAm7vua5+iYtZoL1r5DO2wkXvLCr/Wk3Zk4hxdW4E0YqF7wf9X2uRTlXOek74dnJKaSg3PSpIuf6AtpzSpaWPNqVtF6nrG++38AAAD//wcc5Jg="
>>>>>>> 1f77e6e69c ([vSphere][resourcepool] Add support for new metrics in resourcepool metricset (#40456))
}
