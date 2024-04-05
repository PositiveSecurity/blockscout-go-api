package common

import "errors"

// Custom errors
var (
	ErrBigInt          = errors.New("failed to convert string to big.Int")
	ErrBigFloat        = errors.New("failed to convert string to big.Float")
	ErrNegativeWei     = errors.New("a negative value of wei is not allowed")
	ErrBigIntTooLarge  = errors.New("*big.Int is too large to fit in a uint64")
	ErrNegativeBigInt  = errors.New("cannot convert negative *big.Int to uint64")
	ErrBalanceNotFound = errors.New("balance not found")
)

func CheckError(response *Response) error {
	if response.Message == "OK" || response.Message == "" {
		return nil
	} else if response.Error == "" {
		return nil
	} else {
		return errors.New("error response")
	}
}
