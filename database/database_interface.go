package database

type IDatabase interface {
	HashExists(hash string) (exists bool)
	SaveHash(hash, value string) (err error)
	GetHashValue(hash string) (value string, err error)
}
