package models

import (
	"github.com/jinzhu/gorm"
)

type Wallet struct {
	gorm.Model
	ID  		uint 			`gorm:"not null;unique"     json:"id"`
	Amount 		int 		`gorm:"not null"      		json:"amount"`
	CreatedBy   User	        `gorm:"foreignKey:UserID;"  json:"_"`
	UserID 		uint  			`gorm: "not null"        	json:"user_id"`
}

func (w *Wallet) Prepare() {
	w.CreatedBy = User{}
}

func (w *Wallet) deposit(db *gorm.DB) (Wallet, error) {
	var amount float64
	err := db.Exec("update wallet set amount = amount + $1 where id = $2 RETURNING amount", w.Amount, w.ID).
		Scan(&amount)
	if err != nil {
		return Wallet{}, err.Error
	}
	return Wallet{ID: w.ID, Amount: int(w.Amount)}, nil
}

func (w *Wallet) withdraw(db *gorm.DB) (Wallet, error) {
	var amount float64
	err := db.Exec("update wallet set amount =amount - $1 where id = $2 RETURNING amount", w.Amount, w.ID).
		Scan(&amount)
	if err != nil {
		return Wallet{}, err.Error
	}
	return Wallet{ID: w.ID, Amount: int(amount)}, nil
}

func (w *Wallet) create(db *gorm.DB)  error {
	err := db.Exec("insert into wallet(amount) values($1) RETURNING id", w.Amount).Scan(&w.ID)

	if err != nil {
		return err.Error
	}
	return nil

}


func(w *Wallet) GetWalletById(id int, db *gorm.DB) (*Wallet, error) {
	wallet := &Wallet{}
	if err := db.Debug().Table("wallet").Where("id = ?", id).First(wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func(w *Wallet) GetWallet(db *gorm.DB) (*[]Wallet, error) {
	venues := []Wallet{}
	if err := db.Debug().Table("venues").Find(&venues).Error; err != nil {
		return &[]Wallet{}, err
	}
	return &venues, nil
}

func(w *Wallet) DeleteWallet(id int, db *gorm.DB) error {
	if err := db.Debug().Table("venues").Where("id = ?", id).Delete(&Wallet{}).Error; err != nil {
		return err
	}
	return nil
}

func (v *Wallet) Save(db *gorm.DB) (*Wallet, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&v).Error
	if err != nil {
		return &Wallet{}, err
	}
	return v, nil
}