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
	BucketExists(bucketName string) bool
	Put(bucketName string, key string, value []byte) error
	List(bucketName string) ([]KeyValue, error)
	Get(bucketName string, key string) ([]byte, error)
	Delete(bucketName string, key string) error
}

type kvStore struct {
	boltdb *bolt.DB
}

type KeyValue struct {
	Key   []byte
	Value []byte
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

func (db *kvStore) BucketExists(bucketName string) bool {
	var exists bool
	db.boltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		exists = b != nil
		return nil
	})
	return exists
}

func (db *kvStore) Put(bucketName string, key string, value []byte) error {
	return db.boltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("bucket does not exist")
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
			return errors.New("bucket does not exist")
		}
		value = b.Get([]byte(key))
		return nil
	})
	return value, err
}

func (db *kvStore) List(bucketName string) ([]KeyValue, error) {
	var list []KeyValue
	err := db.boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("bucket does not exist")
		}
		c := b.Cursor()
		list = []KeyValue{}
		for key, value := c.First(); key != nil; key, value = c.Next() {
			list = append(list, KeyValue{
				Key:   key,
				Value: value,
			})
		}
		return nil
	})
	return list, err
}

func (db *kvStore) Delete(bucketName string, key string) error {
	err := db.boltdb.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("bucket does not exist")
		}
		return b.Delete([]byte(key))
	})
	return err
}

func (db *kvStore) Close() error {
	return db.boltdb.Close()
}
