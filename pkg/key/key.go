package key

import (
	"io/ioutil"
)

type Key struct {
	AppleAuthKey string
}

func NewKey() *Key {
	bytes, _ := ioutil.ReadFile("./pkg/key/AuthKey.p8")
	return &Key{
		AppleAuthKey: string(bytes),
	}
}
