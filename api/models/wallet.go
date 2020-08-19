package models

import (
	"database/sql"
	"fmt"
)

type Wallet struct {
	ID int `json:"id"`
	Amount int `json:"amount"`
	UserID uint `json:"user_id"`
	CreatedBy User `json:"-"`
}

func (w *Wallet) Prepare() {
	w.CreatedBy = User{}
}


func(w *Wallet) DepositMoney(db *sql.DB) (Wallet, error) {

	var amount float32
	err := db.QueryRow("update wallet set amount =amount + $1 where id =$2 RETURNING amount", w.Amount, w.ID).Scan(&amount)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{ID: w.ID, Amount: int(amount)}, nil
}

func(w *Wallet) WithdrawMoney(db *sql.DB) (Wallet, error) {
	var amount float32
	err := db.QueryRow("update wallet set amount =amount - $1 where id =$2 RETURNING amount", w.Amount, w.ID).Scan(&amount)
	if err != nil {
		return Wallet{}, err
	}
	return Wallet{ID: w.ID, Amount: int(amount)}, nil
}

func (w *Wallet) CreateWallet(db *sql.DB) error {
	err := db.QueryRow("insert into wallet(amount) values($1) returning id", w.Amount).Scan(&w.ID)

	if err != nil {
		return err
	}
	return nil
}

func(w *Wallet) CheckBalance(db *sql.DB) error {
	return db.QueryRow("select amount from wallet where id = $1", w.ID).Scan(&w.Amount)
}

func (w *Wallet) getAccounts(db *sql.DB, start, count int) ([]Wallet, error) {
	rows, err := db.Query("select id, amount from wallet LIMIT $1 OFFSET $2", count, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	wallets := []Wallet{}

	for rows.Next() {
		var w Wallet
		if err := rows.Scan(&w.ID,&w.Amount); err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}
	if err := rows.Err(); err != nil {
		return wallets, fmt.Errorf("unable to iterate: %v", err)
	}

	return wallets, nil
}

