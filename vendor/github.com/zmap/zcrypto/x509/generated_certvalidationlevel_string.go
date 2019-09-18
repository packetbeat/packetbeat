// Code generated by "stringer -type=CertValidationLevel -output=generated_certvalidationlevel_string.go"; DO NOT EDIT.

package x509

import "strconv"

const _CertValidationLevel_name = "UnknownValidationLevelDVOVEV"

var _CertValidationLevel_index = [...]uint8{0, 22, 24, 26, 28}

func (i CertValidationLevel) String() string {
	if i < 0 || i >= CertValidationLevel(len(_CertValidationLevel_index)-1) {
		return "CertValidationLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CertValidationLevel_name[_CertValidationLevel_index[i]:_CertValidationLevel_index[i+1]]
}
