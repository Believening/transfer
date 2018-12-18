/*
字符串比较

You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.  The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.  The digit-logs should be put in their original order.

Return the final order of the logs.



Example 1:

Input: ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4 7"]

 */
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

//func reorderLogFiles(logs []string) []string {
//	letterParten := "^\\w*[0-9]*\\s([a-z][a-z\\s]*)"
//	letterRgx := regexp.MustCompile(letterParten)
//	letterRgx.Longest()
//	letterMap := make(map[string]string)
//	var digitSlice, letterSlice []string
//
//	for _, log := range logs {
//		if letterRgx.MatchString(log) {
//			matched := letterRgx.FindStringSubmatch(log)
//			letterSlice = append(letterSlice, matched[1])
//			letterMap[matched[1]] = matched[0]
//			continue
//		}
//
//		digitSlice = append(digitSlice, log)
//	}
//	sort.Strings(letterSlice)
//
//	for i, v := range letterSlice {
//		letterSlice[i] = letterMap[v]
//	}
//
//	return append(letterSlice, digitSlice...)
//}

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		idx1 := strings.Index(logs[i], " ")
		idx2 := strings.Index(logs[j], " ")
		isDigit1 := unicode.IsDigit(rune(logs[i][idx1+1]))
		isDigit2 := unicode.IsDigit(rune(logs[j][idx2+1]))
		var ret bool
		switch {
		case isDigit1 && isDigit2:
			ret = false
		case isDigit1 && !isDigit2:
			ret = false
		case !isDigit1 && isDigit2:
			ret = true
		case !isDigit1 && !isDigit2:
			ret = logs[i][idx1+1:] < logs[j][idx2+1:]
		}
		return ret
	})
	return logs
}

func main() {
	logs := []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"}
	out := reorderLogFiles(logs)
	fmt.Println(out)

}
