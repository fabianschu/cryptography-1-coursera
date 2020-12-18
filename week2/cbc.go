package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	keyHex := "140b41b22a29beb4061bda66b6747e14"
	ciphertextHex := "4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81"

	cipherByteArray, _ := hex.DecodeString(ciphertextHex)
	key, _ := hex.DecodeString(keyHex)
	initializationVector := cipherByteArray[0:16]

	ciphertext := cipherByteArray[16:len(cipherByteArray)]

	plainTextByteStore := []byte{}

	for i := 0; i < len(ciphertext); i += 16 {
		decryptedBlock := ciphertext[i : i+16]

		blockCipher, _ := aes.NewCipher(key)
		dst := make([]byte, 16)
		blockCipher.Decrypt(dst, decryptedBlock)

		var plainTextBytes []byte

		if i != 0 {

			plainTextBytes = fixedXorDecrypt(ciphertext[i-16:i], dst)
		} else {
			plainTextBytes = fixedXorDecrypt(initializationVector, dst)
		}

		plainTextByteStore = append(plainTextByteStore, plainTextBytes...)
	}

	padding := plainTextByteStore[len(plainTextByteStore)-1]
	result := plainTextByteStore[0 : len(plainTextByteStore)-int(padding)]
	fmt.Println(string(result))
}

func fixedXorDecrypt(input1, input2 []byte) []byte {
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret
}
