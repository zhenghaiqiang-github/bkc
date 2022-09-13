package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {

		log.Fatal(err)

	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if nil != err {
			return fmt.Errorf("create bucket: %s", err)
		}
		//写入数据
		if b != nil {
			err := b.Put([]byte("1"), []byte("11"))
			if nil != err {
				return err
			}
		}

		return nil
	})
	//db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("MyBucket"))
	//	v := b.Get([]byte("1"))
	//	fmt.Printf("The answer is: %s\n", v)
	//	return nil
	//})
	//read
	db.View(func(tx *bolt.Tx) error {
		// 获取桶
		b := tx.Bucket([]byte("MyBucket"))
		if nil != b {
			value := b.Get([]byte("1"))
			fmt.Printf("value : %s \n", string(value))
		}
		return nil
	})
}
