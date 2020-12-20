package shared

func fixedXorDecrypt(input1, input2 []byte) []byte {
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret
}
