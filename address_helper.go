package main

import(
  "strings"
  "github.com/op/go-logging"
  //tropo "bitbucket.org/voxeolabs/go-tropo-utils"
)

func testinf(){
  logging.SetLevel(logging.ERROR, "test")
  AddressCleaner("cf77f59c97b9a44c950446eff97a37bc100212a34ed68a4b0115dfc79d9f3b73446b9995c9170bdec6e1a889")
  AddressCleaner("9995551212@sip.tropo.com")
  AddressCleaner("tel:+14074740214")
  //tropo.encodeToken("cf77f59c97b9a44c950446eff97a37bc100212a34ed68a4b0115dfc79d9f3b73446b9995c9170bdec6e1a889")
}

func AddressCleaner(address string) string {
  addr := strings.TrimPrefix(strings.Split(address,"@")[0],"tel:")
  LOGGER.Debug("Sanitized address: %s", addr)
  return addr
}

func AddressType(address string) string {
  if len(address) == 88 {
    return "token"
  }else if strings.HasPrefix(address, "999"){
    return "pin"
  }else{
    return "number"
  }
}
