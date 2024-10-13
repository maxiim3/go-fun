package main

import (
	"database/sql/driver"
	"encoding/json"
)

type Person struct {
	Name  string `json:"name"`
	Age   uint8  `json:"age"`
	Email string `json:"email"`
}

func (t *Person) Scan(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Person) Value() (driver.Value, error) {
	b, err := json.Marshal(t)
	return string(b), err
}
