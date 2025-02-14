package products

//# Товары (структура)
//Поля структуры Товары:
//- ID товара (int)
//- Поле для привязки товара к категории по ID (int)
//- название товара (string)
//- ссылка на картинку (картинка) //в бд наверное лучше сохранять ссылку на картинку (string)
//- цена (int)
//- количество товара (остаток на складе) (int)
//- описание товара (string)

type Product struct {
	ProductID         int64
	ProductIDcategory int64
	ProductName       string
	ProductPictureUrl string
	ProductPrice      int64
	ProductQuantity   int64
	ProductDescripion string
}
