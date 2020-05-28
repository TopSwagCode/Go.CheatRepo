package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {

	log.Println("Log messages is per default 'Stderr'")
	log.SetOutput(os.Stdout)
	log.Println("But we can change it by changing the output to be something else. Like now we are using 'log.SetOutput(os.Stdout)'")
	log.Println("We could pick file logging or something else if we would like")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("now 'log.SetFlags(log.LstdFlags | log.Lshortfile)' including Time and filename / line number for all logs in future.")

	log.SetOutput(os.Stderr)

	log.Println("Setting log output back to error 'log.SetOutput(os.Stderr)'")

	basicErrorHandling()

	fmt.Println()
	fmt.Println("- - - - - - - - - - - - - - - - - - - ")
	fmt.Println()

	basicErrorHandlingUsingPackageHandling()

	fmt.Println()
	fmt.Println("- - - - - - - - - - - - - - - - - - - ")
	fmt.Println()

	errorWrapping()


}

func basicErrorHandling(){
	var myStoredPerson person
	fmt.Println("Basic Go")
	fmt.Println()
	fmt.Println("Trying to Unmarshal invalid json")
	myJsonString := `This is not json!`

	// parsed data in
	err := json.Unmarshal([]byte(myJsonString), &myStoredPerson)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(myStoredPerson)
	}

	fmt.Println()
	fmt.Println("Trying to Unmarshal valid json with wrong fields")
	myJsonString = `{"Navn":"Joshua Ryder", "Alder": 33, "KanLideGo":true}`

	err = json.Unmarshal([]byte(myJsonString), &myStoredPerson)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(myStoredPerson)
	}

	fmt.Println()
	fmt.Println("Trying to Unmarshal valid json with right fields")
	myJsonString = `{"Name":"Joshua Ryder", "Age": 33, "DoesLikeGo":true}`

	err = json.Unmarshal([]byte(myJsonString), &myStoredPerson)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(myStoredPerson)
	}
}

func errorWrapping () {
	dataStuff, err := DoDataStuff()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dataStuff)
}

func DoDataStuff () (string, error) {
	//...
	data, err := fetchDataFromApi()

	if err != nil {
		errorMsg := "Can't handle invalid or missing data"
		wrappedError := errors.Wrap(err, errorMsg)
		return "", wrappedError
	}

	return data, nil
}

func fetchDataFromApi () (string, error) {
	//...

	errorMsg := "Could not connect to server"
	return "ADASD", errors.New(errorMsg)
}

type person struct {
	Name string
	Age int
	DoesLikeGo bool
}

