package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func ctrDecrypt(ciphertextHex string) {
	keyHex := "36f18357be4dbd77f050515c73fcf9f2"
	// ciphertextHex3 := "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329"
	// ciphertext2 := "770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451"

	cipherByteArray, _ := hex.DecodeString(ciphertextHex)
	key, _ := hex.DecodeString(keyHex)
	initializationVector := cipherByteArray[0:16]
	ciphertext := cipherByteArray[16:len(cipherByteArray)]

	// plainTextByteStore := []byte{}

	for i := 0; i < len(ciphertext); i += 16 {
		decryptedBlock := ciphertext[i : i+16]

		blockCipher, _ := aes.NewCipher(key)
		dst := make([]byte, 16)
		blockCipher.Encrypt(dst, initializationVector)

		// var plainTextBytes []byte

		blu := fixedXorDecrypt(dst, decryptedBlock)

		fmt.Println(blu)

		// plainTextByteStore = append(plainTextByteStore, plainTextBytes...)
	}
}
