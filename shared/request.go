package shared

type ParamProduct struct {
	ProductName string `json:"product_name" form:"product_name" url:"product_name"`
	Price       int    `json:"price" form:"price" url:"price"`
	Description string `json:"description" form:"description" url:"description"`
	Quantity    int    `json:"quantity" form:"quantity" url:"quantity"`
}
