package database

import (
	"fmt"

	"github.com/tommy-sho/golang-sandbox/web/payjp-vue/backend-api/domain"
)

func SelectAllItems() (items domain.Items, err error) {
	stmt, err := Conn.Query("SELECT * FROM items")
	if err != nil {
		return
	}
	defer stmt.Close()

	for stmt.Next() {
		var id int64
		var name string
		var description string
		var amount int64
		if err := stmt.Scan(&id, &name, &description, &amount); err != nil {
			continue
		}
		item := domain.Item{
			ID:          id,
			Name:        name,
			Description: description,
			Amount:      amount,
		}
		items = append(items, item)
	}
	return
}

func SelectItem(identifier int64) (domain.Item, error) {
	stmt, err := Conn.Prepare(fmt.Sprintf("SELECT * FROM items WHERE id = ? LIMIT 1"))
	if err != nil {
		return domain.Item{}, fmt.Errorf("SelectItems error: %v", err)
	}
	defer stmt.Close()
	var id int64
	var name string
	var description string
	var amount int64
	err = stmt.QueryRow(identifier).Scan(&id, &name, &description, &amount)
	if err != nil {
		return domain.Item{}, fmt.Errorf("SelectItems error: %v", err)
	}

	return domain.Item{
		ID:          id,
		Name:        name,
		Description: description,
		Amount:      amount,
	}, nil
}
