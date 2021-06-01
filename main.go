package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bad33ndj3/costs/pkg/sharedexpenses"
)

func main() {
	arg1 := flag.Int("person1", 1000, "Net monthly income of person 1")
	arg2 := flag.Int("person2", 1000, "Net monthly income of person 2")
	arg3 := flag.Int("expenses", 1000, "Total shared monthly expenses")
	flag.Parse()

	cost := &sharedexpenses.Costs{}
	err := cost.Init(arg1, arg2, arg3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TO RATIO")
	toRatio := &sharedexpenses.ToRatio{}
	toRatio.Fill(*cost)
	toRatio.Render()

	fmt.Println("50/50")
	splitEqual := &sharedexpenses.SplitEqual{}
	splitEqual.Fill(*cost)
	splitEqual.Render()
}
