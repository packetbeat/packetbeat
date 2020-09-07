// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "eJzUV8tu2zoQ3fsrBl4GSD7AiwsEt5suuuteGJNjhY1EEuQorfv1hR6kKVHyo1aKxIsAEcnzEM8MqUd4peMO8HfjaAPAiivawfa5/X+7AZDkhVOWldE7+G8DAP1cqI1sqnaJo4rQ0w5K3AAcFFXS77qJj6CxphN4++Ojbac609jhyQzDGCaFYlVT6VDpOBIgX+n40ziZPJ8F7n/fXwieexvETokZ3MDoyJvGCcoIUw9X0AUc8JaEOihKpU7tjiwfLY0Glh1fkBGktMvBHIATWbPUU4srcMfXkGNHw1j6p4dZWrP/QYInQ/3D4pywZEpRo7VKl8P87cP2NhN9bKKNTmwWmvavtziTmptjGqHAU0WCk9wENt/sI0Sh5P2cKSB8/ZIRSlWT9sro8T4t7NGF/bl2b85oHpVyIi4T3k/xTw836z5UBhcG/1b1t14MOOLGaZK5XLS2UNqr8oX9xf4T+/CeGK9TgNZWSmC3zTM857pSsnQcObi/R6SykvSl9J7RcSGR51vjzMAVvB1ovjZwkpZrM5KWy3zzYYUbmuE0tLBqK2wNDBqz5O5VVSldvkdoB2hALaHxWBJIYlTVtckVjXOkxXHdzM6iBkrriPFXIYyf7sbyPl3g/H+KFXszWXRck+aifbCuyxM4ZOAns0Y2WS7vJB5A88tQYO2SUHQVvGJ99vm60Bd6atLzXfAe4rPtYaiDwpJTRq7ehEOZ9fAwgY/HgBCmeZeoxTLvGZbzhoIbrFYtru4Q6mAhgw20B+NIoOfViQPwMnUfuJWPoqGVpmtjxzS1bZiKt3pyFo0//cYq/vl16qTyjIHCC6zIE38OJ1FuHJ/xphmVJtfeFBm1oI9tbVALQe0ZQ45K5dkdP4ehoPaMIU/uTX2WDRrE5h9/yLhHT8XQmT+ym6A1nCL5B6HRio27/p5aXnlLzYGnd9LlkvZsHJYfOiaDxPhe/wQAAP//ZxBMvQ=="
}
