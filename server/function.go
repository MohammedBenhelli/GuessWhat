package server

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetDrawer(l *Lobby) (*User, error) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(l.persons - 1)
	fmt.Println(r)
	in := 0
	for i := range l.Teams {
		for j := range l.Teams[i].Users {
			if in == r {
				return l.Teams[i].Users[j], nil
			}
			in++
		}
	}
	return nil, errors.New("Can't get a random drawer !")
}

//TODO: func parseWord(word string, letter int) string
//parseWord("bonjour", 3) => b__j_u_

func RandomWord() (string, error) {
	index := 0
	rand.Seed(time.Now().UnixNano())
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
		if index == randNumber {
			return scanner.Text(), nil
		}
		scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "", errors.New("Can't get a random word!")
}

func createJSONResp() *JSONResp {
	return &JSONResp{
		Error:   "",
		Message: "",
		Data:    "",
	}
}
