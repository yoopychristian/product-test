package shared

type Product struct {
	IDProduct   string `json:"id_product"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
