package readFile

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var lines []string
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		readString = strings.TrimSpace(readString)
		lines = append(lines, readString)
		if err == io.EOF {
			break
		}
	}
	return lines, err
}
