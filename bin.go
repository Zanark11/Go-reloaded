
package main

import (
	"strconv"
)

func ConvertirBin(nombre string) (string, error) {

	decimal, err := strconv.ParseInt(nombre, 2, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(decimal, 10), nil
}
