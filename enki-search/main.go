package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func buildEnkiQuery(queryParams map[string]string) string {
	req := "https://api.enkigis.com/volumes?"

	for key, element := range queryParams {
		req += "q." + key + "=" + element + "&"
	}

	return req
}

func main() {
	// Set up query
	queryParams := map[string]string{
		"target":     "mars",
		"instrument": "ultraviolet",
	}

	req := buildEnkiQuery(queryParams)

	// Run query
	resp, err := http.Get(req)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("GET request to Enki successful")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Format
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, body, "", "\t"); err != nil {
		log.Fatalln(err)
	}

	// Write to file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err := f.WriteString(dst.String()); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Successfully wrote request response to file")
	}
}
