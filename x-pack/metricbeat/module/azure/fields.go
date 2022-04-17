// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/menderesk/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded zlib format compressed contents of module/azure.
func AssetAzure() string {
	return "eJzkmc9u67YSxvd5ikGWAU4ewIsL5LZdnEXRon/WxJgaK2wkkoccOsd9+oKiJMuSLMu27CaoFwZiSt/vG3JmJDJf4I12K8C/g6MHAFZc0AoeX+Lfjw8AGXnplGVl9Ar+9wAA6VooTRaKeIujgtDTCnJ8ANgoKjK/qi78AhpL2ovHD+9svNSZYOtfRgiHMl0pViXlDpVuRxrJN9q9G5d1fh8VTp8/XgleUhjETskR3YboyJvgJA2A3Rhm4Bod8Jak2ijqWu2HexDyztLBwPGIT9horMTbwWyAO7ZG0fF7WXRUnIXuz+4C7HYFhtrtXGPun59GsWb9F0nuDaUfxZSxziWiRGuVzuvrH58ezwsiZWwbRmV2kK/x21scSdizK6SVAk8FSe6kbEPzYd1KCJVdz+wKwtcfB0C0tlASF+N19MZwmSpJe2X0YVocSYkT6TA3FSYsHzStjrmB8XSJf3462/emMHhk8FLXPycz4IiD05QN7aK1Qmmv8lf2Jztt+8RZE+M8B91lHuFM9V/P6FhkyOOtcGRgRi+qRIf3NkzS2dJE0tlx3ni2wBnNr581sGjriwHUHkdTx/NhZDfJmz7kv540jr4F8uyFNEHzsw/lKHksM06gf0vKUAlPszeoCspuAU/KEx6CJ9cEH7T6FsZn/gILf3pyE2BPvur6N2H/nsRPxo2BX0lzLA7KFvbw0tVOvAk/a2fePTnBqlQ690ITvxv3JrLgqsJ9xm2+lLP/JxbULKhZ0LBmGfRVpd7Bna/q9xxrjiSpLYk7mKtR5/mzzsiYoDq/ywTucefZZMNY3MVhRZo2R98lVXo36NQ/NdoT9dnh1/N0QweA3KwGFLSl4pQlT257c0cJMmHIktsYV6KWlBYpttiSSuN2AreoClwXJNY7Jr9kNv26x0KDhYSFFgsVdr7tumiEdWqLfC/XNRVq6qWmpQ3CkpOkGXNKhXw379IG2MNTbV8bwr9k/mzbylQrFp0LT9Lo7G7Ov/6SsiX6g8Qe7DLWqiiUzm+xx6ilAXV814kLnxGjKuZuUGVwjrTcjc7Vxadlo6rtQjpi/C6k8f11uHiFfuhrtUcwZNFxSZrF8ueRe/HhYec+WJOFwe73SnAtOjzd3b9lx/5TbfkW3NCl/DqxkUxo0v2wrgdP7ifrOogdQJns8HQPrp/xpsySPPTk29MEWbWHG6RaW+aJcDzfUHLAYtHiqk46K1kYyDbYjXEk0fPi4Eb4ODol3MJnF3Ur7d7bdkxT2sAktmXvxOvwf1mHLu5+arp3ORGA8BIL8sSfI5LWbjs+EptmVJqcUNpzfGJ/7NBqt9C4nQjIUa48u93nCKhxOxFQ3Neoz7JAtdlBOBkyrtGTqDvzR46m8do8RQbBlEYrNm7+e2o+8y11KNx/Jz1e0p6Ni1uRDzyxtcWj81qPf4Yc6YfyTwAAAP//nCIVNA=="
}
