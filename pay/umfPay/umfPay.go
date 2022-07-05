package umfPay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 联动优势
type umfRes struct {
	Amount    string `json:"amount"`
	MerId     string `json:"mer_id"`
	OrderId   string `json:"order_id"`
	OrderData string `json:"order_data"`
	Sign      string `json:"sign"`
}

func (c *Config) UmfPayApp(goodsId, goodsDesc, orderId, amount string) (umfRes, error) {
	merId := c.MER_ID
	// 请求参数
	urlStr := c.URL_STR
	data := make(map[string]string, 0)
	// 协议参数
	data1 := map[string]string{
		"service":    "pay_req",
		"charset":    "UTF-8",
		"res_format": "HTML",
		"version":    "4.0",
	}
	// 业务参数
	data2 := map[string]string{
		// 商户编号
		"mer_id":     merId,
		"notify_url": c.NOTIFY_URL,
		// 商品信息
		"goods_id":  goodsId,
		"goods_inf": goodsDesc,
		// 订单号
		"order_id": orderId,
		// 商户订单日期
		"mer_date": time.Now().Format("20060102"),
		"amt_type": "RMB",
		"amount":   amount,
		"user_ip":  c.SERVER_IP,
	}
	for k, v := range data1 {
		data[k] = v
	}
	for k, v := range data2 {
		data[k] = v
	}

	// 签名
	sign, err := GetSign(data, c.PRIVATE_KEY)
	if err != nil {
		return umfRes{}, err
	}
	data["sign"] = string(sign)
	data["sign_type"] = "RSA"

	// 请求
	header := make(map[string]string, 0)
	resp, err := GetHttp(urlStr, data, header)
	if err != nil {
		return umfRes{}, err
	}
	// 获取结果
	re := regexp.MustCompile(`CONTENT=".+"`)
	i := re.FindStringIndex(resp)
	res := resp[i[0]+9 : i[1]-1]

	// 验签
	err = UmfPayVerifySig(res, c.PRIVATE_KEY)
	if err != nil {
		return umfRes{}, err
	}

	// 签名
	signData := map[string]string{
		"amount":    amount,
		"merId":     merId,
		"orderId":   orderId,
		"orderDate": time.Now().Format("20060102"),
	}
	sign, err = GetSign(signData, c.PRIVATE_KEY)
	if err != nil {
		return umfRes{}, err
	}

	// 返回
	return umfRes{
		Amount:    amount,
		MerId:     merId,
		OrderId:   orderId,
		OrderData: time.Now().Format("20060102"),
		Sign:      string(sign),
	}, nil
}

// 联动优势签名
func GetSign(data map[string]string, PRIVATE_KEY string) ([]byte, error) {
	// 获取待签名字符串
	strSign := getSignStr(data)

	// 读取秘钥
	//privateKey,err := ioutil.ReadFile("../../../lib/pay/umfPay/private.pem")
	//if err !=nil {
	//	return nil,err
	//}
	// RSA私钥签名
	encrypt, err := RsaSign([]byte(strSign), []byte(PRIVATE_KEY))
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(encrypt)), nil
}

// 联动优势验证签名
// data: 传入参数，不能含有空格
func UmfPayVerifySig(data, PUBLIC_KEY string) error {
	// 提取待签名数据，和sign
	signed, sign, err := getSigned(data)
	if err != nil {
		return err
	}
	// sign用base64解码
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	// 读取秘钥
	//publicKey,err := ioutil.ReadFile("../../../lib/pay/umfPay/public.pem")
	//if err !=nil {
	//	return err
	//}

	// RSA公钥验签
	//err = VerRsaSign([]byte(signed),publicKey,signBytes)
	err = VerRsaSign([]byte(signed), []byte(PUBLIC_KEY), signBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取待签名数据
func getSigned(data string) (signed, sign string, err error) {
	dataMap := make(map[string]string, 0)
	keys := make([]string, 0)

	dataSli := strings.Split(data, "&")
	for _, v := range dataSli {
		// 采用=符号分割键值对，且每个v只有一组键值对
		vSli := strings.Split(v, "=")
		keys = append(keys, vSli[0])
		value := strings.Replace(v, vSli[0]+"=", "", 1)

		// urlDecode
		value1, err := url.QueryUnescape(value)
		if err == nil {
			value = strings.Replace(value1, " ", "+", -1)
		}
		// GBK
		vGBK, err := Utf8ToGbk([]byte(value))
		if err != nil {
			return "", "", errors.New("编码过程错误")
		}
		dataMap[vSli[0]] = string(vGBK)
	}

	sort.Strings(keys)
	for k, v := range keys {
		if v == "sign" {
			sign = dataMap[v]
		}
		if v != "sign_type" && v != "sign" {
			signed += v + "=" + dataMap[v]
			if k < len(keys)-1 {
				signed += "&"
			}
		}
	}
	return
}

//  "http://pay.soopay.net/spay/pay/payservice.do"
// http-get
// paramUrl 请求url
// data     请求数据
// header   请求头
func GetHttp(paramUrl string, data map[string]string, header map[string]string) (string, error) {
	Url, err := url.Parse(paramUrl)
	if err != nil {
		return "", err
	}
	// 设置请求参数
	params := url.Values{}
	for k, v := range data {
		params.Set(k, v)
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()

	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	// 设置请求头
	for k, v := range header {
		resp.Header.Set(k, v)

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

// 获取签名字符串
func getSignStr(data map[string]string) string {
	keys := make([]string, 0)
	var str string
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for k, v := range keys {
		str += v + "=" + data[v]
		if k < len(keys)-1 {
			str += "&"
		}
	}
	return str
}

//私钥签名
func RsaSign(data, privateKey []byte) ([]byte, error) {
	h := sha1.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hashed)
}

// UTF-8转GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// RSA公钥验证
func VerRsaSign(data, publicKey, signature []byte) error {
	hashed := sha1.Sum(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed[:], signature)
}
