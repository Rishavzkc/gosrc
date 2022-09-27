package Models

type Book struct {
	Name   string  `json:"name"`
	Id     uint    `json:"id"`
	Author string  `json: "author"`
	price  float64 `json:"price"`
}

func (b *Book) TableName() string {
	return "book"
}
