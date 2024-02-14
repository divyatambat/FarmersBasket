package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/divyatambat/FarmersBasket/internal/pkg/logger"
	"go.uber.org/zap"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Divya@2003"
	dbname   = "farmers_basket_db"
)

func InitializeDatabase() (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		logger.Errorw(context.Background(), "error opening database connection",
			zap.Error(err),
		)
		return nil, err
	}

	//logger.Info("Database connection established successfully!")

	err = db.Ping()
	if err != nil {
		logger.Errorw(context.Background(), "error pinging database",
			zap.Error(err),
		)
		return nil, err
	}

	return db, err
}
