// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package microsoft

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "microsoft", asset.ModuleFieldsPri, AssetMicrosoft); err != nil {
		panic(err)
	}
}

// AssetMicrosoft returns asset data.
// This is the base64 encoded gzipped contents of module/microsoft.
func AssetMicrosoft() string {
	return "eJy0lkFu6zgMhvc9BXeZAaY9gBcDBFN0kEXb4MVZF4xE20JlyZDoGLn9A5WkcYKgTl9VbQxT1s+PpCT6Ht5pV0BrVPDRV3wHwIYtFfA8MmmKKpiOjXcF/HsHAKdpePa6t3QHUBmyOhZp+h4ctnQuLIN3HRVQB993B8sVbRlPSQyq4NuRr3m5PHwxdjZ2qKkipym8IXcfk9f8yghkCSMVsCHGkV1Thb3lt+SkgAptpLPpq8gpLSkZUPkAxtUU2bh6xP94oJNAHkYLL6MZR2Qx8rrTyFSals4+Ocb1TrvBB30x9wmmjLIhEFVAp4FNS/CXcbAu//sbuCFAS4FhwJgAoE8E+gR9lTVQ9HZL+qdJjYOhMapJpJGR+wi+uuBWDbqaNLCH2a8D2GwiAOOU0eR4cYn4TfzFQRcWj0fQuYBO4mxlD9Uo2vmZRuICFshKjSVhfwS4YuTMZVd9CJI2qTEdM3cGPgGIMZrakS59PrDXwVE4228TEPsdmg9g1ZEylaGYCMY5ujgHD7D0MZqNJdii7SkCBipgtnbvzg9u9g/MXmiQx8Itg68DxThLx+zmE6OspLgyKlUje4x72a8G9SQ39tJHw2ZLYihDf3qfiEgTU2iN+5mADkU7c3JTeC+e51s0Fjc2hTTvWB7PaAcMybIi1QfDuyWF6J0jOzaW+2YkprUb0DHpla/4uPaVGwpTmeEmEPITtsbuXjDvFS/KUCXpqSazQfW/dPK8COtIYf+HcHkRprJ8zkRbud4VPWjfonF5yR6TZvL0HTTTzbWWE56PbLEE3GtKJ0j3hfTmr5Ihasl+zgZ36rO91PVbdEr53nH+3ZYKmouSHMsp33WZO7AsF8ijoxtxJJ5lkN+pDm3e1H3Iwo0J/B0AAP//QjN3xA=="
}
