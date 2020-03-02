package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//initPages
func initPages() {
	http.HandleFunc(MainSiteURL+"/", indexPage)
	GetCatagories()
}

func indexPage(w http.ResponseWriter, r *http.Request) {

}

var catagories []string

// GetCatagories - gets all of the catagories in /data
func GetCatagories() {
	catagories = nil
	//catagories = append(catagories, "Blah", "Blah")
	//catagories = append(catagories, "Blah")
	count := 0
	err := filepath.Walk(ExecPath+"/data/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			count = count + 1
			//fmt.Println(count)
			if count > 1 {
				path = strings.ReplaceAll(path, ExecPath+"\\data\\", "")
				path = strings.ReplaceAll(path, ".csv", "")
				if path == "" {
				} else if path == "\n" {
				} else {
					//fmt.Println(path, info.Size())
					//fmt.Printf("length: %d\n", len(catagories))
					catagories = append(catagories, path)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Loaded Catagories")
	}
}
