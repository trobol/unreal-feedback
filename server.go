package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type FeedbackSubmission struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	SessionID string `json:"session_id"`
	Category  string `json:"category"`
	Mood      uint8  `json:"mood"`
	BuildID   string `json:"build_id"`
	Text      string `json:"text"`
	Playtime  int    `json:"playtime"`
	LevelName string `json:"level_name"`
	LevelPos  string `json:"level_pos"`
	LevelSeed string `json:"level_seed"`
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func httpFeedback(w http.ResponseWriter, req *http.Request) {
	err, code := feedback(w, req)
	if err != nil {
		http.Error(w, err.Error(), code)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("success"))
}

func feedback(w http.ResponseWriter, req *http.Request) (error, int) {
	if req.Method != "POST" {
		return errors.New("only post allowed"), http.StatusBadRequest
	}

	var submission FeedbackSubmission
	err := json.NewDecoder(req.Body).Decode(&submission)
	if err != nil {
		return err, http.StatusBadRequest
	}

	err = validateSubmission(submission)
	if err != nil {
		return err, http.StatusBadRequest
	}

	err = saveSubmission(submission)
	if err != nil {
		return err, http.StatusBadRequest
	}

	return nil, 200
}

func validateSubmission(submission FeedbackSubmission) error {
	var valid_categories = []string{"general", "bug", "gameplay", "performance"}

	if !contains(valid_categories, submission.Category) {
		return errors.New("invalid category")
	}

	// TODO: validate timestamp
	// TODO: validate PlayerID
	// TODO: validate BuildID
	// TODO: validate mood

	return nil
}

func saveSubmission(submission FeedbackSubmission) error {

	path := fmt.Sprintf("submissions/%s.json", submission.ID)

	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsExist(err) {
			// TODO: a feedback with the same id or an existing feedback was sent again
			//       check if the new and old feedbacks are the same
			// for now just end
			log.Println("duplicate id uploaded")
			return nil
		} else {
			return err
		}
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.Encode(submission)
	log.Println("wrote file")
	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ServerMain() {

	http.HandleFunc("/feedback", httpFeedback)

	http.ListenAndServe(":8080", nil)
}
