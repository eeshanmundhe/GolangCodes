package main

import "fmt"

func main() {
	jewel := make(map[string]float64)
	jewel["necklace"] = 57.7
	jewel["earrings"] = 100.8
	my, ok := jewel["earrings"]
	fmt.Println(my, ok)
	cloth := map[string]float64{"pants": 70.6, "shirt": 40.9}
	{
		fmt.Println(jewel, cloth)
	}

}
