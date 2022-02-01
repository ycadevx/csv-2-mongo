package main

import (
	"fmt"
	"github/csvToMongoDB/source/core"
	"github/csvToMongoDB/source/model"
	"strings"

	"time"
)

func main() {
	defer calculationTime(time.Now(), "Get All Users")
	userArray := make([]model.User, 0)

	dataPerson := core.ReadUserList()

	personChannel := make(chan model.User, 0)

	go PrintPerUser(personChannel, dataPerson)

	for person := range personChannel {
		userArray = append(userArray, person)
		nameUpper := strings.ToUpper(person.Name)
		fmt.Printf("Name : %s Number :%d\n", nameUpper, person.Number)
		time.Sleep((300 * time.Millisecond))
	}
}

//belirli bir nick verilip channel'a gönderildi.
func PrintPerUser(personChannel chan model.User, personList [][]string) {
	for _, person := range personList {
		modelUser := model.User{
			Name:   core.CreateNick(person),
			Number: core.GenerateNumber(1, 512),
		}
		personChannel <- modelUser
	}
}

func calculationTime(start time.Time, description string) {

	elapsed := time.Since((start))

	fmt.Printf("%s Total time :%s", description, elapsed)
}
