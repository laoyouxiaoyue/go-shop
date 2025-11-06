package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
)

func main() {
	AppID := "9021000157645392"
	PrivateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCFFjZpQYHV5NdXcjcODgYkkxupHMcOGWqEzIOoGrFykLfhcKGbsAGqOSDUvWuI+vc0pU+oNJ4JqsHSKzrFx60WteCs6HJKy/rjnHerS+pkVpgcJy05r0IpG570jjKVhCiYf+/ROK2Ew9/YxA3uyfkWQTQBQqgmpWzwSkpdCN2sbrDALvx1Awml0552+kCCW0TovcACyupega7k2Ww0VnCAjNA08ngo9LPCasr3prhBJuirKpVPHBaf27qDpBzvey2VGH0E9/NMjC9M04sTTNE80J7ajwFcjdHRxAFngts+LwLVMSX+4vCUiPwd3gv3i7jjzgzqUsS6DcSWpLvDjcrNAgMBAAECggEAAJnhFQietYCbGGIDraSSkoe3kEP5Ai9LM95YmeHE+2d77SC9Gh7pYwNvCobwfWXkx/AXNANI03JZ/cEEOvBz765SnXVPTtctAuoqADQPkvRxK29h3OjVu6nMRf5+a/500HuDccZ3winAURJRncp7vYX93iOW7tXAcDVlsJXhqm2z3lqBkqpASKI2cbCeLsHTleMmaKd59Yupz/nNy4eyCB353u8i4XmMQ6fNHVURBSptE7+lgjkLHsal4qSOMK/4JIAN36q3Rs8PrIy/wkgDs3K1lJaY164u8SLx5NExoAFhwq4W5v0T1gg0HEWSd4Sb4xRGMJfWDErMhmERxRbKYQKBgQDYla5e2rOiRReOS5VRBmlU+eYRp1A1c0eHlMCJAAMQP4cvKGER9iC5oWXAqrr5DK6i2qHa/AGVvWhYjsklMdU4wGoj6CP8lcDtHJ1qtHo9DY4wwYSkQKEduWfEudgngYEU9wJkHfARN53ut5GvDoEI7oNOKkZgQX2lVxteGl4dzwKBgQCdToCt4xxrzUc9OgsZicZlzVdszALTK428YbnzWRc25xG/ZaSWGj23W8Yang/EMX1AeWv2+mjo0xVnwk/ecyHHc6VgRdTgo3mLLGLVGdRS63H5FW1zWwmSFrp7J4l0bfiCkjROyL//RKrt2nb3pKbgqQtgmD4ACSC2cgyyR9YwowKBgQC4yURDY579pKtLWYtk3O6DdM7pYvHZGKj/sVdvy9XRSBJjlrTzKSaUmnKBJmlt0XcyNQ0BFB+aDUb4v6I4HMNJr0kfp5EEkRRQeSiJgMAGVQlq6cDsjsOn46b4PoOAeUqTg4fDuZM4bVZO9MnULLDE2RKoXwM6XsX/ikksSZzhGwKBgBKz8ri59/cZQQ8Wh5tRtjUEZRCacPuKgh+Tvvgf75KnhoSrIRZ3qD7UuokPofBBshKoXR3QSAjmj/T2NWjNZ4a/STpZEyZiSWEytc2AdK4nMDXdRlYgzNKBwUpDOSSOrq1XlMCJPcqr72a4QszO1sh/UMr7TuPSMgF/LeNEh0LhAoGBANWEFNzPj2fsMI8su8YaMZ9QMO3lDdxYKnPeXhU8zVlwMkcq6+ARPhcKugW6c/r9OpjaI602WfFNvqnDtsoxqBEkc0FJsyfU4g9+MKRHDkX3XUh2dxELfgNjciACtU218H0/cV6E9k6n4Ao/Wfvzhd0eiFXM4pKvHrrcWohNH5L5"
	AlipayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhRY2aUGB1eTXV3I3Dg4GJJMbqRzHDhlqhMyDqBqxcpC34XChm7ABqjkg1L1riPr3NKVPqDSeCarB0is6xcetFrXgrOhySsv645x3q0vqZFaYHCctOa9CKRue9I4ylYQomH/v0TithMPf2MQN7sn5FkE0AUKoJqVs8EpKXQjdrG6wwC78dQMJpdOedvpAgltE6L3AAsrqXoGu5NlsNFZwgIzQNPJ4KPSzwmrK96a4QSboqyqVTxwWn9u6g6Qc73stlRh9BPfzTIwvTNOLE0zRPNCe2o8BXI3R0cQBZ4LbPi8C1TEl/uLwlIj8Hd4L94u4484M6lLEug3ElqS7w43KzQIDAQAB"
	client, err := alipay.New(AppID, PrivateKey, false)
	if err != nil {
		zap.S().Errorw("实例化支付宝失败")
		return
	}
	err = client.LoadAliPayPublicKey(AlipayPublicKey)
	if err != nil {
		zap.S().Errorw("加载支付宝的公钥失败")
		return
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://xxx"
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, errS = client.TradePagePay(p)
	if errS != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
}
