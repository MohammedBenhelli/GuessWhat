package server

import (
	"encoding/json"
	"log"
)

func (j *JSONResp) toJSON() []byte {
	b, err := json.Marshal(j)
	if err != nil {
		log.Fatal("Can't marshal JSON response !")
	}
	return b
}
