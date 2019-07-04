package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strconv"
	"time"
)

//Get a common interface for enterprise authentication information
func CorpAuth(url string, accessKey string, suiteTicket string, authCorpId string) (string, error) {
	timestamp := time.Now().UnixNano() / 1e6
	nativeSignature := strconv.FormatInt(timestamp, 10) + "\n" + suiteTicket

	afterHmacSHA256 := encrypt.SHA256(nativeSignature)
	afterBase64 := encrypt.BASE64(afterHmacSHA256)
	afterUrlEncode := encrypt.URLEncode(afterBase64)

	params := map[string]string{
		"timestamp":   strconv.FormatInt(timestamp, 10),
		"accessKey":   accessKey,
		"suiteTicket": suiteTicket,
		"signature":   afterUrlEncode,
	}
	body, _ := json.ToJson(map[string]string{
		"auth_corpid": authCorpId,
	})
	return http.Post(url, params, body)
}

//Desc: 获取企业凭证
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/hv357q
func GetCorpToken(accessKey string, suiteTicket string, authCorpId string) (GetCorpTokenResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_corp_token", accessKey, suiteTicket, authCorpId)
	resp := GetCorpTokenResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取企业授权信息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/fmdqvx
func GetAuthInfo(accessKey string, suiteTicket string, authCorpId string) (GetAuthInfoResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_auth_info", accessKey, suiteTicket, authCorpId)
	resp := GetAuthInfoResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
