package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"strings"
)

func fullAPIURL(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tropoAppConfig.API.InsecureSkipVerify},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(tropoAppConfig.Credentials.Username, tropoAppConfig.Credentials.Password)

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	cleanedResponse := strings.Replace(string(strings.Replace(string(bodyText), " ", "", -1)), "\n", "", -1)

	logger.Debug("URL: %s || Response: %s [ code: %d ]", url, cleanedResponse, resp.StatusCode)

	if err != nil {
		r, _ := "provisioningApiRequest PAPI Response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	if resp.StatusCode == 404 {
		logger.Fatal("Not found")
	}

	if resp.StatusCode == 401 {
		logger.Fatal("Authentication error")
	}
	return bodyText, err
}

func provisioningAPIPost(url string, postData []byte) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tropoAppConfig.API.InsecureSkipVerify},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	req.SetBasicAuth(tropoAppConfig.Credentials.Username, tropoAppConfig.Credentials.Password)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	cleanedResponse := strings.Replace(string(strings.Replace(string(bodyText), " ", "", -1)), "\n", "", -1)
	str := fmt.Sprintf("URL: %s || Response: %s [ code: %d ]", url, cleanedResponse, resp.StatusCode)
	logger.Debug(str) //"URL: %s || Response: %s [ code: %d ]", url, clean_response, resp.StatusCode)

	if err != nil {
		r, _ := "provisioningApiPost response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	if resp.StatusCode == 404 {
		logger.Fatal("Not found")
	}

	if resp.StatusCode == 401 {
		logger.Fatal("Authentication error")
	}
	return bodyText, err
}

func provisioningAPIGet(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tropoAppConfig.API.InsecureSkipVerify},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(tropoAppConfig.Credentials.Username, tropoAppConfig.Credentials.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	cleanedResponse := strings.Replace(string(strings.Replace(string(bodyText), " ", "", -1)), "\n", "", -1)
	str := fmt.Sprintf("URL: %s || Response: %s [ code: %d ]", url, cleanedResponse, resp.StatusCode)
	logger.Debug(str) //"URL: %s || Response: %s [ code: %d ]", url, clean_response, resp.StatusCode)

	if err != nil {
		r, _ := "provisioningApiGET response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	if resp.StatusCode == 404 {
		logger.Fatal("Not found")
	}

	if resp.StatusCode == 401 {
		logger.Fatal("Authentication error")
	}
	return bodyText, err
}

func provisioningAPIDelete(u, p, url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tropoAppConfig.API.InsecureSkipVerify},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(tropoAppConfig.Credentials.Username, tropoAppConfig.Credentials.Password)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	cleanedResponse := strings.Replace(string(strings.Replace(string(bodyText), " ", "", -1)), "\n", "", -1)
	str := fmt.Sprintf("URL: %s || Response: %s [ code: %d ]", url, cleanedResponse, resp.StatusCode)
	logger.Debug(str) //"URL: %s || Response: %s [ code: %d ]", url, clean_response, resp.StatusCode)

	if err != nil {
		r, _ := "provisioningApiPost response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	if resp.StatusCode == 404 {
		logger.Fatal("Not found")
	}

	if resp.StatusCode == 401 {
		logger.Fatal("Authentication error")
	}
	return bodyText, err
}

// curl -v -H "Content-Type: application/json" -X POST -d '{"feature":"3","featureFlag":"i"}' api.admin:x@localhost:8080/rest/v1/users/pengyongqian/features
func getAppData(applicationID string) (string, Application) {
	fullAPIURL := []string{tropoAppConfig.API.Protocol, tropoAppConfig.API.URL, "/applications/", applicationID}
	bodyText, err := provisioningAPIGet(strings.Join(fullAPIURL, ""))

	var data Application
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		r, _ := "GetAppData PAPI Response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}
	_, user := getUserData(strconv.Itoa(data.UserID))
	data.UserData = user
	str := string(bodyText)
	return str, data
}

func getAddressData(address string) (string, Address) {

	fullAPIURL := []string{tropoAppConfig.API.URL, "/addresses/", addressType(address), "/", address}
	bodyText, err := provisioningAPIGet(strings.Join(fullAPIURL, ""))

	var data Address
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		r, _ := "GetAddressData PAPI Response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	str := string(bodyText)
	return str, data
}

func getUserData(accountName string) (string, User) {
	logger.Debug("Looking up user - %s", accountName)
	fullAPIURL := []string{tropoAppConfig.API.URL, "/users/", string(accountName)}
	bodyText, err := provisioningAPIGet(strings.Join(fullAPIURL, ""))
	logger.Debug("PAPI Response: %s", bodyText)
	var data User
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		r, _ := "GetUserData PAPI Response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	str := string(bodyText)
	return str, data
}

func getUsersApplications(accountName string) Applications {
	logger.Debug("Looking up user's applications - %s", accountName)
	fullAPIURL := []string{tropoAppConfig.API.URL, "/users/", string(accountName), "/applications"}
	bodyText, err := provisioningAPIGet(strings.Join(fullAPIURL, ""))

	var data Applications
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		r, _ := "GetUsersApplications PAPI Response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	return data
}

func getApplicationAddresses(application string) ApplicationAddresses {
	logger.Debug("Looking up addresses for applications - %s", application)
	fullAPIURL := []string{tropoAppConfig.API.URL, "/applications/", string(application), "/addresses"}
	bodyText, err := provisioningAPIGet(strings.Join(fullAPIURL, ""))

	var data ApplicationAddresses
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		r, _ := "GetApplicationAddresses PAPI Response: %s", bodyText
		logger.Error(r)
		panic(err.Error())
	}

	return data
}

func getUserFeatures(accountName string) []string {
	logger.Debugf("Looking up accountName - %s", accountName)

	fullAPIURL := []string{tropoAppConfig.API.URL, "/users/", accountName, "/features"}

	bodyText, err := provisioningAPIGet(strings.Join(fullAPIURL, ""))

	var data PapiFeaturesResponse
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		panic(err.Error())
	}

	logger.Debugf("Features list: %+v ", data)

	str := []string{}
	for _, v := range data {
		str = append(str, v.FeatureFlag)
	}

	//str := string(bodyText)
	return str
}
