package main

import "fmt"

func main() {

	//Choose region
	regions := getRegions()
	regions.showMenu()
	var selectedRegionIndex int
	fmt.Println("Choose a Region: ")
	_, err := fmt.Scanf("%d", &selectedRegionIndex)
	if err != nil {
		fmt.Println(err)
	}

	//Choose province
	provinces := getProvinces(regions[selectedRegionIndex])
	provinces.showMenu()
	var selectedProvinceIndex int
	fmt.Println("Choose a Province: ")
	_, err = fmt.Scanf("%d", &selectedProvinceIndex)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(provinces[selectedProvinceIndex])

}
