package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

func BuildHex(bytes []byte) string {
	return strings.ToUpper(hex.EncodeToString(bytes))
}

func ConvertErr(height int64, txHash, errTag string, err error) error {
	return fmt.Errorf("%v-%v-%v-%v", err.Error(), errTag, height, txHash)
}

func MarshalJsonIgnoreErr(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func UnMarshalJsonIgnoreErr(data string, v interface{}) {
	json.Unmarshal([]byte(data), &v)
}
