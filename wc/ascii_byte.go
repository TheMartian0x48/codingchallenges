package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CountResult struct {
	Bytes         int
	Character     int
	Line          int
	MaxByteInLine int
	Word          int
	File          string
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

		result := CountResult{
			File:          os.Args[fileNo],
			MaxByteInLine: 0,
		}
		isInWord := false
		words := 0
		lines := 0
		bytes := 0
		var maxByteInLine int = 0

		file, err := os.Open(os.Args[fileNo])
		if err != nil {
			panic(err)
		}

		data := []byte{1}
		for {
			c, err := file.Read(data)
			if c == 0 {
				break
			}
			if err != nil {
				panic(err)
			}
			b := data[0]
			if b == 10 {
				result.MaxByteInLine = max(result.MaxByteInLine, maxByteInLine)
				lines++
				maxByteInLine = 0
			} else {
				maxByteInLine++
			}
			if isWhiteSpace(b) {
				if isInWord {
					words++
				}
				isInWord = false
			} else {
				isInWord = true
			}
			bytes += 1
		}
		result.Line = lines
		result.Word = words
		result.Bytes = bytes
		result.Character = bytes

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
