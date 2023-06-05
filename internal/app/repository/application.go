package repository

import (
	"emivn/internal/models"
	"log"
)

func (r *Repository) ApplicationInsert(creater, cardID string, sum int) error {
	stmt := `INSERT INTO Application (Daimyo, ID, Sum) VALUES(?, ?, ?)`
	_, err := r.DB.Exec(stmt, creater, cardID, sum)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) ApplicationGetAll() ([]models.Application, error) {
	stmt := "SELECT * FROM Application"
	applications := make([]models.Application, 0)
	rows, err := r.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := models.Application{}
		err = rows.Scan(&temp.Creater, &temp.Sum, &temp.CardID, &temp.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		applications = append(applications, temp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return applications, nil
}
func (r *Repository) ApplicationDelete(cardID string) error {
	stmt := `DELETE FROM Application WHERE ID=?`
	_, err := r.DB.Exec(stmt, cardID)
	if err != nil {
		return err
	}
	return nil
}
