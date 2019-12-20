package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/***********************************************
CAB 7.1.2.2c
With the exception of stapling, which is noted below, this extension MUST be present. It MUST NOT be
marked critical, and it MUST contain the HTTP URL of the Issuing CA’s OCSP responder (accessMethod
= 1.3.6.1.5.5.7.48.1). It SHOULD also contain the HTTP URL of the Issuing CA’s certificate
(accessMethod = 1.3.6.1.5.5.7.48.2).
************************************************/

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaOcspUrl struct{}

func (l *subCaOcspUrl) Initialize() error {
	return nil
}

func (l *subCaOcspUrl) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c) && !util.IsRootCA(c)
}

func (l *subCaOcspUrl) Execute(c *x509.Certificate) *LintResult {
	for _, url := range c.OCSPServer {
		if strings.HasPrefix(url, "http://") {
			return &LintResult{Status: Pass}
		}
	}
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_aia_does_not_contain_ocsp_url",
		Description:   "Subordinate CA certificates authorityInformationAccess extension must contain the HTTP URL of the issuing CA’s OCSP responder",
		Citation:      "BRs: 7.1.2.2",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subCaOcspUrl{},
	})
}
