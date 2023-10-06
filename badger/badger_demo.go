package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v3"
)


type badgerDb struct {
	path string
	db *badger.DB
}

func (b *badgerDb) Open() error {
	if _, err := os.Stat(b.path); os.IsNotExist(err) {
		os.MkdirAll(b.path, 0755)
	}
	opts := badger.DefaultOptions(b.path)

	// 默认情况下，Badger 确保所有数据都保存在磁盘上。它还支持纯内存模式。
	// 当 Badger 在内存模式下运行时，所有数据都存储在内存中。
	// 在内存模式下读写速度要快得多，但在崩溃或关闭的情况下，存储在 Badger 中的所有数据都会丢失。
	// 要在内存模式下打开 badger，请设置InMemory选项。
	//opts := badger.DefaultOptions(path).WithInMemory(true)

	opts.Dir = b.path
	opts.ValueDir = b.path
	opts.SyncWrites = false
	opts.ValueThreshold = 256
	opts.CompactL0OnClose = true

	db, err := badger.Open(opts)
	if err != nil {
		log.Println("badger open failed", "path", b.path, "err", err)
		return err
	}
	b.db = db
	return nil
}

func (b *badgerDb) Close() {
	err := b.db.Close()
	if err == nil {
		log.Println("database closed", "err", err)
	} else {
		log.Println("failed to close database", "err", err)
	}
}

//Set 要保存键/值对，请使用以下Txn.Set()方法；
//键/值对也可以通过首先创建来保存Entry，然后 Entry使用Txn.SetEntry(). Entry还公开了在其上设置属性的方法。
func (b *badgerDb) Set(key []byte, value []byte) {
	wb := b.db.NewWriteBatch()
	defer wb.Cancel()
	err := wb.SetEntry(badger.NewEntry(key, value).WithMeta(0))
	if err != nil {
		log.Println("Failed to write data to cache.","key", string(key), "value", string(value), "err", err)
	}
	err = wb.Flush()
	if err != nil {
		log.Println("Failed to flush data to cache.","key", string(key), "value", string(value), "err", err)
	}
}

// SetWithTTL Badger 允许在键上设置可选的生存时间 (TTL) 值。一旦 TTL 过去，密钥将不再可检索，并且将有资格进行垃圾收集。
//可以使用和API 方法将 TTL 设置为time.Duration值。Entry.WithTTL() Txn.SetEntry()
func (b *badgerDb) SetWithTTL(key []byte, value []byte, ttl int64) {
	wb := b.db.NewWriteBatch()
	defer wb.Cancel()
	err := wb.SetEntry(badger.NewEntry(key, value).WithMeta(0).WithTTL(time.Duration(ttl * time.Second.Nanoseconds())))
	if err != nil {
		log.Println("Failed to write data to cache.","key", string(key), "value", string(value), "err", err)
	}
	err = wb.Flush()
	if err != nil {
		log.Println("Failed to flush data to cache.","key", string(key), "value", string(value), "err", err)
	}
}

// Get 要读取数据，我们可以使用以下Txn.Get()方法。
func (b *badgerDb) Get(key []byte) string {
	var ival []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		ival, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		log.Println("Failed to read data from the cache.","key", string(key), "error", err)
	}
	return string(ival)
}

func (b *badgerDb) Has(key []byte) (bool, error) {
	var exist bool = false
	err := b.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		} else {
			exist = true
		}
		return err
	})
	// align with leveldb, if the key doesn't exist, leveldb returns nil
	if strings.HasSuffix(err.Error(), "not found") {
		err = nil
	}
	return exist, err
}

// Delete 使用Txn.Delete()方法删除 key。
func (b *badgerDb) Delete(key []byte) error {
	wb := b.db.NewWriteBatch()
	defer wb.Cancel()
	return wb.Delete(key)
}

// IteratorKeysAndValues 要迭代键，我们可以使用Iterator，可以使用 Txn.NewIterator()方法获得。迭代以按字节排序的字典顺序发生。
func (b *badgerDb) IteratorKeysAndValues(){
	err := b.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Failed to iterator keys and values from the cache.","error", err)
	}
}

// IteratorKeys Badger 支持一种独特的迭代模式，称为key-only迭代。
// 它比常规迭代快几个数量级，因为它只涉及对 LSM 树的访问，它通常完全驻留在 RAM 中。
// 要启用仅键迭代，您需要将该IteratorOptions.PrefetchValues 字段设置为false.
// 这也可用于在迭代期间对选定键进行稀疏读取，item.Value()仅在需要时调用。
func (b *badgerDb) IteratorKeys(){
	err := b.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			fmt.Printf("key=%s\n", k)
		}
		return nil
	})

	if err != nil {
		log.Println("Failed to iterator keys from the cache.","error", err)
	}
}

//SeekWithPrefix 要遍历一个键前缀，您可以组合Seek()and ValidForPrefix()
func (b *badgerDb) SeekWithPrefix(prefixStr string){
	err := b.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(prefixStr)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Failed to seek prefix from the cache.", "prefix", prefixStr,"error", err)
	}
}


func main() {
	db := badgerDb{
		path: "/Users/wanglu51/goland_project/src/myself/study_go/data/badger",
	}

	_ = db.Open()
	defer db.Close()


	key1 := []byte("test")
	value1 := []byte("test")
	db.Set(key1, value1)

	key2 := []byte("test2")
	value2 := []byte("test2")
	db.Set(key2, value2)

	key3 := []byte("hello")
	value3 := []byte("hello")
	db.Set(key3, value3)

	println("---- Get ----")
	res := db.Get(key1)
	println("res ---> " + res)

	println("---- IteratorKeysAndValues ----")
	db.IteratorKeysAndValues()
	println("---- IteratorKeys ----")
	db.IteratorKeys()
	println("---- SeekWithPrefix ----")
	db.SeekWithPrefix("test")
	println("---- end ----")
}