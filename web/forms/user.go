package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile" label:"手机号"`
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=10" label:"密码"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5" `
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile" label:"手机号"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10" label:"密码"`
	Code     string `form:"code" json:"code" binding:"required,min=2,max=10"`
	Subject  string `form:"subject" json:"subject" binding:"required"`
}
