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
	"crypto/dsa"
	"math/big"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type dsaSubgroup struct{}

func (l *dsaSubgroup) Initialize() error {
	return nil
}

func (l *dsaSubgroup) CheckApplies(c *x509.Certificate) bool {
	if c.PublicKeyAlgorithm != x509.DSA {
		return false
	}
	if _, ok := c.PublicKey.(*dsa.PublicKey); !ok {
		return false
	}
	return true
}

func (l *dsaSubgroup) Execute(c *x509.Certificate) *LintResult {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return &LintResult{Status: NA}
	}
	output := big.Int{}

	// Enforce that Y^Q == 1 mod P, e.g. that Order(Y) == Q mod P.
	output.Exp(dsaKey.Y, dsaKey.Q, dsaKey.P)
	if output.Cmp(big.NewInt(1)) == 0 {
		return &LintResult{Status: Pass}
	}
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_correct_order_in_subgroup",
		Description:   "DSA: Public key value has the unique correct representation in the field, and that the key has the correct order in the subgroup",
		Citation:      "BRs: 6.1.6",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &dsaSubgroup{},
	})
}
