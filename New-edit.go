package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type NewCat struct {
	CatName string
}
type NewCount struct {
	Counts string
}

func initNew() {
	http.HandleFunc("/newCategory/", newCategory)
	http.HandleFunc("/editCategory/", makeHandlerEdit(editCategory))
}

var NewCatForm string = "<center><h1 style=\"color:white;\">New Category</h1>                <div class=\"container\">                    <br>                    <form method=\"POST\">                        <table>                            <thead>                                <tr>                                    <th>                                        <h2 style=\"color:white;\">Type</h2></th>                                    <th>                                        <h2 style=\"color:white;\">Value</h2></th>                                </tr>                            </thead>                            <tbody>                                <tr>                                    <td>                                        <h3 style=\"color:white;\">Category Name:</h3></td>                                    <td>                                        <input type=\"textbox\" name=\"CatName\" id=\"CatName\" placeholder=\"Category Name\" style=\"display: inline-block;\">                                    </td>                                </tr>                            </tbody>                        </table>                        <br>                        <input type=\"submit\" value=\"Create Category\"> </form> 						</div></center>"

func newCategory(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles(ExecPath + "/html/NewCategory.html"))
	if r.Method != http.MethodPost {
		p := MainIndexPage{Data: template.HTML(NewCatForm), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
		//tmpl.Execute(w, nil)
		return
	}

	CategoryNew := NewCat{
		CatName: r.FormValue("CatName"),
	}
	if CategoryNew.CatName == "" {
		p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:red;\">Category Name cannot be blank</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
	} else {
		if _, err := os.Stat(ExecPath + "/data/" + CategoryNew.CatName + ".csv"); os.IsNotExist(err) {
			err := WriteToFile(ExecPath+"/data/"+CategoryNew.CatName+".csv", "item1,100f,100,10,\"This is a cool item, and it always will be.\"")
			if err == nil {
				fmt.Printf("/data/" + CategoryNew.CatName + ".csv" + " file created.\n")
				// if successfully created cat.
				p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:green;\">Category has been created</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
				t, _ := template.ParseFiles(ExecPath + "/html/index.html")
				t.Execute(w, p)
			} else {
				// If error creating cat
				p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:red;\">There was an error with creating that category</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
				t, _ := template.ParseFiles(ExecPath + "/html/index.html")
				t.Execute(w, p)
			}
		} else {
			// If file exists
			p := MainIndexPage{Data: template.HTML("<center> <h1 style=\"color:red;\">Category already exists</h1> </center> <br>" + NewCatForm), ProjectName: ProgramName}
			t, _ := template.ParseFiles(ExecPath + "/html/index.html")
			t.Execute(w, p)
		}
	}
}

func editCategory(w http.ResponseWriter, r *http.Request, EditURL string) {
	//fmt.Println(EditURL)
	if EditURL == "" {
		//Load select category page.
		if r.Method == http.MethodPost {
			SCategory := NewCat{
				CatName: r.FormValue("SelectedCategory"),
			}
			http.Redirect(w, r, "/editCategory/"+SCategory.CatName, 303)
			return
		}
		var PreSelectHTML string = "<center><h1 style=\"color:white;\">Select a Category</h1><div class=\"container\"><br><form method=\"POST\"><table><thead><tr><th><h2 style=\"color:white;\">Type</h2></th><th><h2 style=\"color:white;\">Value</h2></th></tr></thead><tbody><tr><td><h3 style=\"color:white;\">Category:</h3></td><td><select id=\"SelectedCategory\" name=\"SelectedCategory\">"
		var PostSelectHTML string = "</select></td></tr></tbody></table><br><input type=\"submit\" value=\"Select Category\"> </form></div></center>"
		var MiddleSelectHTML string = ""
		GetCatagories()
		if len(catagories) != 0 {
			for i := 0; i <= len(catagories)-1; i++ {
				MiddleSelectHTML = MiddleSelectHTML + " <option value=\"" + catagories[i] + "\">" + catagories[i] + "</option>"
			}
		} else {
			MiddleSelectHTML = ""
		}

		var outputDisplay string = PreSelectHTML + MiddleSelectHTML + PostSelectHTML
		p := MainIndexPage{Data: template.HTML(outputDisplay), ProjectName: ProgramName}
		t, _ := template.ParseFiles(ExecPath + "/html/index.html")
		t.Execute(w, p)
	} else {
		if _, err := os.Stat(ExecPath + "/data/" + EditURL + ".csv"); os.IsNotExist(err) {
			//Throw error if category is not found. Need to add category selection options here
			var outputDisplay string = "<center> <h1 style=\"color:red;\">There was an error with locating that category</h1> </center> <br>"
			p := MainIndexPage{Data: template.HTML(outputDisplay), ProjectName: ProgramName}
			t, _ := template.ParseFiles(ExecPath + "/html/index.html")
			t.Execute(w, p)
		} else {
			var PostSuccess string = "N/A"
			if r.Method == http.MethodPost {
				counts := NewCount{
					Counts: r.FormValue("Count"),
				}
				Reported_count, _ := strconv.Atoi(counts.Counts)
				// Use os.Create to create a file for writing.
				f, _ := os.Create(ExecPath + "/data/" + EditURL + ".csv")
				// Create a new writer.
				w := bufio.NewWriter(f)
				for i := 1; i <= Reported_count; i++ {
					iToString := strconv.Itoa(i)
					w.WriteString("\"" + r.FormValue(iToString+"-1") + "\",")
					w.WriteString("\"" + r.FormValue(iToString+"-2") + "\",")
					w.WriteString(r.FormValue(iToString+"-3") + ",")
					w.WriteString(r.FormValue(iToString+"-4") + ",")
					w.WriteString("\"" + r.FormValue(iToString+"-5") + "\"\n")
				}
				// Flush.
				w.Flush()
				PostSuccess = "Success"
			}

			//Display entire categories data in an editable table.
			var count int = 1
			var outputPrepares string = ""
			csvfile, err := os.Open(ExecPath + "/data/" + EditURL + ".csv")
			if err != nil {
				log.Fatalln("Couldn't open the csv file", err)
			}
			// Parse the file
			r := csv.NewReader(csvfile)
			//r := csv.NewReader(bufio.NewReader(csvfile))
			// Iterate through the records
			for {
				// Read each record from csv
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				// Form the html for the table.
				var RowTemplate1 string = "<tr><td><input id=\"" + strconv.Itoa(count) + "-1\" name=\"" + strconv.Itoa(count) + "-1\" size=\"10%\" style=\"background-color:transparent;color:white;border:0;\" value=\""
				var RowTemplate2 string = "\"></td><td><input id=\"" + strconv.Itoa(count) + "-2\" name=\"" + strconv.Itoa(count) + "-2\" size=\"6%\" style=\"background-color:transparent;color:white;border:0;\" value=\""
				var RowTemplate3 string = "\"></td><td><input id=\"" + strconv.Itoa(count) + "-3\" name=\"" + strconv.Itoa(count) + "-3\" size=\"6%\" style=\"background-color:transparent;color:white;border:0;\" value=\""
				var RowTemplate4 string = "\"></td><td><input id=\"" + strconv.Itoa(count) + "-4\" name=\"" + strconv.Itoa(count) + "-4\" size=\"6%\" style=\"background-color:transparent;color:white;border:0;\" value=\""
				var RowTemplate6 string = "\"></td><td><input id=\"" + strconv.Itoa(count) + "-5\" name=\"" + strconv.Itoa(count) + "-5\" size=\"66%\" style=\"background-color:transparent;color:white;border:0;\" value=\""
				var RowTemplateEnd string = "\"></td></tr>"
				outputPrepares = outputPrepares + RowTemplate1 + record[0] + RowTemplate2 + record[1] + RowTemplate3 + record[2] + RowTemplate4 + record[3] + RowTemplate6 + record[4] + RowTemplateEnd
				count = count + 1
			}
			count = count - 1
			var CATNAME string = EditURL
			var ErrorReport string = ""
			if PostSuccess == "Success" {
				ErrorReport = "<h2 style=\"color:Green;\">Data was saved successfully!</h2><br>"
			} else if PostSuccess == "Error" {
				ErrorReport = "<h2 style=\"color:Red;\">Data was NOT saved.</h2><br>"
			}
			var templatePart1 string = "<center><h1 style=\"color:white;\">Editing Category: " + CATNAME + "</h1><div class=\"container\"><br><form method=\"POST\"><h2 style=\"color:white;\">Record Count: " + strconv.Itoa(count) + "</h2><br>" + ErrorReport + "<input type=\"hidden\" id=\"Count\" name=\"Count\" value=\"" + strconv.Itoa(count) + "\"><table><thead><tr><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Notes</th></tr></thead><tbody>"
			var templatePartEnd string = "</tbody></table><br><input type=\"submit\" value=\"Save Changes\"> </form></div></center>"
			var outputDisplay string = templatePart1 + outputPrepares + templatePartEnd
			p := MainIndexPage{Data: template.HTML(outputDisplay), ProjectName: ProgramName}
			t, _ := template.ParseFiles(ExecPath + "/html/index.html")
			t.Execute(w, p)
		}
	}
}

func makeHandlerEdit(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urll := r.URL.Path
		// if urll == "/editCategory/" {
		// 	http.Redirect(w, r, "/", 301)
		// } else if urll == "/editCategory" {
		// 	http.Redirect(w, r, "/", 301)
		// } else {
		urll = strings.Replace(urll, "/editCategory/", "", -1)
		urll = strings.ReplaceAll(urll, "/", "")
		//fmt.Println(urll)
		fn(w, r, urll)
		//}
	}
}
