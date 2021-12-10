package utils

import "strings"

func ParseMsisdn(msisdn string) string {
	components := strings.Split(msisdn, "@")

	if len(components) > 1 {
		msisdn = components[0]
	}

	suffix := "@s.whatsapp.net"

	if len(strings.SplitN(msisdn, "-", 2)) == 2 {
		suffix = "@g.us"
	}

	return msisdn + suffix
}
