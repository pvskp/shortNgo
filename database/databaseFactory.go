package database

import "fmt"

func InitiateDB(dbType string) (IDatabase, error) {
	switch dbType {
	case "BadgerIM":
		return newBadgerIM(), nil
	}

	return nil, fmt.Errorf("unknown db type: %s", dbType)
}
