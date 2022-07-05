package kuaidi100

import (
	"cymul_api_go/lib/encrypt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	key      = "XwfGrNXY9402"                     // 客户授权key
	customer = "4B94CC970025825BD87D1AED0C25643D" // 授权码，申请企业版获取
)

// 快递100 物流信息实时查询
func SynQuery(expressCom, expressNnm string) string {
	param := struct {
		Com      string `json:"com"`      //  查询的快递公司的编码，一律用小写字母
		Num      string `json:"num"`      //  查询的快递单号，单号的最大长度是32个字符
		From     string `json:"from"`     //  出发地城市，省-市-区，非必填，填了有助于提升签收状态的判断的准确率，请尽量提供
		Phone    string `json:"phone"`    //  收件人或寄件人的手机号或固话（也可以填写后四位，如果是固话，请不要上传分机号）
		To       string `json:"to"`       //  目的地城市，省-市-区，非必填，填了有助于提升签收状态的判断的准确率，且到达目的地后会加大监控频率，请尽量提供
		ResultV2 string `json:"resultv2"` //  添加此字段表示开通行政区域解析功能。0：关闭（默认），1：开通行政区域解析功能，2：开通行政解析功能并且返回出发、目的及当前城市信息
		Show     string `json:"show"`     //  返回数据格式。0：json（默认），1：xml，2：html，3：text
		Order    string `json:"order"`    //  返回结果排序方式。desc：降序（默认），asc：升序
	}{
		Com:      expressCom,
		Num:      expressNnm,
		Phone:    "",
		From:     "",
		To:       "",
		ResultV2: "1",
		Show:     "0",
		Order:    "desc",
	}

	byteParam, _ := json.Marshal(param)

	// 签名， 用于验证身份， 按param + key + customer 的顺序进行MD5加密（注意加密后字符串一定要转32位大写）， 不需要加上“+”号
	sign := strings.ToUpper(encrypt.Md5String(string(byteParam) + key + customer))

	res := post(string(byteParam), sign)
	return res
}

func post(byteParam, sign string) string {
	url := "https://poll.kuaidi100.com/poll/query.do"                                              // 实时查询请求地址
	payload := strings.NewReader("customer=" + customer + "&param=" + byteParam + "&sign=" + sign) // 拼接数据
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload) // 发送请求
	if err != nil {
		return ""
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded") // 定义请求格式
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(body)
}
