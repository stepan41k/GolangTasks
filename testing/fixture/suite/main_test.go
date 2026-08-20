package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

func Calculate(db *sql.DB, amount float64) (float64, error) {
	var percent float64
	err := db.QueryRow("SELECT percent FROM discounts WHERE active = 1 LIMIT 1").Scan(&percent)
	if err != nil {
		if err == sql.ErrNoRows {

			return amount, nil
		}
		return 0, err
	}

	return amount * (1 - percent/100), nil
}

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	query := `
		CREATE TABLE discounts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			percent REAL,
			active INTEGER
		);
		INSERT INTO discounts (percent, active) VALUES (10.0, 1);
		`
	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	return db
}

type AdServiceTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (s *AdServiceTestSuite) SetupTest() {
	s.db = connectToDB()
}

func (s *AdServiceTestSuite) TearDownTest() {
	s.db.Close()
}

func (s *AdServiceTestSuite) TestCalculateDiscount_Success() {
	res, err := Calculate(s.db, 100)

	s.NoError(err)
	s.Equal(90.0, res)
}

func (s *AdServiceTestSuite) TestCalculateDiscount_ZeroAmount() {
	res, err := Calculate(s.db, 0)

	s.NoError(err)
	s.Equal(0.0, res)
}

func (s *AdServiceTestSuite) TestCalculateDiscount_NoActiveDiscount() {
	_, err := s.db.Exec("DELETE FROM discounts")
	s.Require().NoError(err)

	res, err := Calculate(s.db, 100)

	s.NoError(err)
	s.Equal(100.0, res)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(AdServiceTestSuite))
}
