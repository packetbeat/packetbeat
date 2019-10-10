// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package tomcat

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "tomcat", asset.ModuleFieldsPri, AssetTomcat); err != nil {
		panic(err)
	}
}

// AssetTomcat returns asset data.
// This is the base64 encoded gzipped contents of module/tomcat.
func AssetTomcat() string {
	return "eJzUV91u8kYQvecp5i6tFHgApFaquKiaigpFSdu7amwPeOvdHXd3TKBPX61/wBhkG+J+Unxp5vic+R/mkNFxCcImRpkBiBJNS3h6K188zQAcaUJPS4hIcAaQkI+dykWxXcKPMwCAyhgMJ4WmGcBWkU78svxtDhYNtRjCI8eclrBzXOT1mzakDYsxTun09hYyPBeaViiolUVYBSwYEqdiD1vHBiQl+IOiV/JcuJhemaX1ma6rzdMV1xZoIkJ78UtHzjoYgKQoIKnyQHuyAsoHNhRKQLgDr3zM6PjBLrlJmipZCAvqPuK3lMAWJiIHvAVH/xTkxcOWHbjaf1/p+iBH4MntKTmHqRv5szLNdndTllf/UqVrkUVD0uLCuRAJ8qIMCgWRJ97yU6AsZEpzdBTyDygxeBihw+BBmcJATs4oCQkpuS/UPC5EM2dFPk2y7mMW0QvTxVxxijI0F55rtaeSrXKYrDhFPjhulNbKU8w2GaHg1BZk2B3va9x1ibls2L9xjwuNdgcv6z+/ZKsS5ovC444WMZuqwvo0rBqjElnHEcoP3Jf/FrHBQ6/beJiQrPD9Dr77SX1TVkkf3S9WiUL9WUaWlNxjebRs5xOyj0jmxIyjMjox59i0PkjbUDaD9r5B9dqM59xxTN6zuxxaP2uOUNdWm5PNF51gg7vrt+u9VQfmqm4GKqBcsAtHMal9f839ZLiwEjgTFIQG8hw21gNruiL2ZHuLrkMazB8lrAOk7G5oSYfol2sahBtYeZrcqN0RxOQcOz98kJyTWiHG95SkjjBRF/wjmuqtgV1f7M1F/+Rrqw2znqCdosIfe2d4faIGu9qtlqqzlPsyMGKAT8ZVH9ljnDxfnzX7MwhmZD8vIiPK/8JwYY64g8tyDwgoEcB2Mm5liAsZ6Ldfz9Q14LMSvKATSoZ9bzJRA05VgAIvv69B0540fHd5Fi8D+Q+n3vn+PmmFJ1fGZSAo755cNYM6fwimlRPnxRg1q837NxAzslRbSQpF8j8lKifM+rRsCLPrDp6ocv4LAAD//87tJzI="
}
