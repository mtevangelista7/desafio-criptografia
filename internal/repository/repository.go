package repository

import (
	"database/sql"
	"desafio-criptografia/pkg/models"
	"log"
)

func CreateDb(db *sql.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		credit_card_token VARCHAR(128) NOT NULL,
		user_document VARCHAR(128) NOT NULL,
		value INTEGER
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela: %s", err)
	}

	return err
}

func CreateTransaction(db *sql.DB, transaction models.Transaction) (int64, error) {
	insertSQL := `INSERT INTO transactions (user_document, credit_card_token, value) VALUES (?, ?, ?)`
	result, err := db.Exec(
		insertSQL,
		transaction.UserDocument,
		transaction.CreditCradToken,
		transaction.Value)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, err
}

func ReadTransaction(db *sql.DB, id int64) (*models.Transaction, error) {
	readSQL := `SELECT * FROM transactions WHERE id = ?`
	result, err := db.Query(readSQL, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var transaction models.Transaction

	if result.Next() {
		err := result.Scan(&transaction.Id, &transaction.UserDocument, &transaction.CreditCradToken, &transaction.Value)
		if err != nil {
			return nil, err
		}
		return &transaction, nil
	}

	return nil, sql.ErrNoRows
}

func DeleteTransaction(db *sql.DB, id int64) error {
	deleteSQL := "DELETE FROM transactions WHERE id = ?"
	_, err := db.Exec(deleteSQL, id)

	return err
}
