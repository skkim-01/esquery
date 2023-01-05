package main

import (
	"fmt"

	"github.com/skkim-01/esquery/csvs"
	"github.com/skkim-01/esquery/utils"
	HttpsUtil "github.com/skkim-01/esquery/utils/httpsutil"
	JsonMapper "github.com/skkim-01/json-mapper/src"
)

var g_query string = `
{
	"sort" : [{ "@timestamp" : "desc" },
	"query":{"bool":{
		"must":[
			{"match":{"source_service_code":"%v"}},
			{"match":{"sa_guid":"%v"}}]
}}}
`

func getQuery(guid, code string) string {
	return fmt.Sprintf(g_query, code, guid)
}

var g_test_json string = `
{
	"_shards": {
	  "failed": 0,
	  "skipped": 0,
	  "successful": 5,
	  "total": 5
	},
	"hits": {
	  "hits": [
		{
		  "_id": "Vy10eYUBEKlRkG37ZXEx",
		  "_index": "gklog-api-2023.01.03",
		  "_score": 13.001686,
		  "_source": {
			"@timestamp": "2023-01-03T21:05:36.810Z",
			"@version": "1",
			"agent": {
			  "name": "p1ec1-was02"
			},
			"api.request": "{\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}",
			"api.response": "{\"result_code\":\"Fail\",\"error_code\":\"PR2003\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"sa_guid and service_code are already in progress.\"}",
			"channel_type_code": "CCPA",
			"cloud": {
			  "account": {},
			  "availability_zone": "eu-central-1b",
			  "image": {},
			  "instance": {
				"id": "i-012f7f9810157215b"
			  },
			  "machine": {},
			  "region": "eu-central-1",
			  "service": {
				"name": "EC2"
			  }
			},
			"comp": "API",
			"country_code": "US",
			"ecs": {
			  "version": "8.0.0"
			},
			"error_code": "PR2003",
			"error_message": "sa_guid and service_code are already in progress.",
			"event_sub_type": "Register",
			"event_type": "Erasure",
			"gubun": "GK",
			"host": {
			  "hostname": "p1ec1-was02",
			  "ip": [
				"10.15.22.206",
				"fe80::4b2:5eff:fe4b:8c38"
			  ],
			  "mac": [
				"06:b2:5e:4b:8c:38"
			  ],
			  "os": {
				"type": "linux"
			  }
			},
			"input": {
			  "type": "log"
			},
			"log": {
			  "file": {
				"path": "/home02/tomcat-logs/logs/api/comp/gk/v2/logs/gk/api/elk/GK_API.gk1.log"
			  },
			  "offset": 859883
			},
			"loglevel": "INFO",
			"message": "API§2023-01-03_21:05:36.810§INFO §GK§resource=registrationGlobalRequest§sa_guid=vydvl8qioc§parent_ticket_id=CCPA_UF5SQSH254§ticket_id=CCPA_UF5SQSH254§source_service_code=GKSCD10004§channel_type_code=CCPA§event_type=E§event_sub_type=R§country_code=US§error_code=PR2003§error_message=sa_guid and service_code are already in progress.§response_status=400§response_time=5§api.request={\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}§api.response={\"result_code\":\"Fail\",\"error_code\":\"PR2003\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"sa_guid and service_code are already in progress.\"}§",
			"parent_ticket_id": "CCPA_UF5SQSH254",
			"resource": "registrationGlobalRequest",
			"response_status": "400",
			"response_time": "5",
			"sa_guid": "vydvl8qioc",
			"source_service_code": "GKSCD10004",
			"tags": [
			  "api",
			  "beats_input_codec_plain_applied"
			],
			"ticket_id": "CCPA_UF5SQSH254",
			"timestamp": "2023-01-03_21:05:36.810"
		  },
		  "_type": "_doc"
		},
		{
		  "_id": "cfFPeoUBwE5SqNpdg05r",
		  "_index": "gklog-api-2023.01.04",
		  "_score": 12.823127,
		  "_source": {
			"@timestamp": "2023-01-04T01:05:08.050Z",
			"@version": "1",
			"agent": {
			  "name": "p1ec1-was01"
			},
			"api.request": "{\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}",
			"api.response": "{\"result_code\":\"Fail\",\"error_code\":\"PR2003\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"sa_guid and service_code are already in progress.\"}",
			"channel_type_code": "CCPA",
			"cloud": {
			  "account": {},
			  "availability_zone": "eu-central-1a",
			  "image": {},
			  "instance": {
				"id": "i-0faefdcf6ca938b4f"
			  },
			  "machine": {},
			  "region": "eu-central-1",
			  "service": {
				"name": "EC2"
			  }
			},
			"comp": "API",
			"country_code": "US",
			"ecs": {
			  "version": "8.0.0"
			},
			"error_code": "PR2003",
			"error_message": "sa_guid and service_code are already in progress.",
			"event_sub_type": "Register",
			"event_type": "Erasure",
			"gubun": "GK",
			"host": {
			  "hostname": "p1ec1-was01",
			  "ip": [
				"10.15.21.158",
				"fe80::c:f1ff:fec7:d8fa"
			  ],
			  "mac": [
				"02:0c:f1:c7:d8:fa"
			  ],
			  "os": {
				"type": "linux"
			  }
			},
			"input": {
			  "type": "log"
			},
			"log": {
			  "file": {
				"path": "/home02/tomcat-logs/logs/api/comp/gk/v2/logs/gk/api/elk/GK_API.gk1.log"
			  },
			  "offset": 702350
			},
			"loglevel": "INFO",
			"message": "API§2023-01-04_01:05:08.050§INFO §GK§resource=registrationGlobalRequest§sa_guid=vydvl8qioc§parent_ticket_id=CCPA_UF5SQSH254§ticket_id=CCPA_UF5SQSH254§source_service_code=GKSCD10004§channel_type_code=CCPA§event_type=E§event_sub_type=R§country_code=US§error_code=PR2003§error_message=sa_guid and service_code are already in progress.§response_status=400§response_time=6§api.request={\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}§api.response={\"result_code\":\"Fail\",\"error_code\":\"PR2003\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"sa_guid and service_code are already in progress.\"}§",
			"parent_ticket_id": "CCPA_UF5SQSH254",
			"resource": "registrationGlobalRequest",
			"response_status": "400",
			"response_time": "6",
			"sa_guid": "vydvl8qioc",
			"source_service_code": "GKSCD10004",
			"tags": [
			  "api",
			  "beats_input_codec_plain_applied"
			],
			"ticket_id": "CCPA_UF5SQSH254",
			"timestamp": "2023-01-04_01:05:08.050"
		  },
		  "_type": "_doc"
		},
		{
		  "_id": "5SnifIUBwE5SqNpdlLib",
		  "_index": "gklog-api-2023.01.04",
		  "_score": 12.823127,
		  "_source": {
			"@timestamp": "2023-01-04T13:05:01.068Z",
			"@version": "1",
			"agent": {
			  "name": "p1ec1-was01"
			},
			"api.request": "{\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}",
			"api.response": "{\"result_code\":\"Success\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"g_ticket_id\":\"gk000019633832\"}",
			"channel_type_code": "CCPA",
			"cloud": {
			  "account": {},
			  "availability_zone": "eu-central-1a",
			  "image": {},
			  "instance": {
				"id": "i-0faefdcf6ca938b4f"
			  },
			  "machine": {},
			  "region": "eu-central-1",
			  "service": {
				"name": "EC2"
			  }
			},
			"comp": "API",
			"country_code": "US",
			"ecs": {
			  "version": "8.0.0"
			},
			"event_sub_type": "Register",
			"event_type": "Erasure",
			"gk_id": "gk000019633832",
			"gubun": "GK",
			"host": {
			  "hostname": "p1ec1-was01",
			  "ip": [
				"10.15.21.158",
				"fe80::c:f1ff:fec7:d8fa"
			  ],
			  "mac": [
				"02:0c:f1:c7:d8:fa"
			  ],
			  "os": {
				"type": "linux"
			  }
			},
			"input": {
			  "type": "log"
			},
			"log": {
			  "file": {
				"path": "/home02/tomcat-logs/logs/api/comp/gk/v2/logs/gk/api/elk/GK_API.gk1.log"
			  },
			  "offset": 554756
			},
			"loglevel": "INFO",
			"message": "API§2023-01-04_13:05:01.068§INFO §GK§resource=registrationGlobalRequest§sa_guid=vydvl8qioc§parent_ticket_id=CCPA_UF5SQSH254§ticket_id=CCPA_UF5SQSH254§source_service_code=GKSCD10004§channel_type_code=CCPA§event_type=E§event_sub_type=R§country_code=US§gk_id=gk000019633832§response_status=200§response_time=19§api.request={\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}§api.response={\"result_code\":\"Success\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"g_ticket_id\":\"gk000019633832\"}§",
			"parent_ticket_id": "CCPA_UF5SQSH254",
			"resource": "registrationGlobalRequest",
			"response_status": "200",
			"response_time": "19",
			"sa_guid": "vydvl8qioc",
			"source_service_code": "GKSCD10004",
			"tags": [
			  "api",
			  "beats_input_codec_plain_applied"
			],
			"ticket_id": "CCPA_UF5SQSH254",
			"timestamp": "2023-01-04_13:05:01.068"
		  },
		  "_type": "_doc"
		},
		{
		  "_id": "JXEre4UBw89EdFJ2J5wR",
		  "_index": "gklog-api-2023.01.04",
		  "_score": 12.823127,
		  "_source": {
			"@timestamp": "2023-01-04T05:05:01.943Z",
			"@version": "1",
			"agent": {
			  "name": "p1ec1-was02"
			},
			"api.request": "{\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}",
			"api.response": "{\"result_code\":\"Fail\",\"error_code\":\"PR2003\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"sa_guid and service_code are already in progress.\"}",
			"channel_type_code": "CCPA",
			"cloud": {
			  "account": {},
			  "availability_zone": "eu-central-1b",
			  "image": {},
			  "instance": {
				"id": "i-012f7f9810157215b"
			  },
			  "machine": {},
			  "region": "eu-central-1",
			  "service": {
				"name": "EC2"
			  }
			},
			"comp": "API",
			"country_code": "US",
			"ecs": {
			  "version": "8.0.0"
			},
			"error_code": "PR2003",
			"error_message": "sa_guid and service_code are already in progress.",
			"event_sub_type": "Register",
			"event_type": "Erasure",
			"gubun": "GK",
			"host": {
			  "hostname": "p1ec1-was02",
			  "ip": [
				"10.15.22.206",
				"fe80::4b2:5eff:fe4b:8c38"
			  ],
			  "mac": [
				"06:b2:5e:4b:8c:38"
			  ],
			  "os": {
				"type": "linux"
			  }
			},
			"input": {
			  "type": "log"
			},
			"log": {
			  "file": {
				"path": "/home02/tomcat-logs/logs/api/comp/gk/v2/logs/gk/api/elk/GK_API.gk1.log"
			  },
			  "offset": 375126
			},
			"loglevel": "INFO",
			"message": "API§2023-01-04_05:05:01.943§INFO §GK§resource=registrationGlobalRequest§sa_guid=vydvl8qioc§parent_ticket_id=CCPA_UF5SQSH254§ticket_id=CCPA_UF5SQSH254§source_service_code=GKSCD10004§channel_type_code=CCPA§event_type=E§event_sub_type=R§country_code=US§error_code=PR2003§error_message=sa_guid and service_code are already in progress.§response_status=400§response_time=5§api.request={\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}§api.response={\"result_code\":\"Fail\",\"error_code\":\"PR2003\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"sa_guid and service_code are already in progress.\"}§",
			"parent_ticket_id": "CCPA_UF5SQSH254",
			"resource": "registrationGlobalRequest",
			"response_status": "400",
			"response_time": "5",
			"sa_guid": "vydvl8qioc",
			"source_service_code": "GKSCD10004",
			"tags": [
			  "api",
			  "beats_input_codec_plain_applied"
			],
			"ticket_id": "CCPA_UF5SQSH254",
			"timestamp": "2023-01-04_05:05:01.943"
		  },
		  "_type": "_doc"
		},
		{
		  "_id": "6C4HfIUBEKlRkG37CNSn",
		  "_index": "gklog-api-2023.01.04",
		  "_score": 12.823127,
		  "_source": {
			"@timestamp": "2023-01-04T09:05:01.534Z",
			"@version": "1",
			"agent": {
			  "name": "p1ec1-was02"
			},
			"api.request": "{\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}",
			"api.response": "{\"result_code\":\"Fail\",\"error_code\":\"PR5004\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"The value of the event_type is invalid.\"}",
			"channel_type_code": "CCPA",
			"cloud": {
			  "account": {},
			  "availability_zone": "eu-central-1b",
			  "image": {},
			  "instance": {
				"id": "i-012f7f9810157215b"
			  },
			  "machine": {},
			  "region": "eu-central-1",
			  "service": {
				"name": "EC2"
			  }
			},
			"comp": "API",
			"country_code": "US",
			"ecs": {
			  "version": "8.0.0"
			},
			"error_code": "PR5004",
			"error_message": "The value of the event_type is invalid.",
			"event_sub_type": "Register",
			"event_type": "Erasure",
			"gubun": "GK",
			"host": {
			  "hostname": "p1ec1-was02",
			  "ip": [
				"10.15.22.206",
				"fe80::4b2:5eff:fe4b:8c38"
			  ],
			  "mac": [
				"06:b2:5e:4b:8c:38"
			  ],
			  "os": {
				"type": "linux"
			  }
			},
			"input": {
			  "type": "log"
			},
			"log": {
			  "file": {
				"path": "/home02/tomcat-logs/logs/api/comp/gk/v2/logs/gk/api/elk/GK_API.gk1.log"
			  },
			  "offset": 484237
			},
			"loglevel": "INFO",
			"message": "API§2023-01-04_09:05:01.534§INFO §GK§resource=registrationGlobalRequest§sa_guid=vydvl8qioc§parent_ticket_id=CCPA_UF5SQSH254§ticket_id=CCPA_UF5SQSH254§source_service_code=GKSCD10004§channel_type_code=CCPA§event_type=E§event_sub_type=R§country_code=US§error_code=PR5004§error_message=The value of the event_type is invalid.§response_status=400§response_time=5§api.request={\"ticket_id\":\"CCPA_UF5SQSH254\",\"sa_guid\":\"vydvl8qioc\",\"country_code\":\"US\",\"lang_code\":\"en\",\"event_type\":\"erasure\",\"password\":\"**********\",\"user_request_date\":\"20230101193440\",\"service_code\":\"CD30004\",\"channel_type_code\":\"CCPA\"}§api.response={\"result_code\":\"Fail\",\"error_code\":\"PR5004\",\"ticket_id\":\"CCPA_UF5SQSH254\",\"message\":\"The value of the event_type is invalid.\"}§",
			"parent_ticket_id": "CCPA_UF5SQSH254",
			"resource": "registrationGlobalRequest",
			"response_status": "400",
			"response_time": "5",
			"sa_guid": "vydvl8qioc",
			"source_service_code": "GKSCD10004",
			"tags": [
			  "api",
			  "beats_input_codec_plain_applied"
			],
			"ticket_id": "CCPA_UF5SQSH254",
			"timestamp": "2023-01-04_09:05:01.534"
		  },
		  "_type": "_doc"
		}
	  ],
	  "max_score": 13.001686,
	  "total": {
		"relation": "eq",
		"value": 5
	  }
	},
	"timed_out": false,
	"took": 3
  }  
`

// 1. open csv : o
// 2. get guid/svc : o
// 3. es query : o
// 4. check : o

// 5. update result
// 6. make csv

func _searchTest() {
	// js g_test_json
	// JsonMapper.JsonMap,

	jsonResponse, err := JsonMapper.NewBytes([]byte(g_test_json))
	if err != nil {
		fmt.Println(err)
		return
	}
	slResponseCode := make([]string, 0)
	totalCount := (int)(jsonResponse.Find("hits.total.value").(float64))
	for i := 0; i < totalCount; i++ {
		fmt.Printf("#DBG\t # COUNT %v #\n", i)
		fmt.Printf("#DBG\t timeStamp: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.timestamp", i)))
		fmt.Printf("#DBG\t channel_type_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.channel_type_code", i)))
		fmt.Printf("#DBG\t source_service_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.source_service_code", i)))
		fmt.Printf("#DBG\t sa_guid: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.sa_guid", i)))
		fmt.Printf("#DBG\t ticket_id: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.ticket_id", i)))

		fmt.Printf("#DBG\t resource: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.resource", i)))
		fmt.Printf("#DBG\t event_type: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.event_type", i)))
		fmt.Printf("#DBG\t response_status: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.response_status", i)))

		fmt.Printf("#DBG\t error_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i)))
		fmt.Printf("#DBG\t error_message: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_message", i)))
		fmt.Printf("\n")

		ifaceErrorCode := jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i))
		if ifaceErrorCode == nil {
			slResponseCode = append(slResponseCode, "")
		} else {
			slResponseCode = append(slResponseCode, ifaceErrorCode.(string))
		}
	}

	var bResult bool = false
	for i, v := range slResponseCode {
		fmt.Printf("#DBG\tindex:%v value:%v\n", i, v)
		if v == "" {
			bResult = true
		}
	}
	fmt.Printf("#DBG\tIs Reopened:%v\n\n", bResult)
}

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
		//fmt.Printf("#DBG\tQueryString: %v\n", strQuery)

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

		totalCount := (int)(jsonResponse.Find("hits.total.value").(float64))
		for i := 0; i < totalCount; i++ {
			fmt.Printf("#DBG\t # COUNT %v #\n", i)
			fmt.Printf("#DBG\t timeStamp: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.timestamp", i)))
			fmt.Printf("#DBG\t channel_type_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.channel_type_code", i)))
			fmt.Printf("#DBG\t source_service_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.source_service_code", i)))
			fmt.Printf("#DBG\t sa_guid: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.sa_guid", i)))
			fmt.Printf("#DBG\t ticket_id: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.ticket_id", i)))

			//fmt.Printf("#DBG\t resource: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.resource", i)))
			fmt.Printf("#DBG\t event_type: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.event_type", i)))
			fmt.Printf("#DBG\t response_status: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.response_status", i)))

			fmt.Printf("#DBG\t error_code: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_code", i)))
			fmt.Printf("#DBG\t error_message: %v\n", jsonResponse.Find(fmt.Sprintf("hits.hits.%v._source.error_message", i)))
			fmt.Printf("\n")
		}

		slResponseCode := make([]string, 0)
		var bResult bool = false
		for i, v := range slResponseCode {
			fmt.Printf("#DBG\tindex:%v value:%v\n", i, v)
			if v == "" {
				bResult = true
			}
		}
		fmt.Printf("#DBG\tIs Reopened:%v\n\n", bResult)
	}

}
