package utils

import (
	"os"
)

func WriteByteArrayToFile(data []byte, filename string) error {
	str := string(data)

	err := os.WriteFile(filename, []byte(str), 0644)
	if err != nil {
		return err
	}

	return nil
}
