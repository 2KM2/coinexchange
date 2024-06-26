package domain

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type vaptchaReq struct {
	Id        string `json:"id"`
	Secretkey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}
type CaptchaDomain struct {
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}

// Verify 验证函数
func (d *CaptchaDomain) Verify(
	server string,
	vid string,
	key string,
	token string,
	scene int,
	ip string) bool {
	logx.Info("[CaptchaDomain]")
	//1.发送Post请求
	resp, err := tools.Post(server, &vaptchaReq{
		Id:        vid,
		Secretkey: key,
		Token:     token,
		Scene:     scene,
		Ip:        ip,
	})
	logx.Info("[CaptchaDomain] resp ", resp)
	if err != nil {
		logx.Error(err)
		return false
	}
	result := &vaptchaRsp{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		logx.Error(err)
		return false
	}
	return result.Success == 1
}
