package value_object

const (
	NEW        OrderStatus = "new"
	WORKING_ON OrderStatus = "working"
	READY      OrderStatus = "ready"
)

type OrderStatus string
