package domain

type ProductRepo interface {
	ListProducts() ([]*Product, error)
}
