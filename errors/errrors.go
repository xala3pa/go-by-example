package main

import (
	"errors"
	"fmt"
)

/*
* By convention, errors are the last return value
* and have type error, a built-in interface.
 */

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("42 not works")
	}
	return arg + 10, nil
}

type customError struct {
	num   int
	eText string
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.num, e.eText)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &customError{arg, "can't work with it"}
	}
	return arg + 10, nil
}

func main() {

	for _, val := range []int{1, 42, 7} {
		if res, e := f1(val); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 works: ", res)
		}
	}

	for _, val := range []int{1, 42, 7} {
		if res, e := f2(val); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 works: ", res)
		}
	}

	/*
	* If you want to programmatically use the data in a custom error,
	* you’ll need to get the error as an instance of the custom error
	* type via type assertion.
	 */
	_, e := f2(42)

	if er, ok := e.(*customError); ok {
		fmt.Println(er.num)
		fmt.Println(er.eText)
	}
}
