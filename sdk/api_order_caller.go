package sdk

import (
	"github.com/polaris-team/dingtalk-sdk-golang/encrypt"
	"github.com/polaris-team/dingtalk-sdk-golang/http"
	"github.com/polaris-team/dingtalk-sdk-golang/json"
	"strconv"
)

//内购商品订单处理完成
//doc:https://ding-doc.dingtalk.com/doc#/serverapi3/lg1nb7/NqyuB
func (client *DingTalkClient) OrderFinish(orderId int64) (*BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"biz_order_id": strconv.FormatInt(orderId, 10),
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/appstore/internal/order/finish", params, "")
	resp := BaseResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//获取内购商品SKU页面地址
//doc:https://ding-doc.dingtalk.com/doc#/serverapi3/lg1nb7/Uxjfs
func (client *DingTalkClient) GetSkuPage(goodsCode string, callBackPage string) (GetSkuPageResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"goods_code":   goodsCode,
	}
	if callBackPage != "" {
		params["callback_page"] = encrypt.URLEncode(callBackPage)
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/appstore/internal/skupage/get", params)
	resp := GetSkuPageResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//获取内购订单信息
//doc:https://ding-doc.dingtalk.com/doc#/serverapi3/lg1nb7/MgHqh
func (client *DingTalkClient) GetOrderInfo(orderId int64) (GetOrderInfoResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"biz_order_id": strconv.FormatInt(orderId, 10),
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/appstore/internal/order/get", params)
	resp := GetOrderInfoResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//应用内购商品核销
//doc:https://ding-doc.dingtalk.com/doc#/serverapi3/lg1nb7/XbiVN
func (client *DingTalkClient) OrderConsume(orderId int64, requestId string, quantity int, userId string) (*BaseResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"biz_order_id": strconv.FormatInt(orderId, 10),
		"request_id":   requestId,
		"quantity":     strconv.Itoa(quantity),
		"userid":       userId,
	}
	body, err := http.Post("https://oapi.dingtalk.com/topapi/appstore/internal/order/consume", params, "")
	resp := BaseResp{}
	if err != nil {
		return nil, err
	}
	json.FromJson(body, &resp)
	return &resp, err
}

//获取未处理的已支付订单
//doc:https://ding-doc.dingtalk.com/doc#/serverapi3/lg1nb7/2abS0
func (client *DingTalkClient) GetUnfinishOrderList(itemCode string, page int64, size int64) (GetUnfinishOrderListResp, error) {
	params := map[string]string{
		"access_token": client.AccessToken,
		"page":         strconv.FormatInt(page, 10),
		"page_size":    strconv.FormatInt(size, 10),
	}
	if itemCode != "" {
		params["item_code"] = itemCode
	}
	body, err := http.Get("https://oapi.dingtalk.com/topapi/appstore/internal/unfinishedorder/list", params)
	resp := GetUnfinishOrderListResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
