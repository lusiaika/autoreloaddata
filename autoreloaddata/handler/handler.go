package handler

import (
	"autoreloaddata/entity"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

const htmlPath = "html/web.html"
const jsonPath = "html/status.json"

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	var data entity.DataStatus
	w.Header().Add("Content Type", "text/html")
	// read from json file and write to webData
	file, _ := ioutil.ReadFile(jsonPath)
	json.Unmarshal(file, &data)
	templates, _ := template.ParseFiles(htmlPath)

	if data.Status.Water <= 5 || data.Status.Wind <= 6 {
		data.Status.DataStatus = "Aman"
	} else if (data.Status.Water >= 6 && data.Status.Water <= 8) || (data.Status.Wind >= 7 && data.Status.Wind <= 15) {
		data.Status.DataStatus = "siaga"
	} else if data.Status.Water > 8 || data.Status.Wind > 15 {
		data.Status.DataStatus = "bahaya"
	}

	context := data
	templates.Execute(w, context)
}

func GenerateToJson() {
	var datas entity.DataStatus
	for {
		datas.Status.Water = rand.Intn(100)
		datas.Status.Wind = rand.Intn(100)

		// write to json file
		jsonString, _ := json.Marshal(&datas)
		ioutil.WriteFile(jsonPath, jsonString, os.ModePerm)

		// sleep for 15 seconds
		time.Sleep(15 * time.Second)
	}
}
