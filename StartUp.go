package main

import (
	"fmt"
	"log"

	//"net"
	//"net/http"
	"os"
	//"os/exec"
	//"path/filepath"
	//"runtime"
	"io"
)

// StartUp - This is run when the program begins. Checks to see if files and directories exist. Useful for deployment.
func StartUp() {
	FileStore()
	if _, err := os.Stat(ExecPath + "/data"); os.IsNotExist(err) {
		fmt.Printf("/data directory created.\n")
		err := os.Mkdir(ExecPath+"/data", 0755)
		check(err)
	}
	if _, err := os.Stat(ExecPath + "/html"); os.IsNotExist(err) {
		fmt.Printf("/html directory created.\n")
		err := os.Mkdir(ExecPath+"/html", 0755)
		check(err)
	}
	if _, err := os.Stat(ExecPath + "/HTTPS-key"); os.IsNotExist(err) {
		fmt.Printf("/HTTPS-key directory created.\n")
		err := os.Mkdir(ExecPath+"/HTTPS-key", 0755)
		check(err)
	}
	if _, err := os.Stat(ExecPath + "/html/index.html"); os.IsNotExist(err) {
		err := WriteToFile(ExecPath+"/html/index.html", PageIndex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("/html/index.html directory created.\n")
	}
	if _, err := os.Stat(ExecPath + "/html/assets"); os.IsNotExist(err) {
		fmt.Printf("/html/assets directory created.\n")
		err := os.Mkdir(ExecPath+"/html/assets", 0755)
		check(err)
	}
	if _, err := os.Stat(ExecPath + "/html/assets/css"); os.IsNotExist(err) {
		fmt.Printf("/html/assets/css directory created.\n")
		err := os.Mkdir(ExecPath+"/html/assets/css", 0755)
		check(err)
	}
	if _, err := os.Stat(ExecPath + "/html/assets/css/styles.css"); os.IsNotExist(err) {
		err := WriteToFile(ExecPath+"/html/assets/css/styles.css", cssIndex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("/html/assets/css/styles.css directory created.\n")
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
