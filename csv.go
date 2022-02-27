package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func CsvMain() {

	csvFile, err := os.OpenFile("output.csv", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	files, err := ioutil.ReadDir("submissions")
	if err != nil {
		log.Fatal(err)
	}

	header := []string{
		"id",
		"session_id",
		"category",
		"mood",
		"build_id",
		"timestamp",
		"level_name",
		"level_pos",
		"playtime",
		"level_seed",
		"text",
	}

	w.Write(header)
	for _, file := range files {
		jsonFile, err := os.Open("submissions/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var sub FeedbackSubmission
		json.Unmarshal(byteValue, &sub)
		fmt.Println(string(byteValue))
		record := []string{
			sub.ID,
			sub.SessionID,
			sub.Category,
			strconv.Itoa(int(sub.Mood)),
			sub.BuildID,
			sub.Timestamp,
			sub.LevelName,
			sub.LevelPos,
			strconv.Itoa(sub.Playtime),
			sub.LevelSeed,
			sub.Text,
		}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

	}
}
