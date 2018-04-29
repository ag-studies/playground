package main

import (
	"fmt"
	"reflect"

	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	validate.RegisterValidation("isslice", IsSlice)
	validate.RegisterValidation("isstringelem", IsStringElem)
	validateVariable()
}

func validateVariable() {
	mySlice := []string{"111", "222"}
	//mySlice := []string{"111", "222", "", "", "", "", "", "", "", "", "", "", ""}
	//mySlice := []string{"111", "222000000000000", "0000000000000"}
	//1
	//if reflect.TypeOf(mySlice).Kind() != reflect.Slice {
	errs1 := validate.Var(mySlice, "isslice")
	if errs1 != nil {
		fmt.Println("Invalid in: Variable must be slice")
		fmt.Println(errs1)
		return
	}
	//2
	errs2 := validate.Var(mySlice, "max=10")
	if errs2 != nil {
		fmt.Println("Invalid in: Max len is 10. Original: ")
		fmt.Println(errs2)
		return
	}
	//3
	errs3 := validate.Var(mySlice, "required")
	if errs3 != nil {
		fmt.Println("Invalid in: Slice should not be null. Original: ")
		fmt.Println(errs3)
		return
	}
	//4
	errs4 := validate.Var(mySlice, "isstringelem")
	if errs4 != nil {
		fmt.Println("Invalid in: Element slice are string")
		fmt.Println(errs4)
		return
	}
	//5
	errs5 := validate.Var(mySlice, "dive,max=12") //applied in elements
	if errs5 != nil {
		fmt.Println("Invalid in: Max length of every element is 12. Original: ")
		fmt.Println(errs5)
		return
	}
}

//IsSlice check if field kind is equal to slice
func IsSlice(fl validator.FieldLevel) bool {

	if fl.Top().Kind() == reflect.Slice {
		return true
	}

	return false
}

//IsSlice check if field element kind is equal to string
func IsStringElem(fl validator.FieldLevel) bool {
	t := fl.Top().Type()
	if t.Elem().Kind() == reflect.String {
		return true
	}
	return false
}
