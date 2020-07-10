package sdk

import (
	"testing"
)

func TestDingTalkClient_GetOrderInfo(t *testing.T) {
	client := CreateClient()
	res, err := client.GetOrderInfo(50010505110088)
	t.Log(err)
	t.Log(res)
}

func TestDingTalkClient_GetSkuPage(t *testing.T) {
	client := CreateClient()
	res, err := client.GetSkuPage("FW_GOODS_1111", "https://apptest.bjx.cloud")
	t.Log(err)
	t.Log(res)
}

func TestDingTalkClient_OrderFinish(t *testing.T) {
	client := CreateClient()
	res, err := client.OrderFinish(50010505110088)
	t.Log(err)
	t.Log(res)
}

func TestDingTalkClient_OrderConsume(t *testing.T) {
	client := CreateClient()
	res, err := client.OrderConsume(50010505110088, "1199291922", 1, "12")
	t.Log(err)
	t.Log(res)
}
