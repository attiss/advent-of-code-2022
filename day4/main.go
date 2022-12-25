package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

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

	fullOverlapCount := 0
	overlapCount := 0
	for _, assignment := range input {
		if assignment.FullOverlap() {
			fullOverlapCount++
		}
		if assignment.Overlap() {
			overlapCount++
		}
	}

	logger.Info("got full overlap count", zap.Int("fullOverlapCount", fullOverlapCount))
	logger.Info("got overlap count", zap.Int("overlapCount", overlapCount))
}

func readInputFromFile(filePath string) ([]Assignment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []Assignment

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), ",")
		if len(lineParts) != 2 {
			return nil, errors.New("invalid input")
		}

		getSection := func(s string) (Section, error) {
			sectionParts := strings.Split(s, "-")
			if len(sectionParts) != 2 {
				return Section{}, errors.New("invalid input")
			}

			start, err := strconv.Atoi(sectionParts[0])
			if err != nil {
				return Section{}, err
			}
			end, err := strconv.Atoi(sectionParts[1])
			if err != nil {
				return Section{}, err
			}

			return Section{Start: start, End: end}, nil
		}

		section1, err := getSection(lineParts[0])
		if err != nil {
			return nil, err
		}
		section2, err := getSection(lineParts[1])
		if err != nil {
			return nil, err
		}

		input = append(input, Assignment{Section1: section1, Section2: section2})
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return input, err
}
