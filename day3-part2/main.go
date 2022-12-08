package main

import (
	"bufio"
	"fmt"
	"os"

	"go.uber.org/zap"
)

const (
	inputFile = "input.txt"
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

	prioritySum := 0
	for _, group := range input {
		prioritySum += ItemPriority(group.GetBadge())
		fmt.Println(string(group.GetBadge()))
	}

	logger.Info("got priority sum", zap.Int("prioritySum", prioritySum))
}

func readInputFromFile(filePath string) ([]Group, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []Group

	scanner := bufio.NewScanner(file)
	getRucksackItems := func() ([]rune, bool) {
		if !scanner.Scan() {
			return nil, true
		}
		return []rune(scanner.Text()), false
	}

	for {
		rucksack1Items, fin := getRucksackItems()
		if fin {
			break
		}
		rucksack2Items, fin := getRucksackItems()
		if fin {
			break
		}
		rucksack3Items, fin := getRucksackItems()
		if fin {
			break
		}

		input = append(input, Group{
			rucksack1Items: []rune(rucksack1Items),
			rucksack2Items: []rune(rucksack2Items),
			rucksack3Items: []rune(rucksack3Items),
		})
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return input, err
}
