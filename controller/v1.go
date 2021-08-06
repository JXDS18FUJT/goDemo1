package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"347613781qq.com/demo1/model"
	"github.com/gin-gonic/gin"
)

type V1Test2Parm struct {
}

func V1Test(c *gin.Context) {
	// var body = strings.NewReader(c.PostForm("res"))

	response, err := http.PostForm("https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil || response.StatusCode != http.StatusOK {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "失败",
		})
		return
	}
	// reader := response.Body

	body, _ := ioutil.ReadAll(response.Body)
	// body = string(body)
	fmt.Println(string(body))
	c.JSON(200, gin.H{
		"res": string(body),
	})
	defer response.Body.Close()

	// reader := response.Body
	// contentLength := response.ContentLength
	// contentType := response.Header.Get("Content-Type")

}

func V1Test2(c *gin.Context) {
	// bytexml, err := xml.Marshal(&requestxml)

	bytexml := []byte(`<xml>
	<mch_appid></mch_appid>
	<mchid></mchid>
	<device_info></device_info>
	<nonce_str></nonce_str>
	<sign></sign>
	<partner_trade_no></partner_trade_no>
	<openid></openid>
	<check_name></check_name>
	<re_user_name></re_user_name>
	<amount></amount>
	<desc></desc>
	<spbill_create_ip></spbill_create_ip>
	</xml>`)
	client := &http.Client{}
	url := "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytexml))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	// now POST it
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"res": string(body),
	})

	defer resp.Body.Close()

}
func V1Gorou(c *gin.Context) {
	var nong model.Nong
	c.BindJSON(&nong)
	model.CreateNong(&nong)
	c.JSON(200, gin.H{
		"message": "err",
		"code":    "200",
	})
}
