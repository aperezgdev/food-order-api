package value_object

type Price float32

func NewPrice(price float32) Price {
	return Price(price)
}

func (p *Price) Validate() bool {
	return *p > 0.0
}
