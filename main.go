package main

import (
	"fmt"

	"github.com/skkim-01/esquery/csvs"
	"github.com/skkim-01/esquery/utils"
	HttpsUtil "github.com/skkim-01/esquery/utils/httpsutil"
)

var g_query string = `
{"query":{"bool":{
	"must":[
		{"match":{"source_service_code":"%v"}},
		{"match":{"sa_guid":"%v"}}]
}}}
`

func getQuery(guid, code string) string {
	return fmt.Sprintf(g_query, code, guid)
}

// 1. open csv
// 2. get guid/svc
// 3. es query
// 4. check
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

		fmt.Printf("#DBG\tGUID: %v, SERVICE: %v, ticketid: %v\n", guid, service, ticketid)

		strQuery := getQuery(guid, service)
		fmt.Printf("#DBG\tQueryString: %v\n", strQuery)

		break
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

		fmt.Println("#DBG\t", jsonResponse.PPrint())

		totalCount := jsonResponse.Find("hits.total.value").(int)
		for i := 0; i < totalCount; i++ {
			fmt.Printf("#DBG\t timeStamp: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.timestamp", i)))
			fmt.Printf("#DBG\t channel_type_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.channel_type_code", i)))
			fmt.Printf("#DBG\t source_service_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.source_service_code", i)))
			fmt.Printf("#DBG\t sa_guid: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.sa_guid", i)))
			fmt.Printf("#DBG\t ticket_id: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.ticket_id", i)))

			fmt.Printf("#DBG\t resource: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.resource", i)))
			fmt.Printf("#DBG\t event_type: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.event_type", i)))
			fmt.Printf("#DBG\t response_status: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.response_status", i)))

			fmt.Printf("#DBG\t error_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.error_code", i)))
			fmt.Printf("#DBG\t error_message: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v.error_message", i)))
		}

		// hitCount := len(jsonResponse.Find("hits.hits").([]interface{}))
		// for i := 0; i < hitCount; i++ {
		// 	req := fmt.Sprintf("hits.hits.%v.api.request", i)
		// 	res := fmt.Sprintf("hits.hits.%v.api.response", i)
		// 	reqmsg := jsonResponse.Find(req)
		// 	resmsg := jsonResponse.Find(res)
		// 	fmt.Printf("#DBG\tREQ: %v\n", reqmsg)
		// 	fmt.Printf("#DBG\tREQ: %v\n", resmsg)
		// }
	}

	// conn := HttpsUtil.NewReqInfo()
	// conn.SetURL("http://10.15.34.123:9210/gklog-api-2023.01.04/_search/?pretty")
	// conn.SetMethod("POST")
	// conn.AppendHeader("Authorization", "Basic Z2thZG1pbjpycGRseG1hcHAwMQ==")
	// conn.AppendHeader("Content-Type", "application/json")
	// conn.SetBody([]byte(g_query))
	// response, err := HttpsUtil.SendRequest(conn)
	// if nil != err {
	// 	fmt.Println(err)
	// 	return
	// }

	// jsonResponse, err := HttpsUtil.ResponseBodyParser(response)
	// if nil != err {
	// 	fmt.Println(err)
	// 	return
	// }

	// hitCount := len(jsonResponse.Find("hits.hits").([]interface{}))
	// for i := 0; i < hitCount; i++ {
	// 	req := fmt.Sprintf("hits.hits.%v.api.request", i)
	// 	res := fmt.Sprintf("hits.hits.%v.api.response", i)
	// 	reqmsg := jsonResponse.Find(req)
	// 	resmsg := jsonResponse.Find(res)
	// 	fmt.Printf("#DBG\tREQ: %v\n", reqmsg)
	// 	fmt.Printf("#DBG\tREQ: %v\n", resmsg)
	// }
}

//error_message
//query :=
