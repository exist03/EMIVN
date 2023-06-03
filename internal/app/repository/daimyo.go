package repository

import (
	"log"
	"time"
)

func (r *Repository) DaimyoGetSamuraiAmountList(daimyo string) (map[string]Banks, error) {
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	myMap := make(map[string]Banks)
	stmtSber := `SELECT SamuraiUsername, Amount 
	FROM SamuraiTurnover
	JOIN Samurai S ON SamuraiTurnover.SamuraiUsername = S.TelegramUsername
	WHERE S.Owner=? AND Bank="сбер" AND Date=?`
	stmtTink := `SELECT SamuraiUsername, Amount 
	FROM SamuraiTurnover
	JOIN Samurai S ON SamuraiTurnover.SamuraiUsername = S.TelegramUsername
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
func (r *Repository) DaimyoGetReportByPeriod(daimyo string, dateBegin, dateEnd time.Time) (map[string]InfoPeriod, error) {
	myMap := make(map[string]InfoPeriod)
	stmtTink := `SELECT SamuraiUsername, Amount
	FROM SamuraiTurnover
	JOIN Samurai S ON SamuraiTurnover.SamuraiUsername = S.TelegramUsername
	WHERE S.owner=? AND Bank="тинькофф" AND Date=?`
	stmtSber := `SELECT SamuraiUsername, Amount
	FROM SamuraiTurnover
	JOIN Samurai S ON SamuraiTurnover.SamuraiUsername = S.TelegramUsername
	WHERE S.owner=? AND Bank="сбер" AND Date=?`
	rows, err := r.DB.Query(stmtTink, daimyo, dateBegin)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var samurai string
		var amount float64
		err = rows.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.Begin = amount
			myMap[samurai] = str
		} else {
			s := InfoPeriod{
				Begin: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	rowsSber, err := r.DB.Query(stmtSber, daimyo, dateBegin)
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
		if str, ok := myMap[samurai]; ok {
			str.Begin += amount
			myMap[samurai] = str
		} else {
			s := InfoPeriod{
				Begin: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsSber.Err(); err != nil {
		return nil, err
	}

	rowsTinkEnd, err := r.DB.Query(stmtTink, daimyo, dateEnd)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsTinkEnd.Close()

	for rowsTinkEnd.Next() {
		var samurai string
		var amount float64
		err = rowsTinkEnd.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.End = amount
			myMap[samurai] = str
		} else {
			s := InfoPeriod{
				End: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsTinkEnd.Err(); err != nil {
		return nil, err
	}

	rowsSberEnd, err := r.DB.Query(stmtSber, daimyo, dateEnd)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rowsSberEnd.Close()

	for rowsSberEnd.Next() {
		var samurai string
		var amount float64
		err = rowsSberEnd.Scan(&samurai, &amount)
		if err != nil {
			return nil, err
		}
		if str, ok := myMap[samurai]; ok {
			str.End += amount
			myMap[samurai] = str
		} else {
			s := InfoPeriod{
				End: amount,
			}
			myMap[samurai] = s
		}
	}
	if err = rowsSberEnd.Err(); err != nil {
		return nil, err
	}

	return myMap, err
}
