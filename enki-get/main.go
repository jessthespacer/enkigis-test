package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func buildEnkiGet(uid int) string {
	return "https://api.enkigis.com/volumes/" + fmt.Sprintf("%d", uid)
}

func main() {
	// Get volume
	uid := 10012
	req := buildEnkiGet(uid)

	// Run request
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
