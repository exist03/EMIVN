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
		temp := models.Card{Owner: daimyo}
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
func (r *Repository) CardGetAll() ([]models.Card, error) {
	cards := make([]models.Card, 0)
	stmt := "SELECT ID, BankInfo, LimitInfo, InDispute, BalanceInfo, Owner FROM Card"
	rows, err := r.DB.Query(stmt)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := models.Card{}
		err = rows.Scan(&temp.ID, &temp.Bank, &temp.Limit, &temp.InDispute, &temp.Balance, &temp.Owner)
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
func (r *Repository) CardUpdateBalance(cardID string, balance int) error {
	stmt := "UPDATE Card SET BalanceInfo=? WHERE ID=?"
	_, err := r.DB.Exec(stmt, balance, cardID)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
func (r *Repository) CardInDispute(cardID string) bool {
	var inDispute bool
	result := r.DB.QueryRow("SELECT InDispute FROM Card WHERE ID=?", cardID)
	result.Scan(&inDispute)
	return inDispute
}
func (r *Repository) CardGetByID(cardID string) (models.Card, error) {
	card := models.Card{ID: cardID}
	stmt := "SELECT BankInfo, LimitInfo, InDispute, BalanceInfo, Owner FROM Card WHERE ID=?"
	row := r.DB.QueryRow(stmt, cardID)
	err := row.Scan(&card.Bank, &card.Limit, &card.InDispute, &card.Balance, &card.Owner)
	if err != nil {
		return models.Card{}, err
	}
	return card, nil
}
func (r *Repository) CardSetDisputeTrue(cardID string) error {
	stmt := "UPDATE Card SET InDispute=1 WHERE ID=?"
	_, err := r.DB.Exec(stmt, cardID)
	if err != nil {
		return err
	}
	return nil
}
