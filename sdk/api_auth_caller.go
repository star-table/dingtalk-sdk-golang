package sdk

import (
	"strconv"
	"time"

	"github.com/flyingtime/dingtalk-sdk-golang/encrypt"
	"github.com/flyingtime/dingtalk-sdk-golang/http"
	"github.com/flyingtime/dingtalk-sdk-golang/json"
)

//CorpAuth is the common signature methods for multiple authentication interfaces
//Get a common interface for enterprise authentication information
func CorpAuth(url string, suiteKey string, suiteSecret string, suiteTicket string, authCorpId string, agentId int64) (string, error) {
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
		bodyParams["agentid"] = strconv.FormatInt(agentId, 10)
		bodyParams["suite_key"] = suiteKey
	}

	body, _ := json.ToJson(bodyParams)
	return http.Post(url, params, body)
}

// sso是否获取管理后台免登陆token
// sso = ture;appkey为管理后台免登陆的企业corpId,appsecret为该企业的SSOsecret
// sso = false;appkey,appsecret为应用的信息
func GetToken(appkey, appsecret string, sso bool) (string, error) {
	type AcessTokenResp struct {
		BaseResp
		AcessToken string `json:"access_token"`
	}

	var body string
	var err error

	if sso {
		params := map[string]string{
			"corpid":     appkey,
			"corpsecret": appsecret,
		}

		body, err = http.Get("https://oapi.dingtalk.com/sso/gettoken", params)
		if err != nil {
			return "", err
		}
	} else {
		params := map[string]string{
			"appkey":    appkey,
			"appsecret": appsecret,
		}

		body, err = http.Get("https://oapi.dingtalk.com/gettoken", params)
		if err != nil {
			return "", err
		}
	}

	resp := AcessTokenResp{}
	json.FromJson(body, &resp)

	return resp.AcessToken, nil
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
func (corp *Corp) GetAgent(agentId int64) (GetAgentResp, error) {
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

func (s *DingTalkSDK) GetSuiteToken(suiteTicket string) (*GetSuiteAccessTokenResp, error) {
	params := map[string]string{
		"suite_key":    s.SuiteKey,
		"suite_secret": s.SuiteSecret,
		"suiteTicket":  suiteTicket,
	}

	jsonStr, err := json.ToJson(params)
	if err != nil {
		return nil, err
	}
	body, err := http.Post("https://oapi.dingtalk.com/service/get_suite_token", nil, jsonStr)
	resp := GetSuiteAccessTokenResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

func (s *DingTalkSDK) GetPermanentCode(suiteAccessToken string, tmpAuthCode string) (*GetPermanentCodeResp, error) {
	params := map[string]string{
		"suite_access_token": suiteAccessToken,
	}

	bodyParams := map[string]string{
		"tmp_auth_code": tmpAuthCode,
	}

	jsonStr, err := json.ToJson(bodyParams)
	if err != nil {
		return nil, err
	}
	body, err := http.Post("https://oapi.dingtalk.com/service/get_permanent_code", params, jsonStr)
	resp := GetPermanentCodeResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

func (s *DingTalkSDK) ActivateSuite(suiteAccessToken string, authCorpId string, permanentCode string) (*BaseResp, error) {
	params := map[string]string{
		"suite_access_token": suiteAccessToken,
	}

	bodyParams := map[string]string{
		"suite_key":      s.SuiteKey,
		"auth_corpid":    authCorpId,
		"permanent_code": permanentCode,
	}

	jsonStr, err := json.ToJson(bodyParams)
	if err != nil {
		return nil, err
	}
	body, err := http.Post("https://oapi.dingtalk.com/service/activate_suite", params, jsonStr)
	resp := BaseResp{}
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
