package main

import (
	"encoding/base64"
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"math"
	"bufio"
)

func main() {
	file, err := ioutil.ReadFile("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	freq := InitScore("../../test/mobydick.txt")
	decode := DecodeBase64(file)
	key := BreakXor(decode, freq)
	fmt.Println(string(key))
	decrypted := make([]byte, len(decode))
	for i,b := range decode{
		decrypted[i] = b ^ key[i%len(key)]
	}
	fmt.Printf("%s\n", decrypted)
}

// read english text and return a map for frequency analysis
func InitScore(f string) map[byte]int{
	m := make(map[byte]int)
	file ,err := os.Open(f)
	if err != nil{
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	for{
		b, _, err := r.ReadLine()
		if err !=nil{
			break
		}
		for _, x := range b{
			m[x]++
		}
	}
	return m
}

//assue equal length [] bytes for now
func HammingDistance(w1, w2 []byte) int {
	diff := 0
	for i, b := range w1 {
		xor := b ^ w2[i] //1 if the bits are different
		for j := xor; j > 0; j >>= 1 {
			if j&0x1 == 1 {
				diff++
			}
		}
	}
	return diff
}

func GuessKeySize(txt []byte, keysize int) int {
	b1 := txt[:keysize]
	b2 := txt[keysize : keysize*2]
	diff := HammingDistance(b1, b2)
	return diff  
}

func Split(b []byte, keysize int) [][]byte {
	blocks := [][]byte{}
i := 0
	for {
		if i > len(b)-keysize {
			break
		}
		blocks = append(blocks, b[i:i+keysize])
		i += keysize
	}
	return blocks
}

//transpose a maxtrix of bytes
func Transpose(b [][]byte) [][]byte {
	buf := make([][]byte, len(b[0]))
	for i := 0; i < len(b[0]); i++ {
		tmp := make([]byte, len(b))
		for j := 0; j < len(b); j++ {
			tmp[j] = b[j][i]
		}
		buf[i] = tmp
	}
	return buf
}

func solveXor(buf []byte, m map[byte]int) ([]byte, byte) {
	var key byte
	var max float64
	text := []byte{}
	for i := 0; i < 256; i++ {
	decode := []byte{}
		for _, char := range buf {
			decode = append(decode, char^byte(i))
		}
		count := score(decode, m)
		if count > max {
			max = count
			key = byte(i)
			text = decode
		}
	}
	return text, key
}

//idea is that the key will produce the most characters in the alphabet
func score(s []byte, m map[byte]int) float64 {
	//count the characters if it is in the alphabet or whitespace
	count :=0
	for _, char := range s {
		count += m[char]
	}
	return float64(count)/ float64(len(s))
}


func DecodeBase64(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(buf, b)
	if err != nil {
		log.Fatal(err)
	}
	return buf[:n]
}

func BreakXor(b []byte, m map[byte]int) []byte {
	var key []byte
	min := math.MaxFloat64
	minKeySize :=0
	for x := 2; x < 40; x++ {

		//take 8 keysize blocks and normalize the distance whichever distance is the lowest is most likely the key
		dist :=0
		for i:=0; i<10; i++{
		a := b[i*x : i*x+x]
		c := b[i*x+x : i*x+x*2]
		dist += HammingDistance(a, c)
		}
		normalized := float64(dist/x)
		//fmt.Printf("size %d hamming distance %.02f\n", x, normalized)
		if normalized < min {
			min = normalized
			minKeySize = x
		}
	}
	//Split file into blocks of the keysize then rearragne so the first byte of every block is in the same block and so on
	block := Split(b, minKeySize)
	trans := Transpose(block)
	
	//CHANGE SCORE TO CHARACTER ANALYSIS OF A BOOK
	for _, line := range trans {
		_, k := solveXor(line, m)
		key = append(key, k)
	}
	return key
}
