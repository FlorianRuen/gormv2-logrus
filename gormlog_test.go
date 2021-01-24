package gormv2_logrus_test

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	gormv2_logrus "github.com/thomas-tacquet/gormv2-logrus"
)

func TestWithLogrus(t *testing.T) {
	// create a logrus entry to giveit to gormv2_logrus
	logger := &logrus.Entry{}

	// create the gorm compatible logger with logrus instance
	gormLog := gormv2_logrus.New(gormv2_logrus.WithLogrus(logger))

	// open in memory database with previous logger
	db, err := gorm.Open(sqlite.Open(
		"file:unit_test_01?mode=memory&cache=shared"),
		&gorm.Config{Logger: gormLog},
	)

	// check if database correctly created
	require.NoError(t, err)
	require.NotNil(t, db)

	sqlDB, err := db.DB()

	require.NoError(t, err)
	require.NotNil(t, sqlDB)

	defer func() {
		assert.NoError(t, sqlDB.Close())
	}()
}