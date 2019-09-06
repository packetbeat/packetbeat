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

/************************************************
The CRL distribution points extension identifies how CRL information is obtained. The extension SHOULD be non-critical, but this profile RECOMMENDS support for this extension by CAs and applications.
************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtCrlDistributionMarkedCritical struct{}

func (l *ExtCrlDistributionMarkedCritical) Initialize() error {
	return nil
}

func (l *ExtCrlDistributionMarkedCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.CrlDistOID)
}

func (l *ExtCrlDistributionMarkedCritical) Execute(cert *x509.Certificate) *LintResult {
	if e := util.GetExtFromCert(cert, util.CrlDistOID); e != nil {
		if e.Critical == false {
			return &LintResult{Status: Pass}
		} else {
			return &LintResult{Status: Warn}
		}
	}
	return &LintResult{Status: NA}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_crl_distribution_marked_critical",
		Description:   "If included, the CRL Distribution Points extension SHOULD NOT be marked critical",
		Citation:      "RFC 5280: 4.2.1.13",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &ExtCrlDistributionMarkedCritical{},
	})
}
