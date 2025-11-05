package forms

// 设置库存表单
type SetInventoryForm struct {
	GoodsId int32 `json:"goodsId" binding:"required,min=1"`
	Num     int32 `json:"num" binding:"required,min=0"`
}

// 批量扣减/归还库存表单
type InventoryBatchForm struct {
	GoodsInfo []InventoryItem `json:"goodsInfo" binding:"required,min=1,dive"`
	OrderSn   string          `json:"orderSn" binding:"required,min=1"`
}

type InventoryItem struct {
	GoodsId int32 `json:"goodsId" binding:"required,min=1"`
	Num     int32 `json:"num" binding:"required,min=1"`
}
