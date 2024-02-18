package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	tickers := []string{
		"MSFT",
		"GOOGL",
		"AMZN",
		"FB",
		"INTC",
		"NVDA",
		"ADBE",
		"PYPL",
		"CSCO",
		"NFLX",
		"AVGO",
		"TXN",
		"QCOM",
		"CMCSA",
		"ASML",
		"PEP",
		"TMUS",
		"AMAT",
		"SBUX",
		"ISRG",
		"AMD",
		"MU",
		"ADP",
		"INTU",
		"ATVI",
		"CSX",
		"ADI",
	}

	var stocks []Stock

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.company = e.ChildText("h1")
		fmt.Println("Company: ", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price: ", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change: ", stock.change)

		stocks = append(stocks, stock)
	})

	for _, t := range tickers {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}
	fmt.Println(stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	headers := []string{"Company", "Price", "Change"}
	writer.Write(headers)
	for _, stock := range stocks {
		record := []string{stock.company, stock.price, stock.change}
		writer.Write(record)
	}

	defer writer.Flush()
}
