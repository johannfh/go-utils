package helpers

import (
	"fmt"
	"strings"

	"github.com/johannfh/go-utils/internal/utils"
)

func Empty(object interface{}) bool {
	return utils.IsEmpty(object)
}

func NotEmpty(object interface{}) bool {
	return !utils.IsEmpty(object)
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

func MergeFuncs(funcs ...func()) func() {
	return func() {
		for _, fn := range funcs {
			fn()
		}
	}
}
