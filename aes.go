package marysue

// const AESKeySize = 128

// func padding(data []byte) []byte {
// 	blockSize := AESKeySize / 8
// 	paddingCount := blockSize - len(data)%blockSize
// 	if paddingCount == 0 {
// 		return data
// 	} else {
// 		return append(data, bytes.Repeat([]byte{0}, paddingCount)...)
// 	}
// }

// func unPadding(data []byte) []byte {
// 	for i := len(data) - 1; ; i-- {
// 		if data[i] != 0 {
// 			return data[:i+1]
// 		}
// 	}
// }

// func Encrypt(data []byte) ([]byte, error) {
// 	blockSize := AESKeySize / 8
// 	block, err := aes.NewCipher(make([]byte, blockSize))
// 	if err != nil {
// 		return nil, err
// 	}
// 	data = padding(data)
// 	encData := make([]byte, len(data))
// 	tmpDAta := make([]byte, blockSize)
// 	for i := 0; i < len(data); i += blockSize {
// 		block.Encrypt(tmpDAta, data[i:i+blockSize])
// 		copy(encData, tmpDAta)
// 	}
// 	return encData, nil

// }

// func Decrypt(data []byte) ([]byte, error) {
// 	blockSize := AESKeySize / 8
// 	block, err := aes.NewCipher(make([]byte, blockSize))
// 	if err != nil {
// 		return nil, err
// 	}
// 	decData := make([]byte, len(data))
// 	tmpData := make([]byte, blockSize)

// 	for i := 0; i < len(data); i += blockSize {
// 		block.Decrypt(tmpData, data[i:i+blockSize])
// 		copy(decData, tmpData)
// 	}
// 	return decData, nil
// }
