package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)
func main(){
	key :=[]byte("YELLOW SUBMARINE")
	file, err := ioutil.ReadFile("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	decode := DecodeBase64(file)
	block, err := aes.NewCipher(key)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ecbDecrypt(block, decode))
}


func DecodeBase64(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(buf, b)
	if err != nil {
		log.Fatal(err)
	}
	return buf[:n]
}

func ecbDecrypt(block cipher.Block, text []byte)[]byte{
	decrypt :=[]byte{}
	for i :=0; i < len(text); i += aes.BlockSize{
		dst := make([]byte, aes.BlockSize)
		block.Decrypt(dst, text[i:i+aes.BlockSize])
		decrypt = append(decrypt, dst...)
	}
	return decrypt
}
