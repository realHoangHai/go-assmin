package util

import (
	"fmt"
	"regexp"
)

// ValidatePhone validates phone number
/*
MobiFone
089	089
090	090
093	093
0120	070
0121	079
0122	077
0126	076
0128	078
*/
/*
VinaPhone
088	088
091	091
094	094
0123	083
0124	084
0125	085
0127	081
0129	082
*/
func ValidatePhone(val interface{}) error {
	phone := val.(string)
	r09 := "^[0]{1}[9]{1}[0-9]{8}"
	r08 := "^[0]{1}[8]{1}[1-9]{1}[0-9]{7}"
	r03 := "^[0]{1}[3]{1}[2-9]{1}[0-9]{8}"
	r05 := "^[0]{1}[5]{1}[568]{1}[0-9]{8}"
	r07 := "^[0]{1}[7]{1}[06789]{1}[0-9]{8}"
	r := regexp.MustCompile(fmt.Sprintf("(%s|%s|%s|%s|%s)$", r09, r08, r03, r05, r07))
	if r.MatchString(phone) {
		return nil
	}
	return fmt.Errorf("invalid phone number")
}
