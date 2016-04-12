/*
The SearchIssues function calls GitHub's '/search/issues' API, does error handling,
and returns the resulting list of Issues. 100 per page.

https://developer.github.com/v3/search/#search-issues
*/

package github

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(IssuesSearchURL + "?q=" + q)
    if err != nil {
        return nil, err
    }

    // We must close resp.Body on all execution paths.
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }

    resp.Body.Close()
    return &result, nil
}
