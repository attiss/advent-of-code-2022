package main

import (
	"bufio"
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
	for _, rucksack := range input {
		for _, commonItem := range rucksack.CommonItems() {
			prioritySum += ItemPriority(commonItem)
		}
	}

	logger.Info("got priority sum", zap.Int("prioritySum", prioritySum))
}

func readInputFromFile(filePath string) ([]Rucksack, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []Rucksack

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, Rucksack{
			compartment1Items: []rune(line[:len(line)/2]),
			compartment2Items: []rune(line[len(line)/2:]),
		})
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return input, err
}
