package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	orders := extract()
	orders = transform(orders)
	load(orders)
	fmt.Println(time.Since(start))
}

// Product holds details of a product
type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

// Order holds details of a customer order
type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int

	UnitCost  float64
	UnitPrice float64
}

func extract() []*Order {
	result := []*Order{}

	f, _ := os.Open("./orders.txt")
	defer f.Close()

	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		o := new(Order)
		o.CustomerNumber, _ = strconv.Atoi(record[0])
		o.PartNumber = record[1]
		o.Quantity, _ = strconv.Atoi(record[2])
		result = append(result, o)
	}

	return result
}

func transform(orders []*Order) []*Order {
	f, _ := os.Open("./productList.txt")
	defer f.Close()

	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)

	for _, r := range records {
		p := new(Product)
		p.PartNumber = r[0]
		p.UnitCost, _ = strconv.ParseFloat(r[1], 64)
		p.UnitPrice, _ = strconv.ParseFloat(r[2], 64)
		productList[p.PartNumber] = p
	}

	for idx := range orders {
		time.Sleep(3 * time.Millisecond)
		o := orders[idx]
		o.UnitCost = productList[o.PartNumber].UnitCost
		o.UnitPrice = productList[o.PartNumber].UnitPrice
	}

	return orders

}

func load(orders []*Order) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%16s%13s%13s%16s%16s", "Part Number", "Quantity",
		"Unit Cost", "Unit Price", "Total Cost", "Total Price\n")

	for _, o := range orders {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f%15.2f\n",
			o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
			o.UnitCost*float64(o.Quantity), o.UnitPrice*float64(o.UnitPrice))
	}
}
