package database

import (
	"testing"
)

func TestConnetDB(t *testing.T) {
	db, err := ConnectDB()

	if err != nil {
		t.Errorf("expected not error, but got %v", err.Error())
	}

	if db == nil {
		t.Error("expected not nil, but got nil")
	}
}
