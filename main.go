package main

import (
	"fmt"

	"example.com/hello/errors"
)

func main() {
	_, err := metodoUm()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func metodoUm() (string, *errors.TraceableError) {
	_, err := metodoDois()

	if err != nil {
		return "", err.Throw()
	}

	return "retorno1", nil
}

func metodoDois() (string, *errors.TraceableError) {
	_, err := metodoTres()

	if err != nil {
		return "", err.Throw()
	}

	return "retorno2", nil
}

func metodoTres() (string, *errors.TraceableError) {
	_, err := metodoQuatro()

	if err != nil {
		return "", err.Throw()
	}

	return "retorno3", nil
}

func metodoQuatro() (string, *errors.TraceableError) {
	err := metodoCinco()

	if err != nil {
		return "", err.Throw()
	}

	return "retorno4", nil
}

func metodoCinco() *errors.TraceableError {
	fmt.Println("Hello, Errors!")

	return errors.Throw("FATAL ERROR on metodoCinco")
}
