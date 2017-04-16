package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Generic error checking method
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Writes contents of filled struct to file
func (i *instaBotConfig) writeConfigToFile(file *os.File) {
	file.WriteString(initString)
	rv := reflect.ValueOf(i)
	num := rv.Elem().NumField()
	file.WriteString("user_blacklist={},\r")
	for i := 0; i < num; i++ {
		fi := rv.Elem().Field(i)
		sf := rv.Elem().Type().Field(i).Name
		switch fi.Kind() {
		case reflect.String:
			file.WriteString(sf + "='" + fi.String() + "',\r")
		case reflect.Int:
			file.WriteString(sf + "=" + strconv.FormatInt(fi.Int(), 10) + ",\r")
		case reflect.Slice:
			s := "["
			arr := strings.Split(fi.Index(0).String(), ",")
			for _, val := range arr {
				s += "'" + val + "'" + ","
			}
			s = strings.TrimRight(s, ",") + "]"
			file.WriteString(sf + "=" + s)
			if i == (num - 1) {
				file.WriteString(")\r")
			} else if i == 5 {
				file.WriteString(",\r user_blacklist={},\r")
			} else {
				file.WriteString(",\r")
			}

		default:
			continue
		}
	}
	file.WriteString(endString)
}

// Serves generated config.py file to client - downloads in browser
func downloadFile(res http.ResponseWriter, req *http.Request, file *os.File) {
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	file.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := file.Stat()                         //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	res.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	res.Header().Set("Content-Type", FileContentType)
	res.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	file.Seek(0, 0)
	io.Copy(res, file) //'Copy' the file to the client
}

// Displays config form or creates config depending on request type
func createConfig(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		t, _ := template.ParseFiles("create.gtpl")
		t.Execute(res, nil)
	} else if req.Method == "POST" {
		req.ParseForm()
		// Create and open file for writing
		file, err := os.Create("config.py")
		defer file.Close()
		check(err)

		// Create new instagram config, load form data into struct, and write result to file
		instConfig := new(instaBotConfig)
		LoadModel(instConfig, req.Form)
		instConfig.writeConfigToFile(file)

		// Download config file
		downloadFile(res, req, file)
	}
}

// Main method for routing requests
func main() {
	http.HandleFunc("/", createConfig)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	// http.ListenAndServe(":5000", nil)
}
