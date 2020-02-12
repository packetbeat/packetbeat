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

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IssuerDNTrailingSpace struct{}

func (l *IssuerDNTrailingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNTrailingSpace) Execute(c *x509.Certificate) *LintResult {
	_, trailing, err := util.CheckRDNSequenceWhiteSpace(c.RawIssuer)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if trailing {
		return &LintResult{Status: Warn}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_trailing_whitespace",
		Description:   "AttributeValue in issuer RelativeDistinguishedName sequence SHOULD NOT have trailing whitespace",
		Citation:      "AWSLabs certlint",
		Source:        AWSLabs,
		EffectiveDate: util.ZeroDate,
		Lint:          &IssuerDNTrailingSpace{},
	})
}
