package forms

// 购物车项表单
type CartItemForm struct {
	Id         int32   `json:"id"`
	UserId     int32   `json:"userId" binding:"required,min=1"`
	GoodsId    int32   `json:"goodsId" binding:"required,min=1"`
	GoodsName  string  `json:"goodsName"`
	GoodsImage string  `json:"goodsImage"`
	GoodsPrice float32 `json:"goodsPrice"`
	Nums       int32   `json:"nums" binding:"required,min=1"`
	Checked    bool    `json:"checked"`
}

// 订单创建表单
type CreateOrderForm struct {
	Id      int32  `json:"id"`
	UserId  int32  `json:"userId" binding:"required,min=1"`
	Address string `json:"address" binding:"required,min=1"`
	Name    string `json:"name" binding:"required,min=1"`
	Mobile  string `json:"mobile" binding:"required,min=6"`
	Post    string `json:"post"`
}

// 订单筛选表单
type OrderFilterForm struct {
	UserId      int32 `form:"userId" json:"userId" binding:"required,min=1"`
	Pages       int32 `form:"pages" json:"pages" binding:"min=1"`
	PagePerNums int32 `form:"pagePerNums" json:"pagePerNums" binding:"min=1,max=100"`
}

// 订单状态更新表单
type UpdateOrderStatusForm struct {
	Id     int32  `json:"id" binding:"required,min=1"`
	Status string `json:"status" binding:"required,min=1"`
}
