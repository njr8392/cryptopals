package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"io/ioutil"
	"log"
	)

func pcks7(b []byte, size int)[]byte{
	pads := make([]byte, size- len(b))
	
	for i:=0; i<len(pads);i++{
		pads[i] = byte(size-len(b))
	}
	return append(b, pads...)
}

func DecodeBase64(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(buf, b)
	if err != nil {
		log.Fatal(err)
	}
	return buf[:n]
}

func ex9(){
	b := []byte("YELLOW SUBMARINE")
	size := 20
	fmt.Printf("%q\n", pcks7(b, size))
}

func ex10(){
	key := []byte("YELLOW SUBMARINE")
	data, err := ioutil.ReadFile("ex10data.txt")
	if err != nil{
		log.Fatal(err)
	}

	ciphertxt := DecodeBase64(data)
	block, err := aes.NewCipher(key)

	if err != nil{
		log.Fatal(err)
	}
	iv := make([]byte, aes.BlockSize)
	for i:=0; i< aes.BlockSize; i++{
		iv[i] = byte(0)
	}

	txt := CBCDecrypt(block, ciphertxt, iv)
	fmt.Printf("%s\n", txt)

}

func CBCDecrypt(block cipher.Block, ciphertxt []byte, iv []byte)[]byte{
	var decrypted []byte
	prev := iv
	for i:=0; i<len(ciphertxt); i+=aes.BlockSize{
		cur := ciphertxt[i:i+aes.BlockSize]
		dst := make([]byte, aes.BlockSize)
		block.Decrypt(dst, cur)
		decrypted = append(decrypted, xor(dst, prev)...)
		prev = cur
	}
	return decrypted
}

//only works for bytes that are the same length
func xor(a,b []byte)[]byte{
	d := make([]byte, len(a))
	for i:=0; i < len(a); i++{
	d = append(d, a[i]^b[i])
	}
	return d
}

func main(){
	ex9()
	ex10()
}
