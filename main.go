package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Cdktf_json_file struct {
	Language           string   `json:"language"`
	App                string   `json:"app"`
	TerraformProviders []string `json:"terraformProviders"`
	TerraformModules   []string `json:"terraformModules"`
	ProjectId          string   `json:"projectId"`
	CodeMakerOutput    string   `json:"codeMakerOutput"`
}

func main() {

	var data Cdktf_json_file

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	var path = filepath.Join(cwd, fmt.Sprint("cdktf.json"))

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File exists\n")
		err = os.Remove(path)
		if err != nil {
			log.Fatalf("There was a error deleting the file => %s", err.Error())
		}
	} else {
		fmt.Printf("File does not exist\n")
	}
	language := os.Args[1]
	data.App = "echo 'hello'"
	data.ProjectId = "052b874a-f94b-485f-82c8-ad1dc6482680"
	data.TerraformProviders = []string{
		"aws@~> 3.63",
		"google@~> 3.89",
	}

	if language == "typescript" {
		data.Language = "typescript"
		data.CodeMakerOutput = ".gen"
	} else if language == "python" {
		data.Language = "python"
		data.CodeMakerOutput = "imports"
	}
	buffer := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	_ = enc.Encode(data)

	// json_data, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatalf("Unable to parse JSON => %s", err.Error())
	// }

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		log.Printf("Unable to open the temp file %s", err.Error())
		return
	}
	defer file.Close()

	// fmt.Print(string(data))
	_, err = file.Write(buffer.Bytes())

	if err != nil {
		log.Fatalf("Unable to parse JSON => %s", err.Error())
	}

}
