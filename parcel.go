package main

import (
	"database/sql"
	"fmt"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	// реализуйте добавление строки в таблицу parcel, используйте данные из переменной p

	stmt, err := s.db.Prepare("INSERT INTO parcel (client, status, address, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("error preparing insert statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.Client, p.Status, p.Address, p.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("error executing insert statement: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error retrieving last insert ID: %w", err)
	}

	return int(id), nil

	// верните идентификатор последней добавленной записи
	//return 0, nil
}

func (s ParcelStore) Get(number int) (Parcel, error) {
	// реализуйте чтение строки по заданному number
	// здесь из таблицы должна вернуться только одна строка

	// заполните объект Parcel данными из таблицы
	p := Parcel{}

	err := s.db.QueryRow("SELECT number, client, status, address, created_at FROM parcel WHERE number = ?", number).
		Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, fmt.Errorf("parcel with number %d not found", number)
		}
		return p, fmt.Errorf("error retrieving parcel: %w", err)
	}

	return p, nil
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	// реализуйте чтение строк из таблицы parcel по заданному client
	// здесь из таблицы может вернуться несколько строк

	// заполните срез Parcel данными из таблицы
	var parcels []Parcel

	rows, err := s.db.Query("SELECT number, client, status, address, created_at FROM parcel WHERE client = ?", client)
	if err != nil {
		return nil, fmt.Errorf("error retrieving parcels for client %d: %w", client, err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Parcel
		if err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		parcels = append(parcels, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return parcels, nil
}

func (s ParcelStore) SetStatus(number int, status string) error {
	// реализуйте обновление статуса в таблице parcel

	stmt, err := s.db.Prepare("UPDATE parcel SET status = ? WHERE number = ?")
	if err != nil {
		return fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, number)
	if err != nil {
		return fmt.Errorf("error executing update statement: %w", err)
	}

	return nil
}

func (s ParcelStore) SetAddress(number int, address string) error {
	// реализуйте обновление адреса в таблице parcel
	// менять адрес можно только если значение статуса registered

	parcel, err := s.Get(number)
	if err != nil {
		return err
	}

	// Only allow address change if the status is "registered"
	if parcel.Status != "registered" {
		return fmt.Errorf("cannot change address for a parcel that is not 'registered'")
	}

	// Prepare the SQL statement to update the address
	stmt, err := s.db.Prepare("UPDATE parcel SET address = ? WHERE number = ?")
	if err != nil {
		return fmt.Errorf("error preparing update statement: %w", err)
	}
	defer stmt.Close()

	// Execute the update statement
	_, err = stmt.Exec(address, number)
	if err != nil {
		return fmt.Errorf("error executing update statement: %w", err)
	}
	return nil
}

func (s ParcelStore) Delete(number int) error {
	// реализуйте удаление строки из таблицы parcel
	// удалять строку можно только если значение статуса registered

	parcel, err := s.Get(number)
	if err != nil {
		return err
	}

	if parcel.Status != "registered" {
		return fmt.Errorf("cannot delete a parcel that is not 'registered'")
	}

	stmt, err := s.db.Prepare("DELETE FROM parcel WHERE number = ?")
	if err != nil {
		return fmt.Errorf("error preparing delete statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(number)
	if err != nil {
		return fmt.Errorf("error executing delete statement: %w", err)
	}

	return nil
}
