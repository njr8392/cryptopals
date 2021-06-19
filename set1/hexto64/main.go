package main

import (
	"fmt"
	"encoding/base64"
	)



const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
func tobase64(b []byte)[]byte{
	dst:= make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)
	return dst
}

func main(){
	enc := tobase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	fmt.Println(string(enc))
}
