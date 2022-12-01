package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

const (
	inputFile = "input.txt"
	separator = -1
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	input, err := readInputFromFile(inputFile)
	if err != nil {
		logger.Error("failed to read input", zap.Error(err))
		panic(err)
	}

	var sums []int

	var tmpSum int
	for _, inputItem := range input {
		if inputItem == separator {
			sums = append(sums, tmpSum)
			tmpSum = 0
			continue
		}
		tmpSum += inputItem
	}

	sort.Ints(sums)

	logger.Info("found max of sums", zap.Int("max", sums[len(sums)-1]))
	logger.Info("sum of top three", zap.Int("sum", sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3]))
}

func readInputFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			input = append(input, separator)
			continue
		}

		parsedInt, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		input = append(input, parsedInt)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return input, err
}
