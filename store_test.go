package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFUnc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)

	data := bytes.NewReader([]byte("Some fucking jpeg bytes"))

	if err := store.writeStream("directory", data); err != nil {
		t.Error(err)
	}

}
