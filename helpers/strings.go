package helpers

import (
	"fmt"
	"strings"
)

// Generate a string consisting of '-'
// with the provided length
func Dashstring(length int) string {
    return strings.Repeat(string('-'), length)
}

// Check wether or not a string is empty or only
// consists of whitespace as defined by Unicode
func IsEmpty(s string) bool {
    return len(strings.TrimSpace(s)) == 0
}

// Add a padding string to the left of every
// non empty or whitespace line of the string
func PrependString(str string, pre string) string {
    lines := strings.Split(str, "\n")

    for i := range lines {
        if IsEmpty(lines[i]) {
            continue
        }

        lines[i] = fmt.Sprintf("%s%s", pre, lines[i])
    }

    return strings.Join(lines, "\n")
}

