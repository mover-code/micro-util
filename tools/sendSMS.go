package tools

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type SmsInfo struct {
	AccessKeyId        string `json:"AccessKeyId"`
	AccessKeySecret    string `json:"AccessKeySecret"`
	SignName           string `json:"sign"`
	TemplateCode       string `json:"login_template"`
	TemplateParam      string `json:"apply_template"`
	SupplierRejectCode string `json:"supplier_reject_code"`
	SupplierPassCode   string `json:"supplier_pass_code"`
}

// 阿里云发送短信
func AliSendSMS(sms SmsInfo, phoneNumber, code string) error {
	config := &openapi.Config{
		AccessKeyId:     tea.String(sms.AccessKeyId),     // AccessKey ID
		AccessKeySecret: tea.String(sms.AccessKeySecret), // AccessKey Secret
	}

	config.Endpoint = tea.String("dysmsapi.aliyuncs.com") // 短信服务器
	client, errNewClient := dysmsapi20170525.NewClient(config)
	if errNewClient != nil {
		return errNewClient
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{ // 发送短信参数
		PhoneNumbers:  tea.String(phoneNumber),
		SignName:      tea.String(sms.SignName),
		TemplateCode:  tea.String(sms.TemplateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	resp, errSendSms := client.SendSms(sendSmsRequest)
	if errSendSms != nil {
		return errSendSms
	}

	if *resp.Body.Message != "OK" {
		return errors.New(*resp.Body.Message)
	}
	return nil
}
