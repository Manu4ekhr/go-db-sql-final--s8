package main

import (
	"database/sql"
	"fmt"
)

// Структура для представления посылки
type Parcel struct {
	Number    int
	Client    int
	Status    string
	Address   string
	CreatedAt string
}

// Структура для работы с базой данных
type ParcelStore struct {
	db *sql.DB
}

// Конструктор для создания нового ParcelStore
func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

// Метод для добавления новой посылки в базу данных
func (s ParcelStore) Add(p Parcel) (int, error) {
	// Реализация добавления в БД (пример без реального SQL-запроса)
	fmt.Printf("Добавляем посылку: %+v\n", p)
	return 1, nil // Возвращаем ID добавленной посылки
}

// Метод для получения посылки по её номеру
func (s ParcelStore) Get(number int) (Parcel, error) {
	p := Parcel{
		Number:    number,
		Client:    1,
		Status:    "registered",
		Address:   "Address Example",
		CreatedAt: "2024-11-15T10:00:00Z",
	}
	// Реализация получения из БД (пока статические данные)
	fmt.Printf("Получаем посылку с номером: %d\n", number)
	return p, nil
}

// Метод для получения всех посылок клиента
func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	parcels := []Parcel{
		{Number: 1, Client: client, Status: "sent", Address: "Address 1", CreatedAt: "2024-11-01T10:00:00Z"},
		{Number: 2, Client: client, Status: "registered", Address: "Address 2", CreatedAt: "2024-11-02T11:00:00Z"},
	}
	// Реализация получения посылок по клиенту из БД
	fmt.Printf("Получаем посылки для клиента: %d\n", client)
	return parcels, nil
}

// Метод для изменения статуса посылки
func (s ParcelStore) SetStatus(number int, status string) error {
	// Реализация изменения статуса в БД
	fmt.Printf("Изменяем статус посылки с номером %d на %s\n", number, status)
	return nil
}

// Основная функция
func main() {
	// Настроим подключение к базе данных (симуляция)
	db := &sql.DB{} // Здесь просто заглушка, для реальной работы нужно подключение к БД

	store := NewParcelStore(db)

	// Пример использования методов
	parcel := Parcel{
		Client:    1,
		Status:    "registered",
		Address:   "Address Example",
		CreatedAt: "2024-11-15T10:00:00Z",
	}

	// Добавление посылки
	id, err := store.Add(parcel)
	if err != nil {
		fmt.Println("Ошибка при добавлении посылки:", err)
		return
	}
	fmt.Printf("Посылка добавлена с номером: %d\n", id)

	// Получение посылки по номеру
	p, err := store.Get(id)
	if err != nil {
		fmt.Println("Ошибка при получении посылки:", err)
		return
	}
	fmt.Printf("Получена посылка: %+v\n", p)

	// Изменение статуса посылки
	err = store.SetStatus(id, "sent")
	if err != nil {
		fmt.Println("Ошибка при изменении статуса посылки:", err)
		return
	}

	// Получение всех посылок клиента
	parcels, err := store.GetByClient(1)
	if err != nil {
		fmt.Println("Ошибка при получении посылок клиента:", err)
		return
	}
	fmt.Println("Все посылки клиента 1:", parcels)
}
