package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetJSONValue() (string, error) {
	jsonValues, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonValues), nil
}
