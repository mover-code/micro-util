package aliPay

import (
	"cymul_api_go/lib/tools"
	"github.com/smartwalle/alipay/v3"
)

func DoPay(productName, orderNo string, orderId, payEndTime, appId int64) (url string, err error) {
	var aliClient *alipay.Client
	if aliClient, err = alipay.New(kAppId, kPrivateKey, IsProduction); err != nil {
		return "", err
	}
	var p = alipay.TradeAppPay{}
	p.NotifyURL = kServerDomain + "/api/pay/aliPay/notyfy"
	p.ReturnURL = kServerDomain + "/pages/order/pay-success/pay-success?order_id=" + tools.Int64ToString(orderId)
	p.Body = "body"
	p.Subject = productName
	p.OutTradeNo = orderNo
	p.PassbackParams = tools.Int64ToString(appId)
	//p.TotalAmount = tools.Float64ToString(orderInfo.PayPrice)
	p.TotalAmount = "0.01"
	p.ProductCode = "QUICK_MSECURITY_PAY"
	p.TimeExpire = tools.Timestamp2min(payEndTime)

	url, _ = aliClient.TradeAppPay(p)
	return
}
