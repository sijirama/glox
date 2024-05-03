package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
)


func CASPathTransformFunc(key string) string {
    hash := sha1.Sum([]byte(key))
    hashedString := hex.EncodeToString(hash[:]) //[:] converts it to a byte slice

    return hashedString
}

type PathTransformFunc func(string) string

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

func DefaultPathTransformFunc(str string) string {
	return str
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (store *Store) writeStream(key string, r io.Reader) error {

	pathName := store.PathTransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	filename := "somefilename"
	pathAndFilename := pathName + "/" + filename

	file, err := os.Create(pathAndFilename)
	if err != nil {
		return err
	}

	n, err := io.Copy(file, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) to disk: %s", n, pathAndFilename)

	return nil
}
