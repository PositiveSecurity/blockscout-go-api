package blockscout

import (
	"encoding/json"
	"math/big"
	"strconv"
	"time"
)

func hexStringToBigInt(str string) (*big.Int, error) {
	str = str[2:]
	bigInt := new(big.Int)
	bigInt, ok := bigInt.SetString(str, 16)

	if !ok {
		return nil, ErrBigInt
	}

	return bigInt, nil
}

func decStringToBigInt(str string) (*big.Int, error) {
	var bigintVal big.Int

	_, success := bigintVal.SetString(str, 10)

	if !success {
		return nil, ErrBigInt
	}

	return &bigintVal, nil
}

func decStringToBigFloat(str string) (*big.Float, error) {
	var floatValue big.Float
	_, success := floatValue.SetString(str)

	if !success {
		return nil, ErrBigFloat
	}

	return &floatValue, nil
}

func decStringToUint(str string) (uint, error) {
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

func marshalToBytes(result any) ([]byte, error) {
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
