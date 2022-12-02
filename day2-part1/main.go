package main

import (
	"bufio"
	"errors"
	"os"
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

	var myScore int
	for _, battle := range input {
		myScore += battle.Score()
	}

	logger.Info("got score", zap.Int("score", myScore))
}

func readInputFromFile(filePath string) ([]Battle, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []Battle

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		battleSpec := strings.Split(scanner.Text(), " ")
		if len(battleSpec) != 2 {
			return nil, errors.New("invalid battle spec")
		}

		var b Battle

		switch battleSpec[0] {
		case "A":
			b.OpponentShape = rock
		case "B":
			b.OpponentShape = paper
		case "C":
			b.OpponentShape = scissor
		default:
			return nil, errors.New("invalid value for opponent shape")
		}

		switch battleSpec[1] {
		case "X":
			b.MyShape = rock
		case "Y":
			b.MyShape = paper
		case "Z":
			b.MyShape = scissor
		default:
			return nil, errors.New("invalid value for my shape")
		}

		input = append(input, b)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return input, err
}
