package form

type Product struct {
	Code string
	Name string
	Type string
	ImageUrl string
	Description string
}

type GetProduct struct {
	Limit int
	Offset int
	Code string
	ProductId uint
}
