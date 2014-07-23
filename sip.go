package main

import (
	"github.com/olekukonko/tablewriter"
	"github.com/wsxiaoys/terminal"
	"os"
	"sort"
	// "strconv"
)

type SipMessage struct {
	Reason       string
	Rfc          string
	TropoHandler string
}

// func main() {
// 	//sipMessages := make(map[string]*SipMessage, 10)

// 	BuildSipTable("486", true)
// }

func BuildSipTable(code string, showAll, includeRfcs bool) {
	var sipMessages = make(map[string]*SipMessage, 100)

	sipMessages["100"] = &SipMessage{
		Reason:       "Trying",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["180"] = &SipMessage{
		Reason:       "Ringing",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["181"] = &SipMessage{
		Reason:       "Call Is Being Forwarded",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["182"] = &SipMessage{
		Reason:       "Queued",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["183"] = &SipMessage{
		Reason:       "Session Progress",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["199"] = &SipMessage{
		Reason:       "Early Dialog Terminated",
		Rfc:          "RFC6228",
		TropoHandler: "n/a",
	}
	sipMessages["200"] = &SipMessage{
		Reason:       "OK",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["202"] = &SipMessage{
		Reason:       "Accepted (Deprecated)",
		Rfc:          "RFC6665",
		TropoHandler: "n/a",
	}
	sipMessages["204"] = &SipMessage{
		Reason:       "No Notification",
		Rfc:          "RFC5839",
		TropoHandler: "n/a",
	}
	sipMessages["300"] = &SipMessage{
		Reason:       "Multiple Choices",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["301"] = &SipMessage{
		Reason:       "Moved Permanently",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["302"] = &SipMessage{
		Reason:       "Moved Temporarily",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["305"] = &SipMessage{
		Reason:       "Use Proxy",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["380"] = &SipMessage{
		Reason:       "Alternative Service",
		Rfc:          "RFC3261",
		TropoHandler: "n/a",
	}
	sipMessages["400"] = &SipMessage{
		Reason:       "Bad Request",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["401"] = &SipMessage{
		Reason:       "Unauthorized",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["402"] = &SipMessage{
		Reason:       "Payment Required",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["403"] = &SipMessage{
		Reason:       "Forbidden",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["404"] = &SipMessage{
		Reason:       "Not Found",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["405"] = &SipMessage{
		Reason:       "Method Not Allowed",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["406"] = &SipMessage{
		Reason:       "Not Acceptable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["407"] = &SipMessage{
		Reason:       "Proxy Authentication Required",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["408"] = &SipMessage{
		Reason:       "Request Timeout",
		Rfc:          "RFC3261",
		TropoHandler: "onTimeout",
	}
	sipMessages["410"] = &SipMessage{
		Reason:       "Gone",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["412"] = &SipMessage{
		Reason:       "Conditional Request Failed",
		Rfc:          "RFC3903",
		TropoHandler: "onCallFailure",
	}
	sipMessages["413"] = &SipMessage{
		Reason:       "Request Entity Too Large",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["414"] = &SipMessage{
		Reason:       "Request-URI Too Long",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["415"] = &SipMessage{
		Reason:       "Unsupported Media Type",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["416"] = &SipMessage{
		Reason:       "Unsupported URI Scheme",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["417"] = &SipMessage{
		Reason:       "Unknown Resource-Priority",
		Rfc:          "RFC4412",
		TropoHandler: "onCallFailure",
	}
	sipMessages["420"] = &SipMessage{
		Reason:       "Bad Extension",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["421"] = &SipMessage{
		Reason:       "Extension Required",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["422"] = &SipMessage{
		Reason:       "Session Interval Too Small",
		Rfc:          "RFC4028",
		TropoHandler: "onCallFailure",
	}
	sipMessages["423"] = &SipMessage{
		Reason:       "Interval Too Brief",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["424"] = &SipMessage{
		Reason:       "Bad Location Information",
		Rfc:          "RFC6442",
		TropoHandler: "onCallFailure",
	}
	sipMessages["428"] = &SipMessage{
		Reason:       "Use Identity Header",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["429"] = &SipMessage{
		Reason:       "Provide Referrer Identity",
		Rfc:          "RFC3892",
		TropoHandler: "onCallFailure",
	}
	sipMessages["430"] = &SipMessage{
		Reason:       "Flow Failed",
		Rfc:          "RFC5626",
		TropoHandler: "onCallFailure",
	}
	sipMessages["433"] = &SipMessage{
		Reason:       "Anonymity Disallowed",
		Rfc:          "RFC5079",
		TropoHandler: "onCallFailure",
	}
	sipMessages["436"] = &SipMessage{
		Reason:       "Bad Identity-Info",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["437"] = &SipMessage{
		Reason:       "Unsupported Certificate",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["438"] = &SipMessage{
		Reason:       "Invalid Identity Header",
		Rfc:          "RFC4474",
		TropoHandler: "onCallFailure",
	}
	sipMessages["439"] = &SipMessage{
		Reason:       "First Hop Lacks Outbound Support",
		Rfc:          "RFC5626",
		TropoHandler: "onCallFailure",
	}
	sipMessages["440"] = &SipMessage{
		Reason:       "Max-Breadth Exceeded",
		Rfc:          "RFC5393",
		TropoHandler: "onCallFailure",
	}
	sipMessages["469"] = &SipMessage{
		Reason:       "Bad Info Package",
		Rfc:          "RFC6086",
		TropoHandler: "onCallFailure",
	}
	sipMessages["470"] = &SipMessage{
		Reason:       "Consent Needed",
		Rfc:          "RFC5360",
		TropoHandler: "onCallFailure",
	}
	sipMessages["480"] = &SipMessage{
		Reason:       "Temporarily Unavailable",
		Rfc:          "RFC3261",
		TropoHandler: "onBusy",
	}
	sipMessages["481"] = &SipMessage{
		Reason:       "Call/Transaction Does Not Exist",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["482"] = &SipMessage{
		Reason:       "Loop Detected",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["483"] = &SipMessage{
		Reason:       "Too Many Hops",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["484"] = &SipMessage{
		Reason:       "Address Incomplete",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["485"] = &SipMessage{
		Reason:       "Ambiguous",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["486"] = &SipMessage{
		Reason:       "Busy Here",
		Rfc:          "RFC3261",
		TropoHandler: "onBusy",
	}
	sipMessages["487"] = &SipMessage{
		Reason:       "Request Terminated",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["488"] = &SipMessage{
		Reason:       "Not Acceptable Here",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["489"] = &SipMessage{
		Reason:       "Bad Event",
		Rfc:          "RFC6665",
		TropoHandler: "onCallFailure",
	}
	sipMessages["491"] = &SipMessage{
		Reason:       "Request Pending",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["493"] = &SipMessage{
		Reason:       "Undecipherable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["494"] = &SipMessage{
		Reason:       "Security Agreement Required",
		Rfc:          "RFC3329",
		TropoHandler: "onCallFailure",
	}
	sipMessages["500"] = &SipMessage{
		Reason:       "Server Internal Error",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["501"] = &SipMessage{
		Reason:       "Not Implemented",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["502"] = &SipMessage{
		Reason:       "Bad Gateway",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["503"] = &SipMessage{
		Reason:       "Service Unavailable",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["504"] = &SipMessage{
		Reason:       "Server Time-out",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["505"] = &SipMessage{
		Reason:       "Version Not Supported",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["513"] = &SipMessage{
		Reason:       "Message Too Large",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["580"] = &SipMessage{
		Reason:       "Precondition Failure",
		Rfc:          "RFC3312",
		TropoHandler: "onCallFailure",
	}
	sipMessages["600"] = &SipMessage{
		Reason:       "Busy Everywhere",
		Rfc:          "RFC3261",
		TropoHandler: "onBusy",
	}
	sipMessages["603"] = &SipMessage{
		Reason:       "Decline",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["604"] = &SipMessage{
		Reason:       "Does Not Exist Anywhere",
		Rfc:          "RFC3261",
		TropoHandler: "onCallFailure",
	}
	sipMessages["606"] = &SipMessage{
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
			table.Append([]string{code, s.Reason, s.TropoHandler, s.Rfc})
			table.Render()
		} else {
			terminal.Stdout.Color("r").Print("Invalid or unknown sip code").Nl().Reset()

		}

	}

}
