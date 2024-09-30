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

func CreateTransaction(db *sql.DB, transaction models.Transaction) error {
	insertSQL := `INSERT INTO transactions (userDocument, creditCardToken, value) VALUES (?, ?, ?)`
	_, err := db.Exec(
		insertSQL,
		transaction.UserDocument,
		transaction.CreditCradToken,
		transaction.Value)

	return err
}

func ReadTransaction(db *sql.DB, id int64) (*models.Transaction, error) {
	readSQL := `SELECT * FROM transactions WHERE id = ?`
	result, err := db.Query(readSQL, id)
	if err != nil {
		return nil, err
	}

	var transaction models.Transaction

	err = result.Scan(&transaction.Id, &transaction.UserDocument, &transaction.CreditCradToken, &transaction.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &transaction, nil
}

func DeleteTransaction(db *sql.DB, id int64) error {
	deleteSQL := "DELETE FROM transactions WHERE id = ?"
	_, err := db.Exec(deleteSQL, id)

	return err
}
