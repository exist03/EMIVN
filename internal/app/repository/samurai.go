package repository

import (
	"emivn/internal/models"
	"log"
)

func (r *Repository) SamuraiInsert(samurai models.Samurai) error {
	stmt := `INSERT INTO Samurai (Nickname, Owner, TelegramUsername) VALUES(?, ?, ?)`
	_, err := r.DB.Exec(stmt, samurai.Nickname, samurai.Owner, samurai.Username)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SamuraiGetListByOwner(nickname string) ([]models.Samurai, error) {
	stmt := `SELECT Owner, Nickname, TelegramUsername FROM Samurai WHERE Owner = ?`

	rows, err := r.DB.Query(stmt, nickname)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]models.Samurai, 0)

	for rows.Next() {
		s := models.Samurai{}
		err = rows.Scan(&s.Owner, &s.Nickname, &s.Username)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) SamuraiSetOwner(ID any, owner string) error {
	stmt := `UPDATE Samurai SET Owner=? WHERE TelegramUsername=?`
	_, err := r.DB.Exec(stmt, owner, ID)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (r *Repository) SamuraiGetByUsername(username string) (models.Samurai, error) {
	stmt := `SELECT TelegramUsername, Owner, Nickname FROM Samurai WHERE TelegramUsername=?`
	row := r.DB.QueryRow(stmt, username)
	samurai := models.Samurai{}
	err := row.Scan(&samurai.Username, &samurai.Owner, &samurai.Nickname)
	if err != nil {
		return models.Samurai{}, err
	}
	return samurai, nil
}

func (r *Repository) SamuraiSetTurnover(id string, amount float64, date, bank string) error {
	stmt := `INSERT INTO SamuraiTurnover (SamuraiUsername, Amount, Date, Bank) VALUES(?, ?, ?, ?)`
	_, err := r.DB.Exec(stmt, id, amount, date, bank)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//func (m *SamuraiModel) GetList(nickname interface{}) ([]string, error) {
//	stmt := `SELECT Owner, Nickname, TelegramUsername FROM Samurai WHERE Owner = ?`
//
//	rows, err := m.DB.Query(stmt, nickname)
//	if err != nil {
//		log.Print(err)
//		return nil, err
//	}
//	defer rows.Close()
//
//	var result []string
//
//	for rows.Next() {
//		s := &users.Samurai{}
//		err = rows.Scan(&s.Owner, &s.Nickname, &s.TelegramUsername)
//		if err != nil {
//			return nil, err
//		}
//		result = append(result, fmt.Sprintf("%s", s))
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
