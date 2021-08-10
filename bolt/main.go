package main

import (
    "os"
    "fmt"
    "github.com/boltdb/bolt"
    "github.com/go-kit/kit/log"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stdout)
    db, err := bolt.Open("mydb.db", 0600, nil)
    if err != nil {
        logger.Log("open", err)
    }
    //put value into bucket
    if err := db.Update(func(tx *bolt.Tx) error {
        if _, err := tx.CreateBucketIfNotExists([]byte("animal")); err != nil {
            logger.Log("create failed", err.Error())
            return err
        }
        b := tx.Bucket([]byte("animal"))
        err = b.Put([]byte("konglong"), []byte("恐龙"))
        err = b.Put([]byte("monkey"), []byte("猴子"))
        err = b.Put([]byte("cat"), []byte("猫"))

        return err
    }); err != nil {
        logger.Log("update error is:", err.Error())
    }

    //get value from bucket
    if err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("animal"))
        v := b.Get([]byte("konglong"))
        fmt.Printf("value=%s\n", v)
        return nil
    }); err != nil {
        logger.Log("view error :", err.Error())
    }

    //Iterate bucket
    if err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("animal"))
        b.ForEach(func(k, v []byte) error {
            fmt.Printf("key=%s, value=%s\n", k, v)
            return nil
            })
        return nil
    }); err != nil {
        logger.Log("view error :", err.Error())
    }
    defer db.Close()
}
