package main

import (
	"fmt"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

type Item struct {
	Key   string `json:"key" xml:"key"`
	Value string `json:"value" xml:"value"`
}

// The one and only database instance.
var db, err = leveldb.OpenFile("c:\\leveldb", nil)

func ConvertKeyVal(key, value string) *Item {
	return &Item{
		Key:   key,
		Value: value,
	}
}

// GetKey returns the requested key/value.
func GetKey(enc Encoder, r *http.Request) (int, string) {
	key := r.FormValue("key")
	value, err := db.Get([]byte(key), nil)
	if err != nil || value == nil {
		// Invalid key, or does not exist
		return http.StatusNotFound, Must(enc.Encode(
			NewError(ErrCodeNotExist, fmt.Sprintf("the key with id %s does not exist", r.FormValue("key")))))
	}

	al := ConvertKeyVal(string(key), string(value))
	return http.StatusOK, Must(enc.Encode(al))
}

// AddKey creates the posted key value.
func AddKey(w http.ResponseWriter, r *http.Request, enc Encoder) (int, string) {
	key, value := r.FormValue("key"), r.FormValue("value")

	if key == "" || value == "" {
		fmt.Println("Missing key/value")
		return http.StatusBadRequest, ""
	}

	// Add the file to the database
	err := db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return http.StatusNotModified, ""
	}

	al := ConvertKeyVal(string(key), string(value))
	return http.StatusCreated, Must(enc.Encode(al))
}

// DeleteKey deletes the corresponding key/value
func DeleteKey(enc Encoder, r *http.Request) (int, string) {
	key := r.FormValue("key")

	if key == "" {
		fmt.Println("Missing key")
		return http.StatusBadRequest, ""
	}

	err = db.Delete([]byte(key), nil)
	if err != nil {
		return http.StatusNotModified, ""
	} else {
		fmt.Printf("Key [%s] deleted\n", key)
	}

	return http.StatusNoContent, ""
}
