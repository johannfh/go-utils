package assert

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"runtime/debug"

	"github.com/johannfh/go-utils/helpers"
)

type AssertData interface{
    Dump() string
}

// all of the collected [AssertData]
var assertData map[string]AssertData = map[string]AssertData{}

// add data which will be logged when an assert fails
func AddAssertData(key string, value AssertData) {
    assertData[key] = value
}

// remove data from being logged when an assert fails
func RemoveAssertData(key string) {
    delete(assertData, key)
}

// aggregates all of the log data and prints to [os.Stderr]
// then ends with a call to [os.Exit], stopping the program
func runAssert(msg string, args ...any) {

    slogValues := []any{
        "msg", msg,
        "area", "Assert",
    }
    slogValues = append(slogValues, args...)
    fmt.Fprintf(os.Stderr, "ARGS: %+s\n", args)

    for k, v := range assertData {
        slogValues = append(slogValues, k, v.Dump())
    }

    fmt.Fprintf(os.Stderr, "ASSERT\n")
    for i := 0; i < len(slogValues); i += 2 {
        fmt.Fprintf(os.Stderr, "\t%s=%v\n", slogValues[i], slogValues[i + 1])
    }

    fmt.Fprintln(os.Stderr, string(debug.Stack()))
    os.Exit(1)
}

// errors if [truth] is `false`
func Assert(truth bool, msg string, data ...any) {
    slog.Info("Assert Check", "truth", truth)
    if !truth {
        slog.Error("Assert#false encountered")
        runAssert(msg, data...)
    }
}

// errors if [a] is not equal to [b]
func Equal[T comparable](a, b T, msg string, data ...any) {
    slog.Info("Equal Check", "a", a, "b", b)

    if a != b {
        slog.Error("Equal#not equal encountered")
        runAssert(msg, data...)
    }
}

// errors if [a] is equal to [b]
func NotEqual[T comparable](a, b T, msg string, data ...any) {
    slog.Info("Not Equal Check", "a", a, "b", b)

    if a == b {
        slog.Error("NotEqual#equal encountered")
        runAssert(msg, data...)
    }
}

// errors if [s] is not empty
func Empty(s string, msg string, data ...any) {
    slog.Info("Empty Check", "s", s)

    if !helpers.IsEmpty(s) {
        slog.Error("EmptyCheck#not empty encountered")
        runAssert(msg, data...)
    }
}

// errors if [s] is empty
func NotEmpty(s string, msg string, data ...any) {
    slog.Info("Not Empty Check", "s", s)

    if helpers.IsEmpty(s) {
        slog.Error("NotEmpty#empty encountered")
        runAssert(msg, data...)
    }
}

// errors if [item] is not `nil`
func Nil(item any, msg string, data ...any) {
    slog.Info("Nil Check", "item", item)

    if item != nil {
        slog.Error("Nil#not nil encountered")
        runAssert(msg, data...)
    }
}


// errors if [item] is `nil`
func NotNil(item any, msg string, data ...any) {
    slog.Info("Not Nil Check", "item", item)

    if item == nil || reflect.ValueOf(item).Kind() == reflect.Ptr && reflect.ValueOf(item).IsNil() {
        slog.Error("NotNil#nil encountered")
        runAssert(msg, data...)
    }
}

