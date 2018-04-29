package main

import (
	"fmt"
	"mydata"
	"os"
	"plugin"
)

type Module interface {
	//RunMod(i int)
	RunMod(in <-chan *mydata.DataUnit)
}

func main() {

	out := make(chan *mydata.DataUnit, 2000)
	plug, err := plugin.Open("./modcounter.so")
	if err != nil {
		fmt.Printf("FATAL (plugin.Open): " + err.Error())
		os.Exit(1)
	}

	symModule, err := plug.Lookup("VarModule")
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	//var mod module
	var mod, ok = symModule.(Module)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	// mod.RunMod(5)
	mod.RunMod(out)
}
