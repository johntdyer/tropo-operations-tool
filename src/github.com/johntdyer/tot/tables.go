package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

func renderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Property", "Value"})
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output
}

func BuildPpidsTable() {
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
		[]string{"WDC Production", "1124", "WDC Production"},
		[]string{"WDC Staging", "1125", "WDC Staging"},
	}
	table.AppendBulk(data)

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render() // Send output

}

func BuildFeaturesTable() {
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

func BuildAddressTable(papi PapiAddressResponse) {

	data := [][]string{
		[]string{"Type", papi.Type},
		[]string{"Prefix", papi.Prefix},
		[]string{"Number", papi.Number},
		[]string{"DisplayNumber", papi.DisplayNumber},
		[]string{"ServiceId", papi.ServiceId},
		[]string{"City", papi.City},
		[]string{"Country", papi.Country},
		[]string{"ProviderName", papi.ProviderName},
		[]string{"SmsEnabled", strconv.FormatBool(papi.SmsEnabled)},
		[]string{"ExcludeFromBilling", strconv.FormatBool(papi.ExcludeFromBilling)},
		[]string{"SmsRateLimit", strconv.Itoa(papi.SmsRateLimit)},
		[]string{"ExchangeId", strconv.Itoa(papi.ExchangeId)},
		[]string{"ApplicationId", strconv.Itoa(papi.ApplicationId)},
		[]string{"RequireVerification", strconv.FormatBool(papi.RequireVerification)},
	}

	if papi.OwnerId == 0 {
		data = append(data, []string{"OwnerId", "none"})
	}

	if papi.State != "" {
		data = append(data, []string{"State", papi.State})
	}

	if papi.Owner != "" {
		data = append(data, []string{"Owner", papi.Owner})
	}

	renderTable(data)
}

func BuildApplicationsTable(apps Applications) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "  Environment  ", "VoiceURL"})
	table.SetRowSeparator("-")
	table.SetRowLine(true)
	//table.SetColWidth(200)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range apps {
		ppids := fmt.Sprintf("v: %s - m: %s", v.VoiceEnvironmentId, v.MessagingEnvironmentId)
		table.Append([]string{
			v.Id,
			v.Name,
			ppids,
			strings.Join([]string{v.VoiceUrl, v.MessagingUrl}, "\n"),
		})

	}
	table.Render() // Send output
}

func BuildApplicationAddressesTable(addresses ApplicationAddresses) {

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

		table.Append([]string{v.Type, address, channel, v.ServiceId})
	}
	table.Render() // Send output
}

func BuildUserTable(papi PapiUserResponse, features []string) {
	fullName := []string{papi.FirstName, papi.LastName}
	address := []string{papi.Address, papi.Address2, papi.State}

	data := [][]string{
		[]string{"Username", papi.Username},
		[]string{"AccountId", papi.Id},
		[]string{"Email", papi.Email},
		[]string{"Name", strings.Join(fullName, " ")},
		[]string{"Address", strings.Join(address, "\n")},
		[]string{"JoinDate", papi.JoinDate},
		[]string{"Status", papi.Status},
		[]string{"PasswordFailedAttempts", strconv.Itoa(papi.PasswordFailedAttempts)},
		[]string{"Feature Flags", strings.Join(features, ",")},
	}

	renderTable(data)

	if papi.Notes != "" {
		clean_response := RemoveNewLines(papi.Notes, " ") //, "\r", " ", -1), "\n", " ", -1)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Notes"})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.Append([]string{clean_response})
		table.Render() // Send output
	}
}

func BuildApplicationTable(papi PapiApplicationResponse) {
	data := [][]string{
		[]string{"AppId", papi.Id},
		[]string{"UserId", strconv.Itoa(papi.UserId)},
		[]string{"App Name", papi.Name},
		[]string{"Platform", papi.Platform},
		[]string{"Environment", papi.Environment},
		[]string{"MessagingUrl", papi.MessagingUrl},
		[]string{"VoiceUrl", papi.VoiceUrl},
		[]string{"Partition", papi.Partition},
	}

	renderTable(data)
}