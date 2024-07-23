package aliyun

import (
	"SMS/app/util"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	tea "github.com/alibabacloud-go/tea/tea"
)

var SMSconfig = &openapi.Config{
	// 您的AccessKey ID
	AccessKeyId: tea.String("LTAI5tM8XhgWyctbXKWngxJC"),
	// 您的AccessKey Secret
	AccessKeySecret: tea.String("bv9xHxEdhq2O0QtMCI1etPLv2FQ1TD"),
}

func SMSAliyun() {
	// 访问的域名
	SMSconfig.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	client, _ := dysmsapi.NewClient(SMSconfig)

	util.Logger.Println(client)

	request := &dysmsapi.AddShortUrlRequest{}

	util.Logger.Println(request)
}
