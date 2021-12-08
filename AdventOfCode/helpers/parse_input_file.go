package helpers

import (
	"fmt"
	"os"
	"sort"
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

func SortString(unsorted string) (sorted string) {
	s := strings.Split(unsorted, "")
	sort.Strings(s)
	sorted = strings.Join(s, "")
	return
}
