package repository

import (
	"emivn/internal/models"
	"log"
)

func (r *Repository) CardGetListByOwner(daimyo string) ([]models.Card, error) {
	cards := make([]models.Card, 0)
	stmt := "SELECT ID, BankInfo, LimitInfo, InDispute, BalanceInfo FROM Card WHERE Owner=?"
	rows, err := r.DB.Query(stmt, daimyo)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := models.Card{}
		err = rows.Scan(&temp.ID, &temp.Bank, &temp.Limit, &temp.InDispute, &temp.Balance)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cards = append(cards, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cards, nil
}
