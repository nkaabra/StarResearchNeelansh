// Init/initialize.go

package Init

import "fmt"

func Initialize() (map[string]int32, map[string]float32) {

	//make a map of customer demands with product name as the key and the quantity as the value
	customerDemands := make(map[string]int32)
	//make a map of seller inventory with product name as the key and the cost as the value
	sellerInventory := make(map[string]float32)

	fmt.Println("Initialization completed.")
	return customerDemands, sellerInventory
}
