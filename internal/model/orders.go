package order

import "time"

//# Заказ (структура)
//Поля струтктуры Заказ:
//- имя (string)
//- ID номер заказа (int)
//- количество товара в заказе (int)
//- дата и время (time.Time)
//- цена (int)
//- итоговая сумма (int)
//- наименование товара (string)
//- статус заказа (принят, отправлен, завершен) (string)
//- id пользователя (string)
//- телефон (string)
//- адрес доставки //не знаю получится сделать или нет, тк по уму там надо города списком выводить, подтягивать кладер адресов и городов (string)

type Order struct {
	OrderID             int64
	OrderQuantity       int64
	OrderDateTime       time.Time
	OrderPrice          int64
	OrderStatus         string
	OrderDeliveryAdress string
}
