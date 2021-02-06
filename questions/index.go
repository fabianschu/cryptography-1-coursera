package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	iv, _ := hex.DecodeString("20814804c1767293b99f1d9cab3bc3e7")
	padding := len(iv)
	plain := []byte("Pay Bob 100$")
	plainTarget := []byte("Pay Bob 500$")

	for i := 0; i < padding; i++ {
		plain = append(plain, byte(padding))
		plainTarget = append(plainTarget, byte(padding))
	}

	xoredPlains := xor(plain, plainTarget)
	newIv := xor(iv, xoredPlains)
	fmt.Println(hex.EncodeToString(newIv))
}

func xor(input1, input2 []byte) []byte {
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret
}
