package main

import (
	"strings"

	"github.com/op/go-logging"
	//tropo "bitbucket.org/voxeolabs/go-tropo-utils"
)

func testinf() {
	logging.SetLevel(logging.ERROR, "test")
	addressCleaner("cf77f59c97b9a44c950446eff97a37bc100212a34ed68a4b0115dfc79d9f3b73446b9995c9170bdec6e1a889")
	addressCleaner("9995551212@sip.tropo.com")
	addressCleaner("tel:+14074740214")
	//tropo.encodeToken("cf77f59c97b9a44c950446eff97a37bc100212a34ed68a4b0115dfc79d9f3b73446b9995c9170bdec6e1a889")
}

func addressCleaner(address string) string {
	addr := strings.TrimPrefix(strings.Split(address, "@")[0], "tel:")
	logger.Debug("Sanitized address: %s", addr)
	return addr
}

func addressType(address string) string {
	if len(address) == 88 {
		return "token"
	} else if strings.HasPrefix(address, "999") {
		return "pin"
	} else {
		return "number"
	}
}
