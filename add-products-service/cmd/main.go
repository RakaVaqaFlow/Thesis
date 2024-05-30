package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Photo       string  `json:"photo"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func createProduct(product Product) error {
	jsonData, err := json.Marshal(product)
	if err != nil {
		return err
	}

	url := os.Getenv("APS_TARGET_URL")
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create product: %s", response.Status)
	}

	return nil
}

func main() {

	num, err := strconv.Atoi(os.Getenv("NUM_OF_PRODUCTS"))
	if err != nil {
		log.Fatal("Num of products should be integer")
	}
	created := 0
	for i := 0; i < num; i++ {
		product := Product{
			ID:          i,
			Name:        fmt.Sprintf("Product %d", i),
			Description: fmt.Sprintf("Description for Product %d", i),
			Photo:       fmt.Sprintf("url_to_photo%d.jpg", i),
			Price:       10.0 + float64(i)*1.5,
			Quantity:    100 - (i % 10),
		}
		err := createProduct(product)
		if err != nil {
			log.Printf("Error creating product %d: %v\n", product.ID, err)
			continue
		} else {
			created++
		}
	}
	log.Printf("Successfuly created %d products", created)
}
