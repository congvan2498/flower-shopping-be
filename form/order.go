package form

type CreateOrderRequest struct {
	Code string
	UserId int
	Status string
	Address Address
	Product []Product
}

type Address struct {
	OrderId uint
	Address string
	Name string
	WardName string
	DistrictName string
	ProvinceName string
	Code string
	Phone string
}
