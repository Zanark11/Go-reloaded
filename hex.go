
package main

import (
	"strconv"
)

func ConvertirHex(nombre string) (string, error) {

	decimal, err := strconv.ParseInt(nombre, 16, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(decimal, 10), nil
}
