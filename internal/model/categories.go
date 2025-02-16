package categories

//# Категория товара (структура)
//Поля структуры категория:
//- название категории (string)
//- ID категории (int)

type Category struct {
	ID        int64
	Name      string
	Categorys map[int64]*Category
	Products  map[int64]*Product
}
