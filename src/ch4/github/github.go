/*
Package github provides a Go API for the GitHub issue tracker.
See https://developer.github.com/v3/search/#search-issues
*/

package github

import "time"

const (
    APIBase = "https://api.github.com"
    IssuesSearchURL = APIBase + "/search/issues"
    IssuesListURL = APIBase + "/repos/:owner/:repo/issues"
)

type IssuesSearchResult struct {
    TotalCount int `json:"total_count"`
    Items      []*Issue
}

type Issue struct {
    Number int
    HTMLURL  string `json:"html_url"`
    Title    string
    State    string
    User     *User
    CreateAt time.Time `json:"created_at"`
    Body     string // In markdown format
}

type User struct {
    Login string
    HTMLURL string `json:"html_url"`
}
