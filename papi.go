package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func provisioningApiRequest(u, p, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(u, p)

	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	clean_response := strings.Replace(string(strings.Replace(string(bodyText), " ", "", -1)), "\n", "", -1)
	logger.Debug("URL: %s || Response: %s [ code: %d ]", url, clean_response, resp.StatusCode)

	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode == 404 {
		logger.Fatal("Not found")
	}

	return bodyText, err
}

func GetAppData(username, password, url, application string) (string, PapiApplicationResponse) {
	fullApiUrl := []string{url, "/applications/", application}
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
	fullApiUrl := []string{url, "/addresses/number/", address}
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
	logger.Debug("Looking up user - %s", accountName)
	fullApiUrl := []string{apiUrl, "/users/", string(accountName)}
	bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

	var data PapiUserResponse
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		panic(err.Error())
	}

	str := string(bodyText)
	return str, data
}

func GetUsersApplications(username, password, apiUrl, accountName string) Applications {
	logger.Debug("Looking up user's applications - %s", accountName)
	fullApiUrl := []string{apiUrl, "/users/", string(accountName), "/applications"}
	bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

	var data Applications
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		panic(err.Error())
	}

	return data
}

func GetApplicationAddresses(username, password, apiUrl, application string) ApplicationAddresses {
	logger.Debug("Looking up addresses for applications - %s", application)
	fullApiUrl := []string{apiUrl, "/applications/", string(application), "/addresses"}
	bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

	var data ApplicationAddresses
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		panic(err.Error())
	}

	return data
}

func GetUserFeatures(username, password, apiUrl, accountName string) []string {
	logger.Debug("Looking up accountName - %s", accountName)

	fullApiUrl := []string{apiUrl, "/users/", accountName, "/features"}
	bodyText, err := provisioningApiRequest(username, password, strings.Join(fullApiUrl, ""))

	var data PapiFeaturesResponse
	err = json.Unmarshal(bodyText, &data)

	if err != nil {
		panic(err.Error())
	}
	logger.Debug("Features list: ", data)

	str := []string{}
	for _, v := range data {
		str = append(str, v.FeatureFlag)
	}

	//str := string(bodyText)
	return str
}
