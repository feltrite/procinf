package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bbrks/wrap"
	"io/ioutil"
	"jaytaylor.com/html2text"
	"net/http"
	"os"
	"strings"
)


func get_file_info(process string) (string){
	url :=  fmt.Sprintf("https://file.net/process/%s.html", strings.ToLower(process))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("NO INFO AVAILABLE ON PROCESS %S ", process)
	}
	content, err := ioutil.ReadAll(resp.Body)
	return string(content)
}

func extractHTMLsection1(htmldocstring string) (string){
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmldocstring))
	if err != nil {
		fmt.Printf("%s", err)
	}
	wanted := doc.Find(".maxwidth").Eq(2)
	wanted.Find("p").Eq(2).Remove()
	wanted.Find("p").Eq(2).Remove()
	tmp, _:= wanted.Html()
	return tmp
}

func extractHTMLsection2(htmldocstring string) (string){
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmldocstring))
	if err != nil {
		fmt.Printf("%s", err)
	}
	wanted := doc.Find(".maxwidth").Eq(3)
	elements := wanted.Find("h2")
	elements.AppendSelection(wanted.Find("p").Eq(0))
	tmp, _:= elements.Html()
	return tmp
}


func formatText(badtext string) (string){
	tmptext, _ := html2text.FromString(badtext, html2text.Options{
		PrettyTables: false,
		OmitLinks:true})
	tmptext = wrap.Wrap(tmptext, 120)
	return tmptext
}

func printInfo(htmldocstring string) {
	section1 := formatText(extractHTMLsection1(htmldocstring))
	section2 := formatText(extractHTMLsection2(htmldocstring))

	fmt.Println(section1)
	fmt.Println(section2)
}


func Usage() {
	fmt.Printf("Usage: %s [OPTIONS] <process name> \n", os.Args[0])
}


func main() {
	fmt.Println("Procinf - Process Info (ver 0.1)")
	args := os.Args

	if len(args) < 2 {
		Usage()
	} else {
		stringRes := get_file_info(args[1])
		printInfo(stringRes)
	}
}
