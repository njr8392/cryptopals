package main

import (
	"encoding/hex"
	"fmt"
	"log"
	)


func main(){
	buf := xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println(string(buf))
}

// take two equal length buffers and xor them
func xor(s1, s2 string)[]byte{
	b1 := []byte(s1)
	dec := make([]byte, hex.DecodedLen(len(b1)))
	_, err := hex.Decode(dec, b1)
	if err != nil{
		log.Fatal(err)
	}
	b2 := []byte(s2)
	dec2 := make([]byte, hex.DecodedLen(len(b2)))
	_, err = hex.Decode(dec2, b2)
	if err != nil{
		log.Fatal(err)
	}
	buf := []byte{}
	for i, _ := range dec{
		buf = append(buf, dec[i] ^ dec2[i]) 
	}
	enc := make([]byte, hex.EncodedLen(len(buf)))
	hex.Encode(enc, buf)
	return enc
}
