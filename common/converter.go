package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func HexStringToBigInt(str string) (*big.Int, error) {
	str = str[2:]
	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(str, 16)

	if !ok {
		return nil, ErrBigInt
	}

	return bigInt, nil
}

func DecStringToBigInt(str string) (*big.Int, error) {
	var bigintVal big.Int

	_, success := bigintVal.SetString(str, 10)

	if !success {
		return nil, ErrBigInt
	}

	return &bigintVal, nil
}

func DecStringToBigFloat(str string) (*big.Float, error) {
	var floatValue big.Float
	_, success := floatValue.SetString(str)

	if !success {
		return nil, ErrBigFloat
	}

	return &floatValue, nil
}

func DecStringToUint(str string) (uint, error) {
	intValue, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(intValue), nil
}

func BigIntToUint64(bigInt *big.Int) (uint64, error) {
	if bigInt.Sign() == -1 {
		return 0, ErrNegativeBigInt
	}

	if bigInt.BitLen() > 64 {
		return 0, ErrBigIntTooLarge
	}

	return bigInt.Uint64(), nil
}

func MarshalToBytes(result any) ([]byte, error) {
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return resultBytes, nil
}

func UnixTimestampToNormal(unixTimestamp int64) string {
	t := time.Unix(unixTimestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// Convert WEI to ETH
func WeiToEther(wei *big.Int) (*big.Float, error) {

	if wei.Sign() == -1 {
		return nil, ErrNegativeWei
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, new(big.Float).SetInt(divisor))

	return ether, nil
}

func SaveToFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func PrettyPrintJSON(data []byte) error {
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		errMsg := fmt.Sprintf("Error decoding JSON: %v\n", err)
		return errors.New(errMsg)
	}
	prettyData, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		errMsg := fmt.Sprintf("Error formatting JSON: %v\n", err)
		return errors.New(errMsg)
	}
	fmt.Println(string(prettyData))
	return nil
}
