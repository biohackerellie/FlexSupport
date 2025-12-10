package utils

import (
	"regexp"
)

func IsMobileUA(ua string) bool {
	mobile := regexp.MustCompile(`(?i)(iphone|android|ipad|mobile)`)
	return mobile.MatchString(ua)
}
