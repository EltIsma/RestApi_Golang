package lavka

//import "time"

type Orders struct {
	Id             int      `json:"order_id" `
	Weight         float32  `json:"weight"  binding:"required"`
	Cost           int      `json:"cost" binding:"required"`
	Region         int      `json:"regions" binding:"required"`
	DeliveryHours  string `json:"delivery_hours"  binding:"required"`
	CompleteOrders string   `json:"complete_time,omitempty" `
}

type OrdersComplete struct {
	Id           int    `json:"completeOrder_id"`
	CurierId     int    `json:"courier_id"  binding:"required"`
	OrderId      int    `json:"order_id" binding:"required"`
	CompleteTime string `json:"complete_time"  binding:"required"`
}

type OrdersAssign struct {
	Id           int     `json:"order_id"`
	Weight       float32 `json:"weight"  binding:"required"`
	Region       int     `json:"region" binding:"required"`
	Cost         int     `json:"cost"  binding:"required"`
}
