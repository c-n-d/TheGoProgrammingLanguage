/*
Exercise 3.12 reports if two strings are anagrams (ignoring case and spacing)

$ go run src/ch3/ex_3_12/main.go
anagram("Quid est veritas", "Est vir qui adest") = true
anagram("Is this an anagram", "No it is not") = false
anagram("Augustus De Morgan", "Great Gun do us a sum") = true
*/

package main

import (
    "fmt"
    "sort"
    "strings"
)

type Pair struct {
    a, b string
}

func main() {
    ps := []Pair{Pair{"Quid est veritas", "Est vir qui adest"},
                 Pair{"Is this an anagram", "No it is not"},
                 Pair{"Augustus De Morgan", "Great Gun do us a sum"}}

    for _, p := range ps {
        fmt.Printf("anagram(\"%s\", \"%s\") = %t\n", p.a, p.b, anagram(p.a, p.b))
    }
}

func anagram(s1, s2 string) bool {
    s1 = clean(s1)
    s2 = clean(s2)

    if len(s1) != len(s2) {
        return false
    }

    return sortString(s1) == sortString(s2)
}

func sortString(s string) string {
    sArr := strings.Split(s, "")

    sort.Strings(sArr)

    return strings.Join(sArr, "")
}

// Ignore spaces and case
func clean(s string) string {
    return strings.ToLower(strings.Replace(s, " ", "", -1))
}
