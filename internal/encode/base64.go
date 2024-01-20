package encode

import "encoding/base64"

func B64Encode(data *string) *string {
	encodedData := base64.URLEncoding.EncodeToString([]byte(*data))

	return &encodedData
}

func B64Decode(data *string) (*string, error) {
	decodedData, err := base64.URLEncoding.DecodeString(*data)
	if err != nil {
		return nil, err
	}

	decodedDataStr := string(decodedData)

	return &decodedDataStr, nil
}
