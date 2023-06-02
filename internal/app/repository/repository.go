package repository

import (
	"database/sql"
	"fmt"
)

type InfoPeriod struct {
	Begin float64
	End   float64
}

type Banks struct {
	Sber    float64
	Tinkoff float64
}

type Repository struct {
	DB *sql.DB
}

func New(DB *sql.DB) *Repository {
	return &Repository{DB: DB}
}
func (r *Repository) ValidUser(senderID, role string) bool {
	var temp int
	stmt := `SELECT COUNT(*) FROM ` + role + ` WHERE TelegramUsername=?`
	row := r.DB.QueryRow(stmt, senderID)
	err := row.Scan(&temp)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if temp == 0 {
		return false
	}
	return true
}
