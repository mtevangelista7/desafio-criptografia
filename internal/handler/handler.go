package handler

import (
	"crypto/sha512"
	"database/sql"
	"desafio-criptografia/internal/repository"
	"desafio-criptografia/pkg/models"
	"encoding/hex"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

const fileName = "sqlite.db"

// @Summary Make a transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Transaction data"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]string
// @Router /makeTransaction [post]
func MakeTransaction(c *gin.Context) {
	var newTransaction models.Transaction

	if err := c.BindJSON(&newTransaction); err != nil {
		c.JSON(400, map[string]string{"error": "Invalid input"})
		return
	}

	sha_512 := sha512.New()
	sha_512.Write([]byte(newTransaction.CreditCradToken))
	newTransaction.CreditCradToken = hex.EncodeToString(sha_512.Sum(nil))

	sha_512 = sha512.New()
	sha_512.Write([]byte(newTransaction.UserDocument))
	newTransaction.UserDocument = hex.EncodeToString(sha_512.Sum(nil))

	db, err := sql.Open("sqlite", fileName)

	defer db.Close()

	if err != nil {
		c.JSON(500, map[string]string{"error": "Database connection error"})
		return
	}

	newTransaction.Id, err = repository.CreateTransaction(db, newTransaction)
	if err != nil {
		c.JSON(500, map[string]string{"error": "Transaction creation failed"})
		return
	}

	c.JSON(200, newTransaction)
}

// @Summary Get a transaction by user ID
// @Description Retrieve a transaction using the user ID from the URL
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 200 {object} models.Transaction
// @Failure 500 {object} map[string]string
// @Router /getTransaction/{id} [get]
func GetTransaction(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(500, map[string]string{"error": "Invalid user ID"})
		return
	}

	db, err := sql.Open("sqlite", fileName)
	defer db.Close()
	if err != nil {
		c.JSON(500, map[string]string{"error": "Database connection error"})
		return
	}

	newTransaction, err := repository.ReadTransaction(db, userID)
	if err != nil {
		c.JSON(500, map[string]string{"error": "transaction"})
		return
	}

	c.JSON(200, newTransaction)
}

// @Summary Delete a transaction by user ID
// @Description Delete a transaction using the user ID from the URL
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int64 true "User ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 500 {object} map[string]string "Database error"
// @Router /deleteTransaction/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(500, map[string]string{"error": "Invalid user ID"})
		return
	}

	db, err := sql.Open("sqlite", fileName)
	defer db.Close()
	if err != nil {
		c.JSON(500, map[string]string{"error": "Database connection error"})
		return
	}

	err = repository.DeleteTransaction(db, userID)
	if err != nil {
		c.JSON(500, map[string]string{"error": "..."})
		return
	}

	c.JSON(204, "...")
}
