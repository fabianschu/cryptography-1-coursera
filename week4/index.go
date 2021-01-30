package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
)

func main() {
	hexString := "f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4"
	cipherText, _ := hex.DecodeString(hexString)
	blockLength := 16
	amountBlocks := len(cipherText) / blockLength
	allBlocks := createBlocks(cipherText, amountBlocks)
	var guessBytes []byte
	lowerCaseBytes := makeRange(97, 122)
	upperCaseBytes := makeRange(65, 90)
	paddingNumbers := makeRange(1, 16)
	var resultContainer []byte

	guessBytes = append(guessBytes, paddingNumbers...)
	guessBytes = append(guessBytes, lowerCaseBytes...)
	guessBytes = append(guessBytes, upperCaseBytes...)
	guessBytes = append(guessBytes, byte(32))

	for blockIdx := 1; blockIdx < amountBlocks; blockIdx++ {
		fmt.Println("BLOCK: ", blockIdx)

		prevBlock := allBlocks[blockIdx-1]
		prevBlockCopy := make([]byte, blockLength)
		tempContainer := make([]byte, blockLength)
		copy(prevBlockCopy, prevBlock)
		fmt.Println(prevBlock)
		fmt.Println(prevBlockCopy)

		for i := 15; i >= 0; i-- {

			padding := blockLength - i
			fmt.Println("PADDING: ", padding)
			for j := 0; j < len(guessBytes); j++ {
				g := guessBytes[j]
				prevBlockCopy[i] = prevBlock[i] ^ g ^ byte(padding)

				var sendBytes []byte
				sendBlocks := allBlocks[0 : blockIdx+1]
				sendBlocks[blockIdx-1] = prevBlockCopy

				for _, block := range sendBlocks {
					sendBytes = append(sendBytes, block...)
				}

				resp, _ := http.Get("http://crypto-class.appspot.com/po?er=" + hex.EncodeToString(sendBytes))
				if resp.StatusCode == 404 {
					tempContainer[i] = g
					for k := 0; k < padding; k++ {
						idx := blockLength - 1 - k
						prevBlockCopy[idx] = prevBlock[idx] ^ tempContainer[idx] ^ byte(padding+1)
					}
					break
				}

				if j == len(guessBytes)-1 {
					fmt.Println("NO VALID PADDING FOUND")
					return
				}
			}
			fmt.Println(tempContainer)
		}
		for _, char := range tempContainer {
			resultContainer = append(resultContainer, char)
		}
		decrypted := string(resultContainer)
		fmt.Println(decrypted)
	}

}

func xor(input1, input2 []byte) []byte {
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret
}

func createBlocks(array []byte, amountBlocks int) [][]byte {
	container := make([][]byte, amountBlocks)
	for i := 0; i < amountBlocks; i++ {
		container[i] = array[i*16 : i*16+16]
	}
	return container
}

func makeRange(min, max int) []byte {
	a := make([]byte, max-min+1)
	for i := 0; i < max-min+1; i++ {
		a[i] = byte(max - i)
	}
	return a
}
