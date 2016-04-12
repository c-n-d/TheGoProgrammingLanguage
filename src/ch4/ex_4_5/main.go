/*
Exercise 4.5 - In place function to eliminate adjacent duplicates in a []string slice

1. Maintain two references i, j
2. i traverses the length of the slice
3. If the string at [i] and [i-1] differ, copy the string at [i] to [j]
4. Increament j when copying
5. Return elements [0:j]

$ go run src/ch4/ex_4_5/main.go
removeDuplicates([a b b c c c d d d d]) = [a b c d]
removeDuplicates([dog dog bird fish fish fish fish cat bird bird]) = [dog bird fish cat bird]
removeDuplicates([one two three]) = [one two three]
*/

package main

import "fmt"

func main() {
    inputs := [][]string{
        []string{"a", "b", "b", "c", "c", "c", "d", "d", "d", "d"},
        []string{"dog", "dog", "bird", "fish", "fish", "fish", "fish", "cat","bird", "bird"},
        []string{"one", "two", "three"}}

    for _, input := range inputs {
        tmp := make([]string, len(input))
        copy(tmp, input)

        fmt.Printf("removeDuplicates(%v) = %v\n", tmp, removeDuplicates(input))
    }
}

func removeDuplicates(s []string) []string {
    j := 1

    for i := 1; i < len(s); i++ {
        if s[i] != s[i-1] {
            s[j] = s[i]
            j++
        }
    }

    return s[:j]
}
