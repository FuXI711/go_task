package lesson01

import (
	"errors"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string
	Balance float64
}

type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func Run(db *gorm.DB) error {
	db.AutoMigrate(&Account{}, &Transaction{})

	db.Create(&Account{Name: "A", Balance: 1000})
	db.Create(&Account{Name: "B", Balance: 0})

	err := db.Transaction(func(tx *gorm.DB) error {
		var fromAccount Account
		if err := tx.Where("name = ?", "A").First(&fromAccount).Error; err != nil {
			return err
		}

		if fromAccount.Balance < 100 {
			return errors.New("账户余额不足")
		}

		var toAccount Account
		if err := tx.Where("name = ?", "B").First(&toAccount).Error; err != nil {
			return err
		}

		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-100).Error; err != nil {
			return err
		}
		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+100).Error; err != nil {
			return err
		}

		if err := tx.Create(&Transaction{
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        100,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
