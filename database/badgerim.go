package database

import (
	"errors"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

type badgerIM struct {
	db *badger.DB
}

func (b *badgerIM) HashExists(hash string) bool {
	err := b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(hash))
		if err != nil {
			log.Printf("Error getting hash: %s", err)
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			log.Printf("Hash not found: %s", err)
		} else {
			log.Printf("Unexpected error: %s", err)
		}
		return false
	}

	return true
}

func (b *badgerIM) GetHashValue(hash string) (value string, err error) {
	var valueCopy []byte

	err = b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(hash))
		if err != nil {
			log.Printf("Error getting hash: %s", err)
			return err
		}

		err = item.Value(func(val []byte) error {
			valueCopy = append([]byte{}, val...)
			return nil
		})

		if err != nil {
			log.Printf("Error getting value: %s", err)
			return err
		}

		return nil
	})

	value = string(valueCopy)

	return
}

func (b *badgerIM) SaveHash(hash, value string) (err error) {
	err = b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(hash), []byte(value))
		if err != nil {
			log.Printf("Error saving hash: %s", err)
			return err
		}
		return nil
	})

	return
}

func newBadgerIM() *badgerIM {
	opts := badger.DefaultOptions("").WithInMemory(true)
	opts.IndexCacheSize = 100 << 20

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Error opening BadgerIM: %s", err)
	}
	return &badgerIM{
		db: db,
	}
}
