package errors

import (
	"fmt"
	"runtime"
	"strconv"
)

// traceableError implements the basic traceable error.
type TraceableError struct {
	Err   error
	Stack []stackTrace
}

type stackTrace struct {
	file         string
	line         int
	functionName string
}

func (t TraceableError) Error() string {
	var e string

	e = "Error Message: " + t.Err.Error() + "\nStacktrace: \n"

	for i := (len(t.Stack) - 1); i >= 0; i-- {

		e = e + t.Stack[i].file + ":" + strconv.Itoa(t.Stack[i].line) + "\n └─" + t.Stack[i].functionName + "()\n"

	}

	return e
}

func (t *TraceableError) Throw() *TraceableError {

	if pc, file, line, ok := runtime.Caller(1); ok {
		function := runtime.FuncForPC(pc)

		var functionName string
		if function != nil {
			functionName = function.Name()
		}

		st := stackTrace{
			file:         file,
			line:         line,
			functionName: functionName,
		}

		t.Stack = append(t.Stack, st)

		return t
	}
	return nil
}

func Throw(msg string) *TraceableError {

	e := fmt.Errorf(msg)

	if pc, file, line, ok := runtime.Caller(1); ok {
		function := runtime.FuncForPC(pc)
		var functionName string

		if function != nil {
			functionName = function.Name()
		}

		t := TraceableError{
			Err: e,
			Stack: []stackTrace{
				{
					file:         file,
					line:         line,
					functionName: functionName,
				},
			},
		}
		return &t
	}
	return nil
}
