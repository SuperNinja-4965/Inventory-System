package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type NewCat struct {
	CatName string
}

func initNew() {
	http.HandleFunc("/newCategory/", newCategory)
	http.HandleFunc("/newItem/", newItem)
}

func newCategory(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles(ExecPath + "/html/NewCategory.html"))
	if r.Method != http.MethodPost {
		p := MainIndexPage{Data: template.HTML("<center><h1 style=\"color:white;\">New Category</h1>                <div class=\"container\">                    <br>                    <form method=\"POST\">                        <table>                            <thead>                                <tr>                                    <th>                                        <h2 style=\"color:white;\">Type</h2></th>                                    <th>                                        <h2 style=\"color:white;\">Value</h2></th>                                </tr>                            </thead>                            <tbody>                                <tr>                                    <td>                                        <h3 style=\"color:white;\">Category Name:</h3></td>                                    <td>                                        <input type=\"textbox\" name=\"CatName\" id=\"CatName\" placeholder=\"Category Name\" style=\"display: inline-block;\">                                    </td>                                </tr>                            </tbody>                        </table>                        <br>                        <input type=\"submit\" value=\"Create Category\"> </form> 						</div></center>"), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
		//tmpl.Execute(w, nil)
		return
	}

	CategoryNew := NewCat{
		CatName: r.FormValue("CatName"),
	}

	if _, err := os.Stat(ExecPath + "/data/" + CategoryNew.CatName + ".csv"); os.IsNotExist(err) {
		err := WriteToFile(ExecPath+"/data/"+CategoryNew.CatName+".csv", "item1,100f,100,10,\"This is a cool item, and it always will be.\"")
		if err == nil {
			fmt.Printf("/data/" + CategoryNew.CatName + ".csv" + " file created.\n")
			// if successfully created cat.
			p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:white;\">Category has been created</h1> </center>"), ProjectName: ProgramName}
			t, _ := template.ParseFiles(ExecPath + "/html/index.html")
			t.Execute(w, p)
		} else {
			// If error creating cat
			p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:white;\">There was an error with creating that category</h1> </center>"), ProjectName: ProgramName}
			t, _ := template.ParseFiles(ExecPath + "/html/index.html")
			t.Execute(w, p)
		}
	} else {
		// If file exists
		p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:white;\">Category already exists</h1> </center>"), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
	}
}

func newItem(w http.ResponseWriter, r *http.Request) {

}
