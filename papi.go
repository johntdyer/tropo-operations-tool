package main

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
  "log"
  "strings"
  "fmt"
)

func GetAppData(username, passwd, url, application string) (string, PapiApplicationResponse) {

  fullApiUrl := []string{url, "/applications/", application};
  client := &http.Client{}
  req, err := http.NewRequest("GET",strings.Join(fullApiUrl, ""), nil)
  req.SetBasicAuth(username, passwd)

  resp, err := client.Do(req)
  if err != nil{ log.Fatal(err) }

  bodyText, err := ioutil.ReadAll(resp.Body)
  if err != nil { panic(err.Error()) }

  var data PapiApplicationResponse
  err = json.Unmarshal(bodyText, &data)
  if err != nil {  panic(err.Error()) }

  str := string(bodyText)
  return str, data
}

func GetAddressData(username, passwd, url, address string) (string, PapiAddressResponse) {

  fullApiUrl := []string{url, "/addresses/number/", address};
  client := &http.Client{}
  req, err := http.NewRequest("GET",strings.Join(fullApiUrl, ""), nil)
  req.SetBasicAuth(username, passwd)

  resp, err := client.Do(req)
  if err != nil{ log.Fatal(err) }

  bodyText, err := ioutil.ReadAll(resp.Body)
  if err != nil { panic(err.Error()) }

  var data PapiAddressResponse
  err = json.Unmarshal(bodyText, &data)
  if err != nil {  panic(err.Error()) }

  str := string(bodyText)
  fmt.Println(bodyText)
  return str, data
}

func GetUserData(username, passwd, apiUrl, accountName string) (string, PapiUserResponse) {

  fullApiUrl := []string{apiUrl, "/users/", accountName};

  client := &http.Client{}
  req, err := http.NewRequest("GET",strings.Join(fullApiUrl, ""), nil)
  req.SetBasicAuth(username, passwd)

  resp, err := client.Do(req)
  if err != nil{ log.Fatal(err) }

  bodyText, err := ioutil.ReadAll(resp.Body)
  if err != nil { panic(err.Error()) }

  var data PapiUserResponse
  err = json.Unmarshal(bodyText, &data)
  if err != nil {  panic(err.Error()) }

  str := string(bodyText)
  return str, data
}

