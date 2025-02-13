Веб-приложение "Мини-приложение для Telegram витрина товаров"

Описание:

Создать веб-приложение Telegram для отображения каталога товаров интернет-магазина с ценой, возможностью заказа товара.

Требования:

Веб-интерфейс отображает список товаров интернет-магазина и (категории товаров - вот это под вопросом пока). Есть возможность заказать товар.
Данные о заказе: номер заказа, дата, время, сумма, наименование товара, статус заказа, данные пользователя сохраняются в базе данных, чтобы они были доступны между сессиями.
Данные о товарах: название, ссылка на картинку, цена, количество, возможно описание - сохраняются в базе данных, чтобы они были доступны между сессиями.
Реализована аутентификация пользователей, чтобы каждый пользователь видел только свою историю заказов.
Реализована возможность импорта администратором данных о товарах из формата json/yml/xls/xlsx/csv.
Реализована административная панель, позволяющая отслеживать заказы.
Применение шаблонов HTML для отображения пользовательского интерфейса.
Обработка ошибок и валидация пользовательского ввода, чтобы предотвратить некорректные данные.
Развертывание

Развертывание сервиса должно осуществляться с использованием docker compose в директории с проектом.

Тестирование

Написаны юнит-тесты на core логику приложения. Плюсом будут тесты на транспортном уровне и на уровне хранения.

Критерии оценивания

Максимум - 15 баллов (при условии выполнения обязательных требований):

Реализован алгоритм - 2 балла.
Реализовано разделение на слои (транспортный, хранения и т.д.) - 2 балла.
Реализовано API сервиса - 2 балла.
Реализован пользовательский интерфейс - 2 балла.
Написаны юнит-тесты - 1 балл.
Написаны интеграционные тесты - 2 балла.
Тесты адекватны и полностью покрывают функциональность - 1 балл.
Понятность и чистота кода - до 3 баллов.
Зачёт от 10 баллов# homework5

Структура приложения

# Пользователь
Категория товара
- название категории
- название подкатегории // какая-то привязка нужна к основной категории для реализации вложенности второго-третего уровня

Товары
- название товара (string)
- ссылка на картинку (картинка) //в бд наверное лучше сохранять ссылку на картинку (string)
- цена (int)
- количество (int)
- описание товара (string)

Заказ
- имя (string)
- номер заказа (int)
- количество (int)
- дата // не знаю делать просто string или int string int string int (int.int.int)
- время // не знаю делать просто string или int string int (int:int)
- цена (int)
- итоговая сумма (int)
- наименование товара (string)
- статус заказа (принят, отправлен, завершен) (string)
- id пользователя //возможно имя (string)
- телефон (string)
- адрес доставки //не знаю получится сделать или нет, тк по уму там надо города списком выводить, подтягивать кладер адресов и городов (string)

# Админка приложения
Ссылка на файл/feed
Возможность загрузки файла

# Поля товаров

# Поля в заказе







