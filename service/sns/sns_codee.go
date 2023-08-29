package sns

import (
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/common/constant"
	"chatgpt-web-new-go/common/email"
	"chatgpt-web-new-go/common/env"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/random"
	"chatgpt-web-new-go/common/redis"
	"chatgpt-web-new-go/common/regexp"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	whiteListPhoneCode = "0000"
)

func SendCode(source string) error {
	// 0. valid first
	if !regexp.IsValidPhoneOrEmail(source) {
		return fmt.Errorf("只支持手机号或者邮箱！")
	}

	// 1. generate sns code, eg:1234
	snsCode := random.GenSmsCode()

	// 2. store code
	key := fmt.Sprintf(redis.KeySnsCode, source)
	if err := config.Redis.Set(key, snsCode, 20*time.Minute).Err(); err != nil {
		return err
	}

	// 3. type based
	if regexp.IsValidPhone(source) {
		return apiSendSnsCode(source, snsCode)
	} else {
		return SendEmailCode(source, snsCode)
	}
}

func SendEmailCode(emailAddress string, code string) error {
	emailContent := fmt.Sprintf(email.SendCodeTemplate, code)
	err := email.SendMail("OurAI", emailAddress, emailContent)
	if err != nil {
		logs.Error("mail.SendMail emailAddress:%v, bizError: %v", emailAddress, err)
	}
	return err
}

type apiResponse struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	SmUuid string `json:"sm_uuid"`
}

func apiSendSnsCode(phone, code string) error {
	if _, found := constant.WhiteListPhone[phone]; found || env.IsDevelop() {
		return nil
	}

	data := url.Values{}

	data.Set("accesskey", "1QQRVwDIBvQ3nSnd")
	data.Set("secret", "Ib87ISCYEfYbsF51WGkXAdt61HhrmW7V")
	data.Set("signin", "【云上AI】")
	data.Set("templateId", "182578")
	data.Set("mobile", phone)
	data.Set("content", code)

	resp, err := http.PostForm("https://api.1cloudsp.com/api/v2/single_send", data)
	if err != nil {
		logs.Error("http.PostForm erorr: %v", err)
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("ioutil.ReadAll erorr: %v", err)
		return err
	}

	logs.Info("sns send api response: %v", string(content))

	result := &apiResponse{}
	err = json.Unmarshal(content, result)
	if err != nil || result == nil {
		logs.Error("json.Unmarshal erorr: %v， content:%v", err, string(content))
		return err
	}

	if result.Code != "0" {
		logs.Error("result code not 0: %v, msg: %v", result.Code, result.Msg)
		return fmt.Errorf("result code not 0: %v, msg: %v", result.Code, result.Msg)
	}
	return nil
}
