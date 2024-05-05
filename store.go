package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

var defaultRootFolderName = "GloxDefaultFolder"

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashedString := hex.EncodeToString(hash[:]) //[:] converts it to a byte slice

	blockSize := 5
	sliceLength := len(hashedString) / blockSize

	paths := make([]string, sliceLength)

	for i := 0; i < sliceLength; i++ {
		paths[i] = hashedString[i*blockSize : (i+1)*blockSize]
	}

	return PathKey{
		Pathname: strings.Join(paths, "/"),
		Filename: hashedString,
	}

}

type PathKey struct {
	Pathname string
	Filename string
}

func (p PathKey) FirstPathName() string {
	paths := strings.Split(p.Pathname, "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}

func (p PathKey) Fullpath() string {
	return fmt.Sprintf("%s/%s", p.Pathname, p.Filename)
}

type PathTransformFunc func(string) PathKey

//INFO: root is the folder name of the root directory, containing all files and folders of the system
type StoreOpts struct {
	Root              string
	PathTransformFunc PathTransformFunc
}

func DefaultPathTransformFunc(key string) PathKey {
	return PathKey{
		Filename: key,
		Pathname: key,
	}
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if opts.Root == "" {
		opts.Root = defaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (store *Store) Has(key string) bool {
	pathkey := store.PathTransformFunc(key)

	_, err := os.Stat(pathkey.Fullpath())

	if err == fs.ErrNotExist {
		return false
	}

	return true
}

func (store *Store) Delete(key string) error {

	pathkey := store.PathTransformFunc(key)

	firstPathNamewithRoot := fmt.Sprintf("%s/%s", store.Root, pathkey.FirstPathName())

	defer func() {
		log.Printf("deleted [%s] from disk", firstPathNamewithRoot)
	}()

	return os.RemoveAll(firstPathNamewithRoot)
}

func (store *Store) Read(key string) (io.Reader, error) {
	f, err := store.readStream(key)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, f)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (store *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := store.PathTransformFunc(key)
	fullPathwithRoot := fmt.Sprintf("%s/%s", store.Root, pathKey.Fullpath())
	return os.Open(fullPathwithRoot)
}

func (store *Store) writeStream(key string, r io.Reader) error {

	pathKey := store.PathTransformFunc(key)
	pathNamewithRoot := fmt.Sprintf("%s/%s", store.Root, pathKey.Pathname)

	if err := os.MkdirAll(pathNamewithRoot, os.ModePerm); err != nil {
		return err
	}

	fullPathwithRoot := fmt.Sprintf("%s/%s", store.Root, pathKey.Fullpath())

	file, err := os.Create(fullPathwithRoot)
	if err != nil {
		return err
	}

	n, err := io.Copy(file, r)
	if err != nil {
		return err
	}

	log.Printf("written (%d) to disk: %s", n, fullPathwithRoot)

	return nil
}
