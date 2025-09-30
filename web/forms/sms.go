package forms

type SendSmsForm struct {
	Mobile  string `form:"mobile" json:"mobile"   binding:"required,mobile"`
	Subject string `form:"subject" json:"subject"  binding:"required"`
}
