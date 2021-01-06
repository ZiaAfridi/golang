package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func readCsvFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file "+path, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path, err)
	}
	return records
}

func getListing(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body_Text := []byte("")
	if response.StatusCode == 200 {
		bodyText, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		body_Text = bodyText
	}
	return body_Text
}

func writCsv(text [][]string) {
	csvfile, err := os.Create("teamprofile.csv")
	if err != nil {
		log.Fatalf("Failed creating file %s", err)
	}
	csvwriter := csv.NewWriter(csvfile)
	err = csvwriter.WriteAll(text) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
	csvwriter.Flush()
	csvfile.Close()

}

func main() {
	records := readCsvFile("./profile.csv")
	//

	for index := range records {

		htmlBody := getListing(records[index][0])
		fmt.Println(string(htmlBody))
		//writCsv(htmlBody)
	}
}
