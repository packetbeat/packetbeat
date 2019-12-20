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
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DNSNameLabelLengthTooLong struct{}

func (l *DNSNameLabelLengthTooLong) Initialize() error {
	return nil
}

func (l *DNSNameLabelLengthTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c) && util.DNSNamesExist(c)
}

func labelLengthTooLong(domain string) bool {
	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if len(label) > 63 {
			return true
		}
	}
	return false
}

func (l *DNSNameLabelLengthTooLong) Execute(c *x509.Certificate) *LintResult {
	if c.Subject.CommonName != "" && !util.CommonNameIsIP(c) {
		labelTooLong := labelLengthTooLong(c.Subject.CommonName)
		if labelTooLong {
			return &LintResult{Status: Error}
		}
	}
	for _, dns := range c.DNSNames {
		labelTooLong := labelLengthTooLong(dns)
		if labelTooLong {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dnsname_label_too_long",
		Description:   "DNSName labels MUST be less than or equal to 63 characters",
		Citation:      "RFC 1035",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &DNSNameLabelLengthTooLong{},
	})
}
