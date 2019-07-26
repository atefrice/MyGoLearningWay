package main

import (
	"fmt"
	"flag"
)

var RoutingCnt int;

func init() {
	flag.IntVar(&RoutingCnt, "cnt", 20, "Create go routing count.(def: 20)");
}


func main() { 

flag.Parse()
fmt.Println(RoutingCnt)

}
