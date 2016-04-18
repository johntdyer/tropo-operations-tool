package main

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/wsxiaoys/terminal"
	// "strconv"
)

type sipMessage struct {
	Reason       string
	Rfc          string
	TropoHandler string
}

// func main() {
// 	//sipMessages := make(map[string]*SipMessage, 10)

// 	BuildSipTable("486", true)
// }

func buildSipTable(code string, showAll, includeRfcs bool) {
	var sipMessages = make(map[string]*sipMessage, 200)

	sipMessages["100"] = &sipMessage{
		Reason:       "Trying",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["180"] = &sipMessage{
		Reason:       "Ringing",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["181"] = &sipMessage{
		Reason:       "Call Is Being Forwarded",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["182"] = &sipMessage{
		Reason:       "Queued",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["183"] = &sipMessage{
		Reason:       "Session Progress",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["199"] = &sipMessage{
		Reason:       "Early Dialog Terminated",
		Rfc:          "RFC6228",
		TropoHandler: "n/a",
	}
	sipMessages["200"] = &sipMessage{
		Reason:       "OK",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["202"] = &sipMessage{
		Reason:       "Accepted (Deprecated)",
		Rfc:          "RFC6665",
		TropoHandler: "n/a",
	}
	sipMessages["204"] = &sipMessage{
		Reason:       "No Notification",
		Rfc:          "RFC5839",
		TropoHandler: "n/a",
	}
	sipMessages["300"] = &sipMessage{
		Reason:       "Multiple Choices",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["301"] = &sipMessage{
		Reason:       "Moved Permanently",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["302"] = &sipMessage{
		Reason:       "Moved Temporarily",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["305"] = &sipMessage{
		Reason:       "Use Proxy",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["380"] = &sipMessage{
		Reason:       "Alternative Service",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["400"] = &sipMessage{
		Reason:       "Bad Request",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["401"] = &sipMessage{
		Reason:       "Unauthorized",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["402"] = &sipMessage{
		Reason:       "Payment Required",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["403"] = &sipMessage{
		Reason:       "Forbidden",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["404"] = &sipMessage{
		Reason:       "Not Found",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["405"] = &sipMessage{
		Reason:       "Method Not Allowed",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["406"] = &sipMessage{
		Reason:       "Not Acceptable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["407"] = &sipMessage{
		Reason:       "Proxy Authentication Required",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["408"] = &sipMessage{
		Reason:       "Request Timeout",
		Rfc:          "RFC3261",
		TropoHandler: "onTimeout",
	}
	sipMessages["410"] = &sipMessage{
		Reason:       "Gone",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["412"] = &sipMessage{
		Reason:       "Conditional Request Failed",
		Rfc:          "RFC3903",
		TropoHandler: "onCallFailure",
	}
	sipMessages["413"] = &sipMessage{
		Reason:       "Request Entity Too Large",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["414"] = &sipMessage{
		Reason:       "Request-URI Too Long",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["415"] = &sipMessage{
		Reason:       "Unsupported Media Type",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["416"] = &sipMessage{
		Reason:       "Unsupported URI Scheme",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["417"] = &sipMessage{
		Reason:       "Unknown Resource-Priority",
		Rfc:          "RFC4412",
		TropoHandler: "onCallFailure",
	}
	sipMessages["420"] = &sipMessage{
		Reason:       "Bad Extension",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["421"] = &sipMessage{
		Reason:       "Extension Required",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["422"] = &sipMessage{
		Reason:       "Session Interval Too Small",
		Rfc:          "RFC4028",
		TropoHandler: "onCallFailure",
	}
	sipMessages["423"] = &sipMessage{
		Reason:       "Interval Too Brief",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["424"] = &sipMessage{
		Reason:       "Bad Location Information",
		Rfc:          "RFC6442",
		TropoHandler: "onCallFailure",
	}
	sipMessages["428"] = &sipMessage{
		Reason:       "Use Identity Header",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["429"] = &sipMessage{
		Reason:       "Provide Referrer Identity",
		Rfc:          "RFC3892",
		TropoHandler: "onCallFailure",
	}
	sipMessages["430"] = &sipMessage{
		Reason:       "Flow Failed",
		Rfc:          "RFC5626",
		TropoHandler: "onCallFailure",
	}
	sipMessages["433"] = &sipMessage{
		Reason:       "Anonymity Disallowed",
		Rfc:          "RFC5079",
		TropoHandler: "onCallFailure",
	}
	sipMessages["436"] = &sipMessage{
		Reason:       "Bad Identity-Info",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["437"] = &sipMessage{
		Reason:       "Unsupported Certificate",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["438"] = &sipMessage{
		Reason:       "Invalid Identity Header",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["439"] = &sipMessage{
		Reason:       "First Hop Lacks Outbound Support",
		Rfc:          "RFC5626",
		TropoHandler: "onCallFailure",
	}
	sipMessages["440"] = &sipMessage{
		Reason:       "Max-Breadth Exceeded",
		Rfc:          "RFC5393",
		TropoHandler: "onCallFailure",
	}
	sipMessages["469"] = &sipMessage{
		Reason:       "Bad Info Package",
		Rfc:          "RFC6086",
		TropoHandler: "onCallFailure",
	}
	sipMessages["470"] = &sipMessage{
		Reason:       "Consent Needed",
		Rfc:          "RFC5360",
		TropoHandler: "onCallFailure",
	}
	sipMessages["480"] = &sipMessage{
		Reason:       "Temporarily Unavailable",
		Rfc:          "RFC3261",
		TropoHandler: "onBusy",
	}
	sipMessages["481"] = &sipMessage{
		Reason:       "Call/Transaction Does Not Exist",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["482"] = &sipMessage{
		Reason:       "Loop Detected",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["483"] = &sipMessage{
		Reason:       "Too Many Hops",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["484"] = &sipMessage{
		Reason:       "Address Incomplete",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["485"] = &sipMessage{
		Reason:       "Ambiguous",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["486"] = &sipMessage{
		Reason:       "Busy Here",
		Rfc:          "RFC3261",
		TropoHandler: "onBusy",
	}
	sipMessages["487"] = &sipMessage{
		Reason:       "Request Terminated",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["488"] = &sipMessage{
		Reason:       "Not Acceptable Here",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["489"] = &sipMessage{
		Reason:       "Bad Event",
		Rfc:          "RFC6665",
		TropoHandler: "onCallFailure",
	}
	sipMessages["491"] = &sipMessage{
		Reason:       "Request Pending",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["493"] = &sipMessage{
		Reason:       "Undecipherable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["494"] = &sipMessage{
		Reason:       "Security Agreement Required",
		Rfc:          "RFC3329",
		TropoHandler: "onCallFailure",
	}
	sipMessages["500"] = &sipMessage{
		Reason:       "Server Internal Error",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["501"] = &sipMessage{
		Reason:       "Not Implemented",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["502"] = &sipMessage{
		Reason:       "Bad Gateway",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["503"] = &sipMessage{
		Reason:       "Service Unavailable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["504"] = &sipMessage{
		Reason:       "Server Time-out",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["505"] = &sipMessage{
		Reason:       "Version Not Supported",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["513"] = &sipMessage{
		Reason:       "Message Too Large",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["580"] = &sipMessage{
		Reason:       "Precondition Failure",
		Rfc:          "RFC3312",
		TropoHandler: "onCallFailure",
	}
	sipMessages["600"] = &sipMessage{
		Reason:       "Busy Everywhere",
		Rfc:          "RFC3261",
		TropoHandler: "onBusy",
	}
	sipMessages["603"] = &sipMessage{
		Reason:       "Decline",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["604"] = &sipMessage{
		Reason:       "Does Not Exist Anywhere",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["606"] = &sipMessage{
		Reason:       "Not Acceptable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}

	table := tablewriter.NewWriter(os.Stdout)

	if includeRfcs {
		table.SetHeader([]string{"Code", "Reason", "Handler", "RFC"})
	} else {
		table.SetHeader([]string{"Code", "Reason", "Handler"})
	}

	table.SetColWidth(500)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	if showAll == true {

		var keys []string
		for k := range sipMessages {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if includeRfcs {
				table.Append([]string{
					k,
					sipMessages[k].Reason,
					sipMessages[k].TropoHandler,
					sipMessages[k].Rfc,
				})
			} else {
				table.Append([]string{
					k,
					sipMessages[k].Reason,
					sipMessages[k].TropoHandler,
				})
			}
		}

		table.Render()
	} else {

		s, ok := sipMessages[code]
		if ok {
			if includeRfcs {
				table.Append([]string{code, s.Reason, s.TropoHandler, s.Rfc})
			} else {
				table.Append([]string{code, s.Reason, s.TropoHandler})
			}

			table.Render()
		} else {
			terminal.Stdout.Color("r").Print("Invalid or unknown sip code").Nl().Reset()

		}

	}

}
