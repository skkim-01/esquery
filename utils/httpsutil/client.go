package HttpsUtil

import (
	"bytes"
	"net/http"
	"time"
)

func SendReq(r *ReqInfo) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.reqBody))
	if nil != err {
		return nil, err
	}

	// set request header
	for k, v := range r.reqHeader {
		req.Header.Add(k, v)
	}

	// set tls enable
	cli := &http.Client{
		Timeout: time.Second * 10,
	}

	req.Close = true

	return cli.Do(req)
}

// SendRequest : client request call wrapper - http/https
func SendRequest(r *ReqInfo) (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.reqBody))
	if nil != err {
		return nil, err
	}

	// set request basic auth
	if r.User != "" {
		req.SetBasicAuth(r.User, r.Passwd)
	}

	// set request header
	for k, v := range r.reqHeader {
		req.Header.Add(k, v)
	}

	// set tls enable
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	cli := &http.Client{
		//Transport: tr,
		Timeout: time.Second * 10,
	}

	req.Close = true

	response, err := cli.Do(req)
	if nil != err {
		return nil, err
	}

	return response, nil
}
