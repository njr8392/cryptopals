package main

import (
	"fmt"
	)

func pcks7(b []byte, size int)[]byte{
	pads := make([]byte, size- len(b))
	
	for i:=0; i<len(pads);i++{
		pads[i] = byte(size-len(b))
	}
	return append(b, pads...)
}

func ex9(){
	b := []byte("YELLOW SUBMARINE")
	size := 20
	fmt.Printf("%q\n", pcks7(b, size))
}

func main(){
	ex9()
}
