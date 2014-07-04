package main

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
  "strings"
)

func provisioningApiRequest(u, p, url string) ( []byte, error ) {
  client := &http.Client{}
  req, err := http.NewRequest("GET",url, nil)
  req.SetBasicAuth(u, p)

  resp, err := client.Do(req)
  if err != nil{
    logger.Fatal(err)
  }

  bodyText, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err.Error())
  }
  return bodyText, err
}

func GetAppData(username, password, url, application string) (string, PapiApplicationResponse) {
  fullApiUrl := []string{url, "/applications/", application};
  bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

  var data PapiApplicationResponse
  err = json.Unmarshal(bodyText, &data)

  if err != nil {
    panic(err.Error())
  }

  str := string(bodyText)
  return str, data
}

func GetAddressData(username, password, url, address string) (string, PapiAddressResponse) {
  fullApiUrl := []string{url, "/addresses/number/", address};
  bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

  var data PapiAddressResponse
  err = json.Unmarshal(bodyText, &data)

  if err != nil {
    panic(err.Error())
  }

  str := string(bodyText)
  return str, data
}

func GetUserData(username, password, apiUrl, accountName string) (string, PapiUserResponse) {

  fullApiUrl := []string{apiUrl, "/users/", accountName};
  bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

  var data PapiUserResponse
  err = json.Unmarshal(bodyText, &data)

  if err != nil {
    panic(err.Error())
  }

  str := string(bodyText)
  return str, data
}

