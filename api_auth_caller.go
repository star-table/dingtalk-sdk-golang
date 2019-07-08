package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strconv"
	"time"
)

//CorpAuth is the common signature methods for multiple authentication interfaces
//Get a common interface for enterprise authentication information
func CorpAuth(url string, suiteKey string, suiteSecret string, suiteTicket string, authCorpId string, agentId int) (string, error) {
	timestamp := time.Now().UnixNano() / 1e6
	nativeSignature := strconv.FormatInt(timestamp, 10) + "\n" + suiteTicket

	afterHmacSHA256 := encrypt.SHA256(nativeSignature, suiteSecret)
	afterBase64 := encrypt.BASE64(afterHmacSHA256)
	afterUrlEncode := encrypt.URLEncode(afterBase64)

	params := map[string]string{
		"timestamp":   strconv.FormatInt(timestamp, 10),
		"accessKey":   suiteKey,
		"suiteTicket": suiteTicket,
		"signature":   afterUrlEncode,
	}

	bodyParams := map[string]string{
		"auth_corpid": authCorpId,
	}
	if agentId != 0 {
		bodyParams["agentid"] = strconv.Itoa(agentId)
		bodyParams["suite_key"] = suiteKey
	}

	body, _ := json.ToJson(bodyParams)
	return http.Post(url, params, body)
}

//GetCorpToken can get corporation's token
//Desc: 获取企业凭证
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/hv357q
func (corp *Corp) GetCorpToken() (GetCorpTokenResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_corp_token", corp.SuiteKey, corp.SuiteSecret, corp.SuiteTicket, corp.CorpId, 0)
	resp := GetCorpTokenResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取企业授权信息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/fmdqvx
func (corp *Corp) GetAuthInfo() (GetAuthInfoResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_auth_info", corp.SuiteKey, corp.SuiteSecret, corp.SuiteTicket, corp.CorpId, 0)
	resp := GetAuthInfoResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取授权应用信息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/vfitg0
func (corp *Corp) GetAgent(agentId int) (GetAgentResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_agent", corp.SuiteKey, corp.SuiteSecret, corp.SuiteTicket, corp.CorpId, agentId)
	resp := GetAgentResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取jsapi_ticket
//Doc: https://open-doc.dingtalk.com/microapp/dev/uwa7vs
func (client *DingTalkClient) GetJsApiTicket(authType string) (*GetJsApiTicketResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"type":         authType,
	}
	body, err := http.Get("https://oapi.dingtalk.com/get_jsapi_ticket", params)
	resp := GetJsApiTicketResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

func CalculateJsApiSign(ticket string, nonceStr string, timestamp int64, url string) string {
	plain := "jsapi_ticket=" + ticket + "&noncestr=" + nonceStr + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&url=" + url
	return encrypt.SHA1(plain)
}
