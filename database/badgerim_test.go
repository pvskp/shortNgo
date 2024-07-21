package database

import (
	badger "github.com/dgraph-io/badger/v4"
	"testing"
)

func setupBadgerIM(t *testing.T) *badgerIM {
	opts := badger.DefaultOptions("").WithInMemory(true)
	opts.IndexCacheSize = 100 << 20

	db, err := badger.Open(opts)
	if err != nil {
		t.Fatalf("Error opening BadgerIM: %s", err)
	}
	return &badgerIM{
		db: db,
	}
}

func tearDownBadgerIM(b *badgerIM) {
	b.db.Close()
}

func TestSaveHash(t *testing.T) {
	b := setupBadgerIM(t)
	defer tearDownBadgerIM(b)

	hash := "foo"
	value := "bar"
	err := b.SaveHash(hash, value)
	if err != nil {
		t.Errorf("Error saving hash %s on TestSaveHash", err)
	}

	savedValue, err := b.GetHashValue(hash)
	if err != nil {
		t.Errorf("Error getting hash %s", err)
	}

	if savedValue != value {
		t.Errorf("Expected %s, got %s", value, savedValue)
	}
}

func TestGetHashValue(t *testing.T) {
	b := setupBadgerIM(t)
	defer tearDownBadgerIM(b)

	hash := "foo"
	expectedValue := "bar"
	err := b.SaveHash(hash, expectedValue)
	if err != nil {
		t.Errorf("Error saving hash %s on TestGetHashValue", err)
	}

	value, err := b.GetHashValue(hash)
	if err != nil {
		t.Errorf("Error getting hash value: %s", err)
	}

	if value != expectedValue {
		t.Errorf("Expected value %s, got %s", expectedValue, value)
	}

	_, err = b.GetHashValue("nofoo")
	if err == nil {
		t.Errorf("Expected error for nonexistent hash")
	}
}

func TestHashExists(t *testing.T) {
	b := setupBadgerIM(t)
	defer tearDownBadgerIM(b)

	hash := "foo"
	value := "bar"
	err := b.SaveHash(hash, value)
	if err != nil {
		t.Errorf("Error saving hash %s on TestHashExists", err)
	}

	if exists := b.HashExists(hash); !exists {
		t.Errorf("Hash %s should exist", hash)
	}

	nonExistentHash := "nofoo"
	if exists := b.HashExists(nonExistentHash); exists {
		t.Errorf("Hash %s should not exist", hash)
	}
}
