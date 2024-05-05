package main

import (
	"bytes"
	"fmt"
	"io"

	//"io"
	"testing"
)

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	store := newStore()
	defer tearDown(t, store)

	for i := 0; i < 10; i++ {

		key := fmt.Sprintf("momsbestpicture__%d", i)

		data := ([]byte("Some fucking jpeg bytes"))

		if err := store.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := store.Has(key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		r, err := store.Read(key)
		if err != nil {
			t.Error(err)
		}
		b, _ := io.ReadAll(r)
		fmt.Println(string(b))
		if string(b) != string(data) {
			t.Errorf("have %v, want %v", string(b), string(data))
		}

		if err := store.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := store.Has(key); ok {
			t.Errorf("expected to not have key %s", key)
		}

	}
}
