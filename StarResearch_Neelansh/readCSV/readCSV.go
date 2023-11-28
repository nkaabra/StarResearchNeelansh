package ReadCSV

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// define a struct to hold seller product data
type SellerProduct struct {
	StockCode int
	Name      string
	Quantity  int
	Price     float32
}

func ReadCSVData(filename string) (map[string]int32, map[string]SellerProduct, error) {
	customerDemands := make(map[string]int32)
	sellerInventory := make(map[string]SellerProduct)

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read and discard the header row
	_, err = reader.Read()
	if err != nil {
		return nil, nil, err
	}
	// Read all the remaining records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	for _, record := range records {
		stockCode := record[0]
		nameOfProduct := record[1]
		quantityStr := record[2] // Quantity is in the third column
		priceStr := record[3]    // Price is in the fourth column

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			return nil, nil, err
		}
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, nil, err
		}

		//make the seller product struct
		sellerProduct := SellerProduct{

			Name:     nameOfProduct,
			Quantity: quantity,
			Price:    float32(price),
		}
		//cast quantity as int32 from int

		customerDemands[nameOfProduct] = int32(quantity)
		sellerInventory[stockCode] = sellerProduct
	}
	fmt.Println("CSV file read successfully.")
	return customerDemands, sellerInventory, nil
}
