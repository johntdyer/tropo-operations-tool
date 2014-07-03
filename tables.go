package main

import (
  "os"
  "strings"
  "strconv"
  "github.com/olekukonko/tablewriter"
)

func BuildAddressTable(papi PapiAddressResponse){

  data := [][]string{
    []string{"Type",                papi.Type},
    []string{"Prefix",              papi.Prefix},
    []string{"Number",              papi.Number},
    []string{"DisplayNumber",       papi.DisplayNumber},
    []string{"ServiceId",           papi.ServiceId},
    []string{"City",                papi.City},
    []string{"State",               papi.State},
    []string{"Country",             papi.Country},
    []string{"ProviderName",        papi.ProviderName},
    []string{"SmsEnabled",          strconv.FormatBool(papi.SmsEnabled)},
    []string{"ExcludeFromBilling",  strconv.FormatBool(papi.ExcludeFromBilling)},
    []string{"SmsRateLimit",        strconv.Itoa(papi.SmsRateLimit)},
    []string{"ExchangeId",          strconv.Itoa(papi.ExchangeId)},
    []string{"ApplicationId",       strconv.Itoa(papi.ApplicationId)},
    []string{"RequireVerification", strconv.FormatBool(papi.RequireVerification)},

  }

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Property", "Value"})

  for _, v := range data {
    table.Append(v)
  }
  table.SetAlignment(tablewriter.ALIGN_LEFT)

  table.Render() // Send output
}


func BuildUserTable(papi PapiUserResponse){
  fullName := []string{papi.FirstName, papi.LastName};
  address :=  []string{papi.Address, papi.Address2, papi.State}
  notes := ""

  if papi.Notes == "" {
    notes = "none"
  }else{
    notes = papi.Notes
  }

  data := [][]string{
    []string{"Username",                papi.Username},
    []string{"AccountId",               papi.Id},
    []string{"Email",                   papi.Email},
    []string{"Name",                    strings.Join(fullName, " ")},
    []string{"Address",                 strings.Join(address, "\n")},
    []string{"JoinDate",                papi.JoinDate},
    []string{"Status",                  papi.Status},
    []string{"Notes",                   notes},
    []string{"PasswordFailedAttempts",  strconv.Itoa(papi.PasswordFailedAttempts)},
  }

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Property", "Value"})

  for _, v := range data {
    table.Append(v)
  }
  table.SetAlignment(tablewriter.ALIGN_LEFT)

  table.Render() // Send output
}


func BuildApplicationTable(papi PapiApplicationResponse){
  data := [][]string{
    []string{"AppId",         papi.Id},
    []string{"UserId",        strconv.Itoa(papi.UserId)},
    []string{"App Name",      papi.Name},
    []string{"Platform",      papi.Platform},
    []string{"Environment",   papi.Environment},
    []string{"MessagingUrl",  papi.MessagingUrl},
    []string{"VoiceUrl",      papi.VoiceUrl},
    []string{"Partition",     papi.Partition},
  }
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Property", "Value"})

  for _, v := range data {
    table.Append(v)
  }
  table.SetAlignment(tablewriter.ALIGN_LEFT)
  table.Render() // Send output
}

