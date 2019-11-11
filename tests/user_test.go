package tests

import (
	"encoding/json"
	"testing"
)

type name struct {
	Id       int
	Username string
	Password string
}

func newName() *name {
	return &name{}
}

func TestUser(t *testing.T) {
	var n name
	str := `{"id":1,"username":"hujie","password":"hujie123"}`
	json.Unmarshal([]byte(str), &n)


}
