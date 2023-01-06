package main

import (
	"fmt"

	"github.com/skkim-01/esquery/csvs"
	"github.com/skkim-01/esquery/utils"
	HttpsUtil "github.com/skkim-01/esquery/utils/httpsutil"
)

var g_query string = `
{
	"sort" : [{ "@timestamp" : "desc" }],
	"query":{"bool":{
		"must":[
			{"match":{"source_service_code":"%v"}},
			{"match":{"sa_guid":"%v"}}]
}}}
`

func getQuery(guid, code string) string {
	return fmt.Sprintf(g_query, code, guid)
}

// 1. open csv : o
// 2. get guid/svc : o
// 3. es query : o
// 4. check : o

// 5. update result
// 6. make csv

func main() {
	checklist := make([]string, 0)
	utils.Sfolder(&checklist, "./resources")

	fmt.Println(checklist)

	csvObject := csvs.NewCSVHandle()
	err := csvObject.OpenCSV(checklist[0])
	if nil != err {
		fmt.Println(err)
		return
	}

	var guid string
	var service string
	var ticketid string

	for row := 1; row < csvObject.RowCount(); row++ {
		guid = csvObject.GetField(row, 35)
		service = csvObject.GetField(row, 36)
		ticketid = csvObject.GetField(row, 38)

		fmt.Printf("> DBG\tGUID: %v, SERVICE: %v, ticketid: %v\n", guid, service, ticketid)

		strQuery := getQuery(guid, service)
		conn := HttpsUtil.NewReqInfo()
		conn.SetURL("http://10.15.34.123:9210/gklog-api-2023.01.*/_search/?pretty")
		conn.SetMethod("POST")
		conn.AppendHeader("Authorization", "Basic Z2thZG1pbjpycGRseG1hcHAwMQ==")
		conn.AppendHeader("Content-Type", "application/json")
		conn.SetBody([]byte(strQuery))

		response, err := HttpsUtil.SendRequest(conn)
		if nil != err {
			fmt.Println(err)
			continue
		}

		jsonResponse, err := HttpsUtil.ResponseBodyParser(response)
		if nil != err {
			fmt.Println(err)
			continue
		}

		slResponseCode := make([]string, 0)
		totalCount := (int)(jsonResponse.Find("hits.total.value").(float64))
		for i := 0; i < totalCount; i++ {
			fmt.Printf("> DBG\t # COUNT %v #\n", i)
			fmt.Printf("> DBG\t timeStamp: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.timestamp", i)))
			fmt.Printf("> DBG\t channel_type_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.channel_type_code", i)))
			fmt.Printf("> DBG\t source_service_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.source_service_code", i)))
			fmt.Printf("> DBG\t sa_guid: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.sa_guid", i)))
			fmt.Printf("> DBG\t ticket_id: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.ticket_id", i)))

			fmt.Printf("> DBG\t event_type: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.event_type", i)))
			fmt.Printf("> DBG\t response_status: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.response_status", i)))

			fmt.Printf("> DBG\t error_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i)))
			fmt.Printf("> DBG\t error_message: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_message", i)))
			fmt.Printf("\n")

			ifaceErrorCode := jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i))
			if ifaceErrorCode == nil {
				slResponseCode = append(slResponseCode, "")
			} else {
				slResponseCode = append(slResponseCode, ifaceErrorCode.(string))
			}
		}

		// PR5004 : 0, nil: 1, !PR5004: X
		// order by timestamp descending
		// newer - older
		// x - x - x : true
		// x - 1 - 0 : true
		// x - 0 - 1 : false
		// after not x, don't care
		var bResult bool = true
		var nSuccessIdx int = -1
		for i, v := range slResponseCode {
			if v == "" {
				nSuccessIdx = i
				bResult = true
				break
			} else if v == "PR5004" {
				bResult = false
				break
			}
		}
		fmt.Println("> DBG\t Response Code Order:", slResponseCode)
		fmt.Printf("> DBG\tIs Reopened:%v\n\n", bResult)
		if bResult {
			fmt.Printf("> DBG\t SuccessTime: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.timestamp", nSuccessIdx)))
		}
	}
}
