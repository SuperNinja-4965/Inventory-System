package main

import (
	"html/template"
	"net/http"
)

func CatagoryPage(w http.ResponseWriter, r *http.Request) {
	GetCatagories()
	var cats string
	//fmt.Println(len(catagories))
	if len(catagories) != 0 {
		for i := 0; i <= len(catagories)-1; i++ {
			cats = cats + ItemView("/"+catagories[i], catagories[i], "1")
		}
		p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
	} else {
		cats = cats + ItemView("", "NO ITEMS FOUND", "0")
		p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
	}
}
