package main

import (

 "github.com/gogosphere/haaws"
 "fmt"
)

func main() {

	results := haaws.DescribeSubnets()

	for k, v := range results {
		fmt.Printf("['aws ")
		for _, vv := range v{
			fmt.Printf("%v ", vv)
		}
		fmt.Println("'],")
		

		_ = k
		_ = v
	}

}
