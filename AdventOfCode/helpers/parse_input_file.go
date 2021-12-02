package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ParseInputFile(filename string) (input []string, err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("could not read from %s: %w", filename, err)
		return
	}

	input = strings.Split(string(data), "\n")
	return
}
