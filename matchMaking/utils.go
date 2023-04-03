package matchMaking

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetNormalFloat() float64 {
	r := rand.New(rand.NewSource(777 + time.Now().UnixMilli()))
	return r.NormFloat64()
}

func WriteFileToLocal(data any, fileName string) {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = os.WriteFile(fileName, fileData, 0644)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
