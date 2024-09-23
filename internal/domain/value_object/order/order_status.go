package order_vo

const (
	NEW        OrderStatus = "new"
	WORKING_ON OrderStatus = "working_on"
	READY      OrderStatus = "ready"
)

type OrderStatus string
