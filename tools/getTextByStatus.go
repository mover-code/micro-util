package tools

import (
	"time"
)

//获取门店状态
func GetStoreStatus(status int64) map[string]interface{} {
	text := [...]string{
		0: "禁用",
		1: "启用",
	}
	return map[string]interface{}{
		"text":  text[status],
		"value": status,
	}
}

//获取is_check中文
func GetIsCheckAttr(value int64) map[string]interface{} {
	text := [...]string{
		0: "不支持",
		1: "支持",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

//获取product_status
func GetProductStatus(value int64) map[string]interface{} {
	//10上架 20下架 30 草稿 40 待审核
	text := [...]string{
		10: "上架",
		20: "下架",
		30: "草稿",
		40: "待审核",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetProductStatusText(value int64) string {
	//10上架 20下架 30 草稿 40 待审核
	text := [...]string{
		10: "上架",
		20: "下架",
		30: "草稿",
		40: "待审核",
	}
	return text[value]
}

//获取 DeliveryStatus 发货状态(10未发货 20已发货)
func GetDeliveryStatus(value int64) map[string]interface{} {
	text := [...]string{
		10: "未发货",
		20: "已发货",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetDeliveryStatusText(value int64) string {
	text := [...]string{
		10: "未发货",
		20: "已发货",
	}
	return text[value]
}

//获取 deliveryType 配送方式(10快递配送 20上门自提 30无需物流)
func GetDeliveryType(value int64) map[string]interface{} {
	text := [...]string{
		10: "快递配送",
		20: "上门自提",
		30: "无需物流",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

//获取 订单来源(10普通 20积分 30拼团 40砍价 50秒杀 60礼包购)
func GetOrderSource(value int64) string {
	text := [...]string{
		10: "普通",
		20: "积分",
		30: "拼团",
		40: "砍价",
		50: "秒杀",
		60: "礼包购",
	}
	return text[value]
}

//获取 订单状态10=>进行中，20=>已取消，30=>已完成
func GetOrderStatus(value int64) map[string]interface{} {
	text := [...]string{
		10: "进行中",
		20: "已取消",
		21: "取消审核中",
		30: "已完成",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetOrderStatusText(value int64) string {
	text := [...]string{
		10: "进行中",
		20: "已经取消",
		30: "已完成",
	}
	return text[value]
}

//获取 付款状态(10未付款 20已付款)
func GetPayStatus(value int64) map[string]interface{} {
	text := [...]string{
		10: "未付款",
		20: "已付款",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetPayStatusText(value int64) string {
	text := [...]string{
		10: "未付款",
		20: "已付款",
	}
	return text[value]
}

//获取 支付方式(10余额支付 20微信支付)
func GetPayType(value int64) map[string]interface{} {
	text := [...]string{
		10: "余额支付",
		20: "微信支付",
		30: "支付宝支付",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

//获取 收货状态(10未收货 20已收货)
func GetReceiptStatus(value int64) map[string]interface{} {
	text := [...]string{
		10: "未收货",
		20: "已收货",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetReceiptStatusText(value int64) string {
	text := [...]string{
		10: "未收货",
		20: "已收货",
	}
	return text[value]
}

//
func GetCouponState(TotalNum, ReceiveNum, ExpireType, EndTime, isReceive int64) map[string]interface{} {
	if isReceive > 0 {
		return map[string]interface{}{
			"text":  "已领取",
			"value": 0,
		}
	}
	if TotalNum > -1 && ReceiveNum >= TotalNum {
		return map[string]interface{}{
			"text":  "已抢光",
			"value": 0,
		}
	}
	if ExpireType == 20 && (EndTime+86400) < time.Now().Unix() {
		return map[string]interface{}{
			"text":  "已过期",
			"value": 0,
		}
	}
	return map[string]interface{}{
		"text":  "",
		"value": 1,
	}
}

// 获取优惠券颜色
func GetColorAttr(value int64) map[string]interface{} {
	text := [...]string{
		10: "blue",
		20: "red",
		30: "violet",
		40: "yellow",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetColorText(value int64) string {
	text := [...]string{
		10: "blue",
		20: "red",
		30: "violet",
		40: "yellow",
	}
	return text[value]
}

// 获取优惠券类型
func GetLiveStatusTextAttr(value int64) map[string]interface{} {
	text := [...]string{
		0:   "待审核",
		100: "未通过",
		101: "直播中",
		102: "未开始",
		103: "已结束",
		104: "暂停",
		107: "已过期",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetLiveStatusText(value int64) string {
	text := [...]string{
		0:   "待审核",
		100: "未通过",
		101: "直播中",
		102: "未开始",
		103: "已结束",
		104: "暂停",
		107: "已过期",
	}
	return text[value]
}

// 获取直播状态
func GetCouponTypeText(value int64) string {
	text := [...]string{
		10: "满减券",
		20: "折扣券",
	}
	return text[value]
}
func GetCouponTypeAttr(value int64) map[string]interface{} {
	text := [...]string{
		10: "满减券",
		20: "折扣券",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// 获取 分销商提现申请状态 (10待审核 20审核通过 30驳回 40已打款)
func GetApplyStatus(value int64) map[string]interface{} {
	text := [...]string{
		10: "待审核",
		20: "审核通过",
		30: "驳回",
		40: "已打款",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// 获取  订单类型(10商城订单)
func GetOrderType(value int64) map[string]interface{} {
	text := [...]string{
		10: "商城订单",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// GetActiveStatus 获取  活动状态(1生效 0未生效)
func GetActiveStatus(value int64) map[string]interface{} {
	text := [...]string{
		1: "生效-进行中",
		0: "未生效",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

func GetTimeAttr(time int64) map[string]interface{} {
	return map[string]interface{}{
		"text":  Timestamp2Date(time),
		"value": time,
	}
}

func GetJoinStatusAttr(status, endTime, joinEndTime int64) int64 {
	if status == 0 {
		return 0
	}
	if endTime < time.Now().Unix() {
		return 0
	}
	if joinEndTime < time.Now().Unix() {
		return 0
	}
	return 1
}
func GetStatusTextAttr(status, startTime, endTime int64) string {
	if status == 0 {
		return "未生效"
	}
	if startTime > time.Now().Unix() {
		return "未开始"
	}
	if endTime < time.Now().Unix() {
		return "已结束"
	}
	if startTime < time.Now().Unix() && endTime > time.Now().Unix() {
		return "生效-进行中"
	}
	return ""
}
func GetSecKillProductStatusTextAttr(status int64) string {
	text := [...]string{
		0:  "待审核",
		10: "通过",
		20: "未通过",
	}
	return text[status]
}

// 售后类型(10退货退款 20换货 30退款)
func GetOrderRefundType(value int64) map[string]interface{} {
	text := [...]string{
		10: "退货退款",
		20: "换货",
		30: "退款",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// 售后单状态(0进行中 10已拒绝 20已完成 30已取消)
func GetOrderRefundStatus(value int64) map[string]interface{} {
	text := [...]string{
		0:  "进行中",
		10: "已拒绝",
		20: "已完成",
		30: "已取消",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// 10申请平台介入20同意30拒绝
func GetOrderRefundPlateStatus(value int64) map[string]interface{} {
	text := [...]string{
		0:  "",
		10: "申请平台介入",
		20: "同意",
		30: "拒绝",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// 商家审核状态(0待审核 10已同意 20已拒绝)
func GetOrderRefundIsAgree(value int64) map[string]interface{} {
	text := [...]string{
		0:  "待审核",
		10: "已同意",
		20: "已拒绝",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

func GetMethodAttr(value int64) map[string]interface{} {
	text := [...]string{
		10: "按件数",
		20: "按重量",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
func GetMethodText(value int64) string {
	text := [...]string{
		10: "按件数",
		20: "按重量",
	}
	return text[value]
}

func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// GetBalanceLogScene 余额变动场景(10用户充值 20用户消费 30管理员操作 40订单退款)
func GetBalanceLogScene(value int64) map[string]interface{} {
	text := [...]string{
		10: "用户充值",
		20: "用户消费 ",
		30: "管理员操作",
		40: "订单退款",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

func GetUpdatePriceAttr(value float64) string {
	if value < 0 {
		return "-"
	}
	return "+"
}

func GetPayTypeText(payType int64) string {
	text := [...]string{
		10: "余额支付",
		20: "微信支付 ",
		30: "支付宝支付",
	}
	return text[payType]
}

// GetAskStatus 提问状态(0=待审核 1=审核通过 2=审核不通过)
func GetAskStatus(value int64) map[string]interface{} {
	text := [...]string{
		0: "待审核 ",
		1: "审核通过 ",
		2: "审核不通过",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}

// GetAskCommentStatus 提问评论状态(0=待审核 1=审核通过 2=审核不通过)
func GetAskCommentStatus(value int64) map[string]interface{} {
	text := [...]string{
		0: "待审核 ",
		1: "审核通过 ",
		2: "审核不通过",
	}
	return map[string]interface{}{
		"text":  text[value],
		"value": value,
	}
}
