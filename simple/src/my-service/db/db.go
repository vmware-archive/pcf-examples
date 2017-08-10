package db

import (
	"errors"
	"io"

	"github.com/boltdb/bolt"
)

//go:generate counterfeiter ./ KVStore

type KVStore interface {
	io.Closer

	CreateBucket(bucketName string) error
	DeleteBucket(bucketName string) error
	Put(bucketName string, key string, value []byte) error
	Get(bucketName string, key string) ([]byte, error)
}

type kvStore struct {
	boltdb *bolt.DB
}

func NewDB(filename string) (KVStore, error) {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &kvStore{
		boltdb: db,
	}, nil
}

func (db *kvStore) CreateBucket(bucketName string) error {
	return db.boltdb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

func (db *kvStore) DeleteBucket(bucketName string) error {
	return db.boltdb.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		return err
	})
}

func (db *kvStore) Put(bucketName string, key string, value []byte) error {
	return db.boltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("Bucket does not exist")
		}
		err := b.Put([]byte(key), value)
		return err
	})
}

func (db *kvStore) Get(bucketName string, key string) ([]byte, error) {
	var value []byte
	err := db.boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("Bucket does not exist")
		}
		value = b.Get([]byte(key))
		return nil
	})
	return value, err
}

func (db *kvStore) Close() error {
	return db.boltdb.Close()
}
