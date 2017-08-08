package db

type KVStore interface {
	Put(bucketName string, key string, value interface{})
	Get(bucketName string, key string) interface{}
}

type memoryDb struct {
	data map[string]map[string]interface{}
}

func NewDB() KVStore {
	return &memoryDb{
		data: map[string]map[string]interface{}{},
	}
}

func (db *memoryDb) Put(bucketName string, key string, value interface{}) {
	bucket, ok := db.data[bucketName]
	if !ok {
		bucket = make(map[string]interface{})
		db.data[bucketName] = bucket
	}
	bucket[key] = value
}

func (db *memoryDb) Get(bucketName string, key string) interface{} {
	return db.data[bucketName][key]
}
