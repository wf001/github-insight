package util

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime/debug"
	"sort"
	"time"

	"runtime"

	"github.com/k0kubun/pp"
)

func Info(v interface{}) {
	// debug utility
	_, file, line, _ := runtime.Caller(1)
	reg := "[/]"
	files := regexp.MustCompile(reg).Split(file, -1)
	fmt.Printf("\033[36m###### %s %d ########\033[0m\n", files[len(files)-1], line)
	pp.Println(v)
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func ParseDecSep(s string) string {
	// add decimal separater to s, then return it.
	// if 1000 recieved, then 1,000 returned.
	res := []rune{}
	l := len(s)
	_s := []rune(s)
	for i := 0; i < l; i += 1 {
		if (l-i)%3 == 0 && i != 0 {
			res = append(res, rune(','))
		}
		res = append(res, _s[i])
	}

	return string(res)
}

func Contains(src string, tgt []string) bool {
	//return whether src contains one of the tgt element or not
	for _, t := range tgt {
		if t == src {
			return true
		}
	}
	return false
}
func GetSortedKeyByValue(m map[string]int) []string {
	// return a list of key sorted by a map value
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}
func GetSortedKeyByKey(m map[string]int) []string {
	// return a list sorted by a map key
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func ParseToDate(s string) string {
	// convert date-time format to date format
	t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", s)
	s = t.Format("2006-01-02")
	return s
}
func Error(msg string) {
	// print Error to standard error
	debug.PrintStack()
	fmt.Fprintln(os.Stderr, fmt.Sprintf("[ERROR] %s", msg))
	os.Exit(1)
}
