package repository

import (
	"log"
	"time"
)

func (r *Repository) ControllerEnterInfo() {

}
func (r *Repository) ControllerGetSamuraiAmountList(daimyo string) (map[string]Banks, error) {
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	myMap := make(map[string]Banks)
	stmtSber := `SELECT SamuraiUsername, Amount 
	FROM TurnoverController
	JOIN Samurai S ON TurnoverController.SamuraiUsername = S.TelegramUsername
	WHERE S.Owner=? AND Bank="сбер" AND Date=?`
	stmtTink := `SELECT SamuraiUsername, Amount 
	FROM TurnoverController
	JOIN Samurai S ON TurnoverController.SamuraiUsername = S.TelegramUsername
	WHERE S.Owner=? AND Bank="тинькофф" AND Date=?`

	rowsSber, err := r.DB.Query(stmtSber, daimyo, yesterday)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsSber.Close()

	for rowsSber.Next() {
		var samurai string
		var amount float64
		err = rowsSber.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		s := Banks{
			Sber: amount,
		}
		myMap[samurai] = s
	}
	if err = rowsSber.Err(); err != nil {
		return nil, err
	}

	rowsTink, err := r.DB.Query(stmtTink, daimyo, yesterday)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTink.Close()

	for rowsTink.Next() {
		var samurai string
		var amount float64
		err = rowsTink.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.Tinkoff = amount
			myMap[samurai] = str
		} else {
			s := Banks{
				Tinkoff: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsTink.Err(); err != nil {
		return nil, err
	}
	return myMap, nil
}
