package main

import (
	"fmt"

	"github.com/skkim-01/esquery/csvs"
	"github.com/skkim-01/esquery/utils"
	HttpsUtil "github.com/skkim-01/esquery/utils/httpsutil"
	slog "github.com/skkim-01/esquery/utils/simplelog"
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
	utils.Sfolder(&checklist, "./res")

	fmt.Println(">DBG\tcheck files:", checklist)
	for _, v := range checklist {
		fmt.Println(">DBG\tfile:", v)
		_search(v)
	}
}

func _search(strCSVFile string) {
	var guid string
	var service string
	var ticketid string
	var strLogFile string = "../" + strCSVFile + ".log"
	fmt.Println(">DBG\tresult is written at", strLogFile)

	csvObject := csvs.NewCSVHandle()
	err := csvObject.OpenCSV(strCSVFile)
	if nil != err {
		fmt.Println(err)
		return
	}

	for row := 1; row < csvObject.RowCount(); row++ {
		guid = csvObject.GetField(row, 35)
		service = csvObject.GetField(row, 36)
		ticketid = csvObject.GetField(row, 38)

		fmt.Printf("GUID: %v, SERVICE: %v, ticketid: %v\n\n", guid, service, ticketid)
		slog.Write(strLogFile, fmt.Sprintf("\tGUID: %v, SERVICE: %v, ticketid: %v", guid, service, ticketid))

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
		//totalCount := (int)(jsonResponse.Find("hits.hits").(float64))
		totalCount := len(jsonResponse.Find("hits.hits").([]interface{}))
		for i := 0; i < totalCount; i++ {
			slog.Write(strLogFile, fmt.Sprintf("# COUNT %v #", i))
			slog.Write(strLogFile, fmt.Sprintf("timeStamp: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.timestamp", i))))
			slog.Write(strLogFile, fmt.Sprintf("channel_type_code: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.channel_type_code", i))))
			slog.Write(strLogFile, fmt.Sprintf("source_service_code: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.source_service_code", i))))
			slog.Write(strLogFile, fmt.Sprintf("sa_guid: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.sa_guid", i))))
			slog.Write(strLogFile, fmt.Sprintf("ticket_id: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.ticket_id", i))))

			slog.Write(strLogFile, fmt.Sprintf("event_type: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.event_type", i))))
			slog.Write(strLogFile, fmt.Sprintf("response_status: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.response_status", i))))

			slog.Write(strLogFile, fmt.Sprintf("error_code: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i))))
			slog.Write(strLogFile, fmt.Sprintf("error_message: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_message", i))))
			slog.Write(strLogFile, "")

			ifaceErrorCode := jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i))
			if ifaceErrorCode == nil {
				slResponseCode = append(slResponseCode, "SUCCESS")
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
			if v == "SUCCESS" {
				nSuccessIdx = i
				bResult = true
				break
			} else if v == "PR5004" {
				bResult = false
				break
			}
		}
		slog.Write(strLogFile, fmt.Sprintf("Response Code Order:", slResponseCode))
		slog.Write(strLogFile, fmt.Sprintf("Is Reopened:%v", bResult))
		if bResult {
			slog.Write(strLogFile, fmt.Sprintf("SuccessTime: %v", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.timestamp", nSuccessIdx))))
		}
		slog.Write(strLogFile, "")
	}
}
