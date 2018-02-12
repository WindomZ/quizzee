package tests

import (
	"testing"

	"github.com/WindomZ/quizzee"
	_ "github.com/WindomZ/quizzee-db/bolt"
	"github.com/WindomZ/testify/assert"
)

func init() {
	quizzee.RegisterDB("testing", "../data/data.db")
}

func TestCloseDB(t *testing.T) {
	assert.NoError(t, quizzee.CloseDB())
}
