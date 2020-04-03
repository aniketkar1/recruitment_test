package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var str []string

func htmlParser(r io.Reader) {
	z, _ := html.Parse(r)
	dfs(z)
	//fmt.Println(str)

}

func text(n *html.Node) {
	if n.Type == html.TextNode {
		if len(strings.TrimSpace(n.Data)) > 1 {
			str = append(str, strings.TrimSpace(n.Data)) //getting text from the html node and its child
		}
	}
	if n.Type != html.ElementNode {
		return
	}
	for tt := n.FirstChild; tt != nil; tt = tt.NextSibling {
		text(tt)
	}
}

func dfs(n *html.Node) { // travelling through every node in html

	if n.Type == html.ElementNode && n.Data == "md-card" { //getting text from md-card
		text(n)
	}

	for tt := n.FirstChild; tt != nil; tt = tt.NextSibling {
		dfs(tt)
	}
}

func main() {
	link := flag.String("link", "https://summerofcode.withgoogle.com/archive/2019/projects/", "link to the site [ string ]")
	flag.Parse()
	res, err := http.Get(*link)
	if err != nil {
		panic(err)
	}

	htmlParser(res.Body)

	vals := make([][]string, len(str)/4)

	for i := 0; i < 400; i++ {
		if i%4 != 3 {
			vals[i/4] = append(vals[i/4], str[i]) //converting []str to [][]vals
		}
	}
	for i := 0; i < len(vals); i++ { // getting correct format i.e name , organisation , project
		temp := vals[i][1]
		vals[i][1] = vals[i][2]
		vals[i][2] = temp
	}

	csvFile, err := os.Create("project.csv") //creating csv file

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for _, val := range vals {
		_ = csvwriter.Write([]string(val))
	}
	csvwriter.Flush()
	csvFile.Close()

}
