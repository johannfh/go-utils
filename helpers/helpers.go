package helpers

import (
	"fmt"
	"strings"

	"github.com/johannfh/go-utils/internal/utils"
)


// Check wether or not the object is empty
func Empty(object interface{}) bool {
    return utils.IsEmpty(object)
}

func NotEmpty(object interface{}) bool {
    return !utils.IsEmpty(object)
}

// Generate a string consisting of '-'
// with the provided length
func Dashstring(length int) string {
    return strings.Repeat(string('-'), length)
}

// Add a padding string to the left of every
// non empty or whitespace line of the string
func PrependString(str string, pre string) string {
    lines := strings.Split(str, "\n")

    for i := range lines {
        if len(strings.TrimSpace(lines[i])) == 0 {
            continue
        }

        lines[i] = fmt.Sprintf("%s%s", pre, lines[i])
    }

    return strings.Join(lines, "\n")
}

