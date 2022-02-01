package core

import (
	"encoding/csv"
	"github/csvToMongoDB/source/shared"
	"log"
	"math/rand"
	"os"
	"time"
)

// Min-max değerler arasında rastgele bir numara üret
func GenerateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min) + min
	return num
}

//user isim / soyisim ayraç
func CreateNick(person []string) string {
	return person[0][0:1] + person[1][0:2]
}

func ReadUserList() [][]string {

	fileCsv := shared.Config.FILEURL
	personData := make([][]string, 0)

	f, err := os.Open(fileCsv)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = shared.Config.SPERATOR

	dataPerson, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range dataPerson {
		if (shared.Config.ISHEADER && i > 0) || !shared.Config.ISHEADER {
			personData = append(personData, []string{line[1], line[2]})
		}
	}

	return personData

}
