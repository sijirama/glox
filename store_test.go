package main

import (
	"bytes"
	"testing"
)

func TestPathTransformFUnc (t *testing.T)  {
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	store := NewStore(opts)

	data := bytes.NewReader([]byte("Some fucking jpeg bytes"))

	if err := store.writeStream("directory", data); err != nil {
		t.Error(err)
	}

}
