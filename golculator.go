package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Store the result in this struct
type CalcResult struct {
	CResult int
}

// index is a simple function using a template to display
// on browser. There are two important parameter:
// w http.ResponseWriter: Responsible to WRITE data to the Browser
// r *http.Request: Responsible to GET data from the browser
func index(w http.ResponseWriter, r *http.Request) {

	// Create the HTML Template
	tmpl := template.New("HTML Calculator")

	// Parse HTML content
	tmpl, _ = tmpl.Parse(
		`<html><form method="post" action="/">
      <input type="text" name="val1" />
      <select name="op">
        <option value="+">+</option>
        <option value="-">-</option>
        <option value="/">/</option>
        <option value="*">*</option>
      </select>
      <input type="text" name="val2" />
      <label value="result">=</label>
      <label name="result">{{.CResult}}</label>
      <input type="submit" value="Go!" />
    </form>
    </html>`)

	// Execute the template
	// w = Write HTML to browser
	tmpl.Execute(w, nil)

	// Get HTML data
	// r = Request from form
	getVal := r.FormValue("val1")  // Request 1st value
	getOp := r.FormValue("op")     // Request the operator
	getVal2 := r.FormValue("val2") // Request 2nd value

	// Convert values to INT
	getValInt, _ := strconv.Atoi(getVal)
	getVal2Int, _ := strconv.Atoi(getVal2)

	// Test the requested operator and perform
	// operations as well
	if getOp == "+" {
		result := getValInt + getVal2Int
		r := CalcResult{CResult: result} // Put result in the struct
		tmpl.Execute(w, r)               // Execute the template again showing the result
		fmt.Println(result)

	} else if getOp == "-" {
		result := getValInt - getVal2Int
		r := CalcResult{CResult: result}
		tmpl.Execute(w, r)
		fmt.Println(result)

	} else if getOp == "/" {
		result := getValInt / getVal2Int
		r := CalcResult{CResult: result}
		tmpl.Execute(w, r)
		fmt.Println(result)

	} else if getOp == "*" {
		result := getValInt * getVal2Int
		r := CalcResult{CResult: result}
		tmpl.Execute(w, r)
		fmt.Println(result)
	}

}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":9000", nil)
}
