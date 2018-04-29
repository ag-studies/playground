package main

import (
	"fmt"
	"mydata"
	//  "strconv"
)

type module string

//func (m module) RunMod(i int) {
func (m module) RunMod(in <-chan *mydata.DataUnit) {
	//fmt.Println("Hello Universe " + strconv.Itoa(i))
	fmt.Println("Hello Universe ")
	n := <-in
	fmt.Printf("%v", n.S)
}

var VarModule module
