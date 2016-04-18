package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
)

func renderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Property", "Value"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output
}

func buildPpidsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ppid", "Environment"})
	table.SetRowSeparator("-")
	table.SetColWidth(200)
	data := [][]string{
		[]string{"Production Voice", "461", "Shared Production"},
		[]string{"Production Messaging", "462", "Shared Production"},
		[]string{"Vegas Production Voice", "653", "Shared Production"},
		[]string{"Vegas Production Messaging", "654", "Shared Production"},
		[]string{"Multisite - Tropo", "663", "Shared Production"},
		[]string{"Multisite - Tropo w/ SMS", "664", "Shared Production"},
		[]string{"Orlando Fallback Voice", "1070", "Shared Production"},
		[]string{"Orlando Fallback Messaging", "1071", "Shared Production"},
		[]string{"Orlando Staging Voice", "410", "Staging"},
		[]string{"Orlando Staging Messaging", "445", "Staging"},
		[]string{"ORL - Tropo w/ SMS, New TropoGateway BETA", "1121", "Alpha"},
		[]string{"LAS - Tropo w/ SMS, New TropoGateway BETA", "1120", "Alpha"},
		[]string{"Alpha environment ", "502", "Alpha"},
		[]string{"Outlook Production Voice", "655", "Dedicated Production"},
		[]string{"Outlook Production Messaging", "656", "Dedicated Production"},
		[]string{"Vegas FiServ Voice", "1122", "Dedicated Production"},
		[]string{"Vegas FiServ Messaging", "1123", "Dedicated Production"},
		[]string{"Orlando FiServ Voice", "832", "Dedicated Production"},
		[]string{"Orlando FiServ Messaging", "833", "Dedicated Production"},
	}
	table.AppendBulk(data)

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output

}

// func buildUserFeaturesTable(s PapiFeaturesResponse) [][]string {
// 	data := [][]string{}
// 	for _, i := range s {
// 		// value := []string{i.Feature, i.FeatureFlag}
// 		data = append(data, []string{i.Feature, i.FeatureFlag})
// 	}
// 	return data
// }

func buildFeaturesTable() {

	data := [][]string{
		[]string{"s", "Outbound SIP"},
		[]string{"b", "SIP Bang Syntax"},
		[]string{"c", "Override Caller id"},
		[]string{"w", "International Outbound SMS"},
		[]string{"i", "International Outbound Voice"},
		[]string{"u", "Domestic Outbound Voice"},
		[]string{"d", "Domestic Outbound SMS"},
		[]string{"r", "SIP REFER"},
		[]string{"x", "Disabled Account"},
	}
	renderTable(data)
}

func buildAddressTable(data Address) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Key", "Value"})

	fields := structs.New(data).Fields()

	for _, f := range fields {
		if (f.Name() == "ServiceId" || f.Name() == "ApplicationId") && f.IsZero() {
			continue
		}
		value := ""
		switch f.Kind() {
		case reflect.Bool:
			value = strconv.FormatBool(f.Value().(bool))
		case reflect.Int, reflect.Int32, reflect.Int64:
			value = strconv.Itoa(f.Value().(int))
		case reflect.String:
			value = f.Value().(string)
		}

		table.Append([]string{
			f.Name(),
			value,
		})

	}

	table.Render()
}

func buildApplicationsTable(apps Applications) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "  Environment  ", "VoiceURL"})
	table.SetRowSeparator("-")
	table.SetRowLine(true)
	//table.SetColWidth(200)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range apps {
		ppids := fmt.Sprintf("v: %s - m: %s", v.VoiceEnvironmentID, v.MessagingEnvironmentID)
		table.Append([]string{
			v.ID,
			v.Name,
			ppids,
			strings.Join([]string{v.VoiceURL, v.MessagingURL}, "\n"),
		})

	}
	table.Render() // Send output
}

func buildApplicationAddressesTable(addresses ApplicationAddresses) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Address", "Channel", "ServiceID"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range addresses {
		address := ""
		channel := ""
		switch {
		case v.Address != "":
			address = v.Address
		case v.Number != "":
			address = v.Number
		case v.Token != "":
			address = v.Token
		}

		if v.Type == "sip" || v.Type == "number" {
			channel = "voice"
		} else {
			channel = v.Channel
		}

		table.Append([]string{v.Type, address, channel, v.ServiceID})
	}
	table.Render() // Send output
}

func buildUserTable(papi User, features []string) {
	fullName := []string{papi.FirstName, papi.LastName}

	data := [][]string{
		[]string{"Username", papi.Username},
		[]string{"AccountId", papi.ID},
		[]string{"Email", papi.Email},
		[]string{"Name", strings.Join(fullName, " ")},
		[]string{"JoinDate", papi.JoinDate},
		[]string{"Status", papi.Status},
		[]string{"PasswordFailedAttempts", strconv.Itoa(papi.PasswordFailedAttempts)},
		[]string{"Feature Flags", strings.Join(features, ",")},
	}

	if papi.Address != "" && papi.State != "" {
		data = append(data, []string{"Address", strings.Join([]string{papi.Address, papi.Address2, papi.State}, "\n")})
	}

	renderTable(data)

	if papi.Notes != "" {
		cleanedResponse := removeNewLines(papi.Notes, " ")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Notes"})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.Append([]string{cleanedResponse})
		table.Render() // Send output
	}
}

func buildApplicationTable(papi Application) {
	u := fmt.Sprintf("%s - [%s]", strconv.Itoa(papi.UserID), papi.UserData.Username)

	data := [][]string{
		[]string{"User", u},
		[]string{"AppId", papi.ID},
		[]string{"App Name", papi.Name},
		[]string{"Platform", papi.Platform},
		[]string{"Environment", papi.Environment},
		[]string{"MessagingUrl", papi.MessagingURL},
		[]string{"VoiceUrl", papi.VoiceURL},
		[]string{"Partition", papi.Partition},
	}
	renderTable(data)

	// Add notes
	if papi.UserData.Notes != "" {
		cleanedResponse := removeNewLines(papi.UserData.Notes, " ")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Notes"})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.Append([]string{cleanedResponse})
		table.Render() // Send output
	}
}
