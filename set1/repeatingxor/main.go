package main

import (
	"fmt"
	"encoding/hex"
	)

func main(){
	txt := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	key := "ICE"
	enc := encrypt([]byte(txt), []byte(key))
	fmt.Println(string(enc))
}

//encrypt will xor the message with each successive byte of the key and continue for the length of the message
func encrypt(msg []byte, key []byte)[]byte{
	encoding := []byte{}
	for i :=0; i < len(msg); i++{
		place := i % len(key)
		encoding = append(encoding, msg[i] ^ key[place])
	}
	hexbytes := make([]byte, hex.EncodedLen(len(encoding)))
	hex.Encode(hexbytes, encoding)
	return hexbytes
}
