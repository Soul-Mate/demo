package models

import "github.com/Soul-Mate/demo/soa_demo/db"

type MemberModel struct {
	Id     int     `json:"member_id"`
	Name   string  `json:"member_name"`
	Level  int     `json:"member_level"`
	Amount float64 `json:"member_amount"`
}

func (m *MemberModel) Find(id int) (*MemberModel, error) {
	conn, err := db.DefaultDB.Connection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	row := conn.QueryRow("SELECT * from members where id = ?", id)
	if err = row.Scan(&m.Id, &m.Name, &m.Level, &m.Amount); err != nil {
		return nil, err
	}
	return m, nil
}
