package main 

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha512"
)

var key = []byte("dannon yes")

func main() {

}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error message") 
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) {

}