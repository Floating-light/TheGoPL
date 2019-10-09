package ch4

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

var data []byte

//GoStruct2JSON convert go struct to JSON, using json.Marshal
func GoStruct2JSON() {
	var err error
	data, err = json.Marshal(movies)
	//data, err := json.MarshalIndent(movies, "", "    ") //format json output
	if err != nil {
		log.Fatalf("JSON marshling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

//JSON2GoStruct convert json to go struct
func JSON2GoStruct() {
	GoStruct2JSON()
	var titles []struct {
		Title    string
		Released int
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%#v\n", titles)
}

//github issue example  <<<<----------------------------------------------------------------

//IssueURL github issue url
const IssuesURL = "https://api.github.com/search/issues"

//IssuesSearchResult issues search result struct
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

//Issue issue struct
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

//User github user infomation struct
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//SearchIssues queries the Github issue tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//We must close resp.Body on all execution paths.
	//(chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	//resp.Body is a streaming
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//FalseMain excute github example
func FalseMain() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
