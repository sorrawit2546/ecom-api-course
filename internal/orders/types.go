package orders

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity	int32	`json:"quantity"` 
}

type createOrderParams	struct {
	CustomerId	int64	`json:"customer_id"`
	Items	[]orderItem	`json:items`
}