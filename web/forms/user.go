package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile" label:"手机号"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=10" label:"密码"`
}
