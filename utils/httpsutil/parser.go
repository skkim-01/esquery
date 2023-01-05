package HttpsUtil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	JsonMapper "github.com/skkim-01/json-mapper/src"
)

// RequestBodyParser : parse request body
func RequestBodyParser(r *http.Request) (*JsonMapper.JsonMap, error) {
	// close request body
	defer func() {
		r.Body.Close()
	}()

	// when body is nil, return
	if nil == r.Body {
		return nil, nil
	}

	// decode requset body to json
	var reqBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if nil != err {
		return nil, err
	}

	jsonBytes, err := JsonMapper.ToJson(reqBody)
	if nil != err {
		return nil, err
	}

	// create json map
	jsonBody, err := JsonMapper.NewBytes(jsonBytes)
	if nil != err {
		return nil, err
	}

	return jsonBody, nil
}

func ResponseBodyParser(res *http.Response) (*JsonMapper.JsonMap, error) {
	// close response body
	defer func() {
		res.Body.Close()
	}()

	if http.StatusOK == res.StatusCode {
		b, _ := ioutil.ReadAll(res.Body)
		jsonBody, err := JsonMapper.NewBytes(b)
		if nil != err {
			fmt.Println(string(b))
			return nil, err
		} else {
			return jsonBody, nil
		}
	} else {
		return nil, errors.New(fmt.Sprintf("%v %v", res.StatusCode, res.Status))
	}
}

// ResponseParser : parse http response
func ResponseParser(res *http.Response) (string, string) {
	// close response body
	defer func() {
		res.Body.Close()
	}()

	if http.StatusOK == res.StatusCode {
		b, _ := ioutil.ReadAll(res.Body)
		return res.Status, string(b)
	}
	return res.Status, ""
}
