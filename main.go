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

		conn := HttpsUtil.NewReqInfo()
		conn.SetURL("http://10.15.34.123:9210/gklog-api-*/_search/?pretty")
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

		hitCount := len(jsonResponse.Find("hits.hits").([]interface{}))
		for i := 0; i < hitCount; i++ {
			req := fmt.Sprintf("hits.hits.%v.api.request", i)
			res := fmt.Sprintf("hits.hits.%v.api.response", i)
			reqmsg := jsonResponse.Find(req)
			resmsg := jsonResponse.Find(res)
			fmt.Printf("#DBG\tREQ: %v\n", reqmsg)
			fmt.Printf("#DBG\tREQ: %v\n", resmsg)
		}
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
