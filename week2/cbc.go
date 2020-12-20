package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	ciphertextHex1 := "4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81"
	ciphertextHex2 := "5b68629feb8606f9a6667670b75b38a5b4832d0f26e1ab7da33249de7d4afc48e713ac646ace36e872ad5fb8a512428a6e21364b0c374df45503473c5242a253"
	ciphertextHex3 := "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329"
	ciphertextHex4 := "770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451"
	cbcDecrypt(ciphertextHex1)
	cbcDecrypt(ciphertextHex2)
	ctrDecrypt(ciphertextHex3)
	ctrDecrypt(ciphertextHex4)
}

func cbcDecrypt(ciphertextHex string) {
	keyHex := "140b41b22a29beb4061bda66b6747e14"
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

func ctrDecrypt(ciphertextHex string) {
	keyHex := "36f18357be4dbd77f050515c73fcf9f2"

	cipherByteArray, _ := hex.DecodeString(ciphertextHex)
	key, _ := hex.DecodeString(keyHex)
	initializationVector := cipherByteArray[0:16]
	ciphertext := cipherByteArray[16:len(cipherByteArray)]

	plainTextByteStore := []byte{}

	for i := 0; i < len(ciphertext); i += 16 {
		decryptedBlock := ciphertext[i : i+16]

		blockCipher, _ := aes.NewCipher(key)
		dst := make([]byte, 16)
		blockCipher.Encrypt(dst, initializationVector)
		plainTextFragment := fixedXorDecrypt(dst, decryptedBlock)

		plainTextByteStore = append(plainTextByteStore, plainTextFragment...)

		fmt.Println(string(plainTextByteStore))
		fmt.Println(plainTextByteStore)
		initializationVector[len(initializationVector)-1]++
	}
}

func fixedXorDecrypt(input1, input2 []byte) []byte {
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret
}
