package models

import (
	"database/sql"
	"fmt"
)

type Log struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Notes string `json:"notes"`
	Date  string `json:"date"`
}

type LogModel struct {
	DB *sql.DB
}

func (model *LogModel) All() ([]Log, error) {
	var logs []Log

	rows, err := model.DB.Query("SELECT `id`, `name`, `notes`, `date` FROM logs")
	if err != nil {
		return nil, fmt.Errorf("select all logs failed: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var log Log

		if err := rows.Scan(&log.ID, &log.Name, &log.Notes, &log.Date); err != nil {
			return nil, fmt.Errorf("populate log struct failed: %v", err)
		}

		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("select all logs failed: %v", err)
	}

	return logs, nil
}

func (model *LogModel) Add(log Log) (int64, error) {
	result, err := model.DB.Exec("INSERT INTO logs (`name`, `notes`, `date`) VALUES(?, ?, ?);", log.Name, log.Notes, log.Date)

	if err != nil {
		return 0, fmt.Errorf("add log failed: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("add log failed: %v", err)
	}

	return id, nil
}
