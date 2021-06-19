package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

//Goal: we have a hex encoded string and need to find the hidden message
//Message was xor with some character so we will have to check all characters and undo the xor to find the hidden message
//To score the message we return the one with the largest count of english characters (including whitespace)
//this is probably isn't the most elegant approached but it should suffice to find the hidden message

var max int

func main() {
	cipher := solveXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(string(cipher))
}

func solveXor(s string) []byte {
	text := []byte{}
	buf := []byte(s)
	hexbytes := make([]byte, hex.DecodedLen(len(buf)))
	_, err := hex.Decode(hexbytes, buf)
	if err != nil {
		log.Fatal(err)
	}

	key := []byte{}
	for i := byte(0); i <= 127; i++ {
		for _, char := range hexbytes {
			key = append(key, char^i)
		}
		count := score(string(key))
		if count > max {
			max = count
			text = key
		}
		key = []byte{}
	}
	return text
}

//
func score(s string) int {
	m := make(map[byte]int)
	//count the characters if it is in the alphabet or whitespace
	for _, char := range s {
		if byte(char) > byte('a') && byte(char) < byte('z') || byte(char) > byte('A') && byte(char) < byte('Z') || byte(char) == 32 {
			m[byte(char)]++
		}
	}
	count := total(m)
	return count
}

func total(m map[byte]int) int {
	sum := 0
	for _, val := range m {
		sum += val
	}
	return sum
}
