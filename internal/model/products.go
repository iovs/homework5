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
	ID         int64
	Name       string
	PictureUrl string
	Price      int64
	Quantity   int64
	Descripion string
}
