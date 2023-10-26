package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Code           string
	Name           string
	PriceCurrent   float64
	QuantityCurrent int
}

type Category struct {
	ID           string
	Name         string
	ProductList  []Product
}

func main() {
	// Open the CSV file for reading
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	categories := make([]Category, 0)

	// Iterate through each record, skipping the header
	for i, record := range records {
		if i == 0 {
			continue // Skip the header
		}

		id := record[0]
		name := record[1]
		productsStr := strings.Trim(record[2], "[]")
		productData := strings.Split(productsStr, "], [")
		products := make([]Product, 0)

		for _, pData := range productData {
			parts := strings.Split(pData, ",")
			if len(parts) != 4 {
				log.Printf("Error parsing product data: %s\n", pData)
				continue
			}

			code := parts[0]
			productName := parts[1]
			price, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				log.Printf("Error parsing price for product: %s\n", productName)
				continue
			}
			quantity, err := strconv.Atoi(parts[3])
			if err != nil {
				log.Printf("Error parsing quantity for product: %s\n", productName)
				continue
			}

			product := Product{
				Code:           code,
				Name:           productName,
				PriceCurrent:   price,
				QuantityCurrent: quantity,
			}

			products = append(products, product)
		}

		category := Category{
			ID:          id,
			Name:        name,
			ProductList: products,
		}

		categories = append(categories, category)
	}

	// Print the extracted data
	for _, category := range categories {
		fmt.Printf("Category: %s\n", category.Name)
		for _, product := range category.ProductList {
			fmt.Printf("Product: %s, Code: %s, Price: %.2f, Quantity: %d\n", product.Name, product.Code, product.PriceCurrent, product.QuantityCurrent)
		}
	}
}
