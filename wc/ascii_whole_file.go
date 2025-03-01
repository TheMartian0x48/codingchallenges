package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CountResult struct {
	Bytes     int
	Character int
	Line      int
	Word      int
	File      string
}

func printResult(result CountResult) {
	// fmt.Printf("%15d%15d%15d%15d\t%s\n", result.bytes, result.character, result.line, result.word, result.file)
	v, _ := json.MarshalIndent(&result, "", "    ")
	fmt.Println(string(v))
}

func isWhiteSpace(b byte) bool {
	return (b > 8 && b < 14) || b == 32
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("Usage: %s [filename]", os.Args[0]))
	}

	totalResult := CountResult{File: "total"}
	for fileNo := 1; fileNo < len(os.Args); fileNo++ {

		data, err := os.ReadFile(os.Args[fileNo])
		if err != nil {
			panic(err)
		}

		result := CountResult{
			Character: len(data),
			Bytes:     len(data),
			File:      os.Args[fileNo],
		}

		isInWord := false
		word := 0
		line := 0
		for i := 0; i < len(data); i++ {
			if data[i] == 10 {
				line++
			}
			if isWhiteSpace(data[i]) {
				if isInWord {
					word++
				}
				isInWord = false
			} else {
				isInWord = true
			}
		}
		result.Line = line
		result.Word = word

		printResult(result)

		totalResult.Bytes += result.Bytes
		totalResult.Character += result.Character
		totalResult.Line += result.Line
		totalResult.Word += result.Word
	}
	if len(os.Args) > 2 {
		printResult(totalResult)
	}
}
