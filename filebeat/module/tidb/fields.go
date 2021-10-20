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

package tidb

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "tidb", asset.ModuleFieldsPri, AssetTidb); err != nil {
		panic(err)
	}
}

// AssetTidb returns asset data.
// This is the base64 encoded zlib format compressed contents of module/tidb.
func AssetTidb() string {
	return "eJzEWkuP3DYSvvtXFHLxzGKmJ0BOmcNik3ECGE6AZOO1sSeBIqsloilSZlH9yK9fFCmp1Wq1x+6RvDrNtMSqrx6sF3kPGzw8QtAqfwUQdDD4CN+9129+/u4VgEKSXtdBO/sI/3wFAMCv4HenGoOvAKh0PmTS2bUuHiH4hn9cazSKHuPn92BFhT0DfsKhxkcovGvq9pcJNvw8uapytiUH2kIocch/1X455DfkmTt16H/s+Abch8GPF1jz8xOshTZuiz5xWJ28fXI2CG0JGlsLT6igQiJRIIGwqsW0gr9K1xgFuNcUwFlzgF2JFniJtkXkgAPKHXTjijPkQ409A70H1+pu7V111J62FISVyFxoKNVYkUNEr//x+uT3DtUGDzvn1ejdZ7Dx85srOmgK19qi6swr6tpoKXjhCt6XGLmAW8NWmIZ16xGE2YkDAQWvbbE6I/5f14AUFhpCEFAYlwsD0tkt+nBThlDT48PDbrdboREUtFxJ91A0WuED2occRaCHtTbIfz1UggL6h3b1qgyVuYXaO4lEzkNwQDVKvT5EnOd2JON2i9qSGXyVHYOucNKQSgTMrLCOvs6WbKTaaxsiaTaVcSOrdLw/NegP2WUErsnN+NUXcMc9yoY/6BEIoCACVmjDNBJXB13pv3E+MB3lndAhC2MdvoymbapMujoLgjbzUo6ha36LdOEt2mPNG6XE50wiXVVrswAY1fgYT9gxGEf0ws4DUqS57K2TSEY55Atw/AR//fnbQAPwvtTtjgZNYF1IewgVCALBUbWLiN0OPw90bwOv5W84kIHRFjsZpxf1Vve4zShUYyGuka6jqXSBNE3wqhTBlltrW6BPwaUV7FSNk0j49fQu+WogfdDc24yC8Bd39tUiRqpxn1AQVR1rB5azsfpTg/D2TQpnwQtLQp47K4wSXmC/irkyJicUXpb9DhxQufdoBLvbOHcchdaUsUd6K8ykzLlzBoX9Opk/lhhK9CBOLcmO3NYmieO5lO99g6Ct4uoA2e1FmKKSkkEsKBIlE3cSEz8jGQs1Yeir6OaHZCBCf0FvVuE+02pmT3n7hrpNEDkgl8RbZ7apeHo25VEj5RJ2nNYSc0OidWPMYRpPLuTGrdfzx3vOwH3yyXHtPILH4NvSe4gYrXQN+wkBeu98a36PnxrNhWZaN+GLJXK+4u6kXSeIRS5BW2kaFSt4uQEnZePpDv6NBScgqo0Od9Ht3ut3H4DQc3ehCfKGLqiJ6bzDA81fqdTmzOgvdNFjIca0+1w07Zf8/CdGLIzdYaZQOoVZXLpubAxT3daF+3sgNCjD2bc3r/f7/Wq1en07MOz9PYe/WN0kDLFK13IE8UJ6NMJmc+azQd6thcfxyhm2YQz7JztRwB+JGzgPv6RtGR0buRH1XHJ8RnpuMjIpZDntcnNiLXVIRczIeyL35yDmHLhtMQ/IjnwpKKucx8wjNeZCxr+aused12GBOjfGO+ksNRWqmPMTqxgLj3XvemSAiyUit7kZNTmv0zhP39FhtU2Vo2cwRwZw01W8HbTbNht7HKZ1EGorrMTzcPKr84B7UdUG7yKlXYkeQcdVN2380Ou2hb7teB+g0kUZIP8CNpf18y0M2m6R1qA9/DObPgufCWUxuWUxA16HexJ1SrWxon3O2dKSYSPUUBMrNha2RiJdaQpaDivXOAnclVqWp87Cy5m3RpUSfWyMOAszg9Gc7YIuFhxNJLWc9qKf18yJN+uToQaHKMgR7dFhQ+kRW7XzPnJBGLhJcvcNwFp7ChB2Ln14e3fG9JdTZUCFwqbwnCiOm+kjxV5Umq6Xxitjz3rUjyb4sx8PQaVtQ2MwF3fiQgF1Wta6FIRw07G97d6GnbtP7wbOGqtEfSHGxlFRDz/X1rhi/jIvIVheO4TSWdWpJ3F9gXIKDFkHfanYSjXvJWeZWYyqPUOWrg+z8wrW8pi9/Tnrqw71haR9df3qkbjfzDiqzg/cOClMZkSQZZamqIsYPZYUaXZeo1XE9u9axi5OxrTRto4LuEDa8Bs8zFtUxUzOqJhyknM09UmciXsjfvOR/4OnX1OX9e7D59CS/ntmY8R8wmQ7RUbYzndHPcFxRcbcA9pjSj9X8oX5VR9Zfey+F6pfYyOf+vvTgcycGYOfX4Qsu0FC8Loo0FOcUFQuYDqTUo1HkMKYaYWEvU3ZdOb8Us8fCyqsskrsZ7VZJfa6aiqm7fwBqBYyjvHUMNgPmlD02qk0fj2dBt60A9pYcuaHgLfT+laaNotJwcS/gQwNoZ93SMR0ux3A1GFXdsVjHL8+35vyqn+Vbu7ThpcDk87aTE+Pd66G9eSsxRQa3r7hLpYo/X27Om0RDifHANhxZFSPP0ycBhhXsIjE6a2jyd7ww7Rw8ym83x/5vJqSjffs20oEkQu60C+k8Tz/PXt1FIdrL2mon802T+54+aDl1x9ZHDc3cV1zIQdw1p2/9DgiPNYeQ6zcr5IU1l5qwjfbLEKbNYvUagGifX29GGWFQehp0lc7Z1vEfTKZR6qdJVxC38neCxTusVZsyfdnOxNpbbKQ5ednlIIDI8cGDnLEXweXyjfpbBs7zOFueISa5oK4l4hqMJ74TDu/kOgnZ1qnF1yelznELHLcjF3V6hvLFaTRlQ6oBns4lB6FovPpUB9ykmI+NdggNHV7sGZMt3IcqoRH2Dm/mbyv9TGdynVnE23pnGgHscGI0XXCB9f5wTm8MD7960YTTU5M3YajBNpKj4LwYvOQXPn/Einb/8axEtKdyHiEBDsdykE0vzsuiv+DcpjulbQnklERzijYoqfYrLg1/P7h6emcx09QeBQBlF6v0aOVCDmGHaIdM4m3FXoM41P0StjDKct4DfJS7VRncaIuttPHOdcqXGzRi2Lqlph09X28UvUMovrH72dF9MeP378AzVINxfWIhFIz9whMEam/79ABaZ0+1q62wJP58eW4zDhjbF7Ks8Yh5wv0lka/C/jVtViW8qpr8XxDn4ot0AnO7lZdcrJppN7JDcWbCAa5stroukY1Y/E/ZrTBwzfgksfZbjx7z0o9ZzMzzYnLhW/DJD+E6yuz/wUAAP//Mhy8Dg=="
}
