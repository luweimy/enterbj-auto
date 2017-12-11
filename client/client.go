package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/amlun/enterbj/request"
	"github.com/amlun/enterbj/response"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
)

type Config struct {
}

type Client struct {
}

func New(config Config) *Client {
	return &Client{}
}

var errRequest = errors.New("generate request error")

// GetPersonInfo 获取用户信息
func (e *Client) GetPersonInfo(userID string) (*response.PersonInfo, error) {
	PersonInfoURL := "https://bjjj.zhongchebaolian.com/industryguild_mobile_standard_self2.1.2/mobile/standard/getpersonalinfor?"
	var reqBody request.PersonInfo
	reqBody.Appkey = ""
	reqBody.Dicver = ""
	reqBody.SN = ""
	reqBody.UserId = userID

	r, err := query.Values(reqBody)
	if err != nil {
		logrus.Error(err)
		return nil, nil
	}
	req, _ := http.NewRequest("GET", PersonInfoURL+r.Encode(), nil)
	req.Header = commonHeader

	if req == nil {
		return nil, errRequest
	}
	var repBody *response.PersonInfo
	if _, err := sendRequest(req, &repBody); err != nil {
		return nil, err
	}
	return repBody, nil

}

func sendRequest(req *http.Request, v interface{}) (resp *http.Response, err error) {
	logrus.Debugf("send request (%v)", req)
	resp, err = httpClient.Do(req)
	logrus.Debugf("receive response (%v)", resp)
	if err != nil {
		logrus.Errorf("send request error (%v)", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logrus.Debugf("receive response body (%s)", body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return resp, nil
}
