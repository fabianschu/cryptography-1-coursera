package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	videoFile, _ := ioutil.ReadFile("test-video-file.mp4")

	const blockSize = 1024
	const hashSize = 32

	lastStartingIndex := (len(videoFile) / blockSize) * blockSize
	lastBlockLength := len(videoFile) - lastStartingIndex

	lastBlock := videoFile[lastStartingIndex : lastStartingIndex+lastBlockLength]
	hash := sha256.Sum256(lastBlock)

	for i := lastStartingIndex - blockSize; i >= 0; i -= blockSize {
		currentDataBlock := videoFile[i : i+blockSize]
		hashSlice := hash[:]
		currentCompleteBlock := append(currentDataBlock, hashSlice...)
		hash = sha256.Sum256(currentCompleteBlock)

		if i == 0 {
			result := hex.EncodeToString(hash[:])
			fmt.Println(result)
		}
	}
}
