package server

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"os"
)

func RandomWord() (string, error) {
	index := 0
	randNumber := rand.Intn(WORDS)

	file, err := os.Open(FILENAME); if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		index++
		fmt.Println(scanner.Text())
		if index == randNumber {
			return scanner.Text(), nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "", errors.New("Can't get a random word!")
}
