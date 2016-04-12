/*

*/

package github

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func CreateIssue() {
    
}

func ReadIssue() {
    scanner := bufio.NewScanner(os.Stdin)
    owner := readFromStdin(scanner, "Please enter the owner:")
    repo := readFromStdin(scanner, "Please enter the repo name:")

    url := strings.Replace(IssuesListURL, ":owner", owner, 1)
    url = strings.Replace(url, ":repo", repo, 1)

    resp, err := getWithErrorCheck(url)

    if err != nil {
        log.Fatal(err)
    }

    var result []Issue
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        fmt.Fprintf(os.Stderr, "readissues: %v\n", err)
    }

    for _, res := range result {
        fmt.Printf("%v\n", res)
    }
    resp.Body.Close()
}

func readFromStdin(in *bufio.Scanner, prompt string) string {
    fmt.Println(prompt)
    in.Scan()
    return in.Text()
}

func UpdateIssue() {
    
}

func DeleteIssue() {
    
}

func getWithErrorCheck(url string) (*http.Response, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }

    // We must close resp.Body on all execution paths.
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    return resp, nil
}