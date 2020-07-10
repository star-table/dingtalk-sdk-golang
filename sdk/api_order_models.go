package sdk

type GetSkuPageResp struct {
	BaseResp
	Result string `json:"result"`
}

type GetOrderInfoResp struct {
	BaseResp
	Result OrderInfo `json:"result"`
}

type OrderInfo struct {
	CreateTimestamp   int64  `json:"create_timestamp"`
	StartTimestamp    int64  `json:"start_timestamp"`
	EndTimestamp      int64  `json:"end_timestamp"`
	PaidTimestamp     int64  `json:"paid_timestamp"`
	Quantity          int    `json:"quantity"`
	Status            int    `json:"status"`
	TotalActualPayFee int64  `json:"total_actual_pay_fee"`
	GoodsCode         string `json:"goods_code"`
	ItemCode          string `json:"item_code"`
	CorpId            string `json:"corp_id"`
	BizOrderId        int64  `json:"biz_order_id"`
}

type GetUnfinishOrderListResp struct {
	BaseResp
	Result UnfinishOrderList `json:"result"`
}

type UnfinishOrderList struct {
	Items []UnfinishOrderInfo `json:"items"`
}

type UnfinishOrderInfo struct {
	BizOrderId        int64  `json:"biz_order_id"`
	CorpId            string `json:"corp_id"`
	CreateTimestamp   int64  `json:"create_timestamp"`
	GoodsCode         string `json:"goods_code"`
	ItemCode          string `json:"item_code"`
	PaidTimestamp     int64  `json:"paid_timestamp"`
	Quantity          int    `json:"quantity"`
	Status            int    `json:"status"`
	TotalActualPayFee int64  `json:"total_actual_pay_fee"`
}
