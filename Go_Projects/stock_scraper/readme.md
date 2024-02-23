# Stock Price Scraper

## Overview
This Go program scrapes stock prices and changes from Yahoo Finance for a list of specified ticker symbols. It then saves the data into a CSV file named `stocks.csv`.

## Features
- Scrapes stock prices and changes from Yahoo Finance.
- Saves the data into a CSV file for further analysis.

## Dependencies
- [github.com/gocolly/colly](https://github.com/gocolly/colly): Used for web scraping functionalities.
- Standard Go libraries for file handling and CSV operations.

## Installation
1. Clone or download the repository.
2. Ensure you have Go installed on your system.
3. Navigate to the project directory.
4. Run `go mod tidy` to install the necessary dependencies.

## Usage
1. Modify the `tickers` array in the `main` function to include the ticker symbols you want to scrape.
2. Run the program by executing `go run main.go`.
3. The program will scrape the stock prices and changes for the specified ticker symbols and save the data into `stocks.csv`.

## How It Works
- The program utilizes the Colly library to perform web scraping on Yahoo Finance pages.
- It extracts the company name, current price, and change percentage from the HTML of the Yahoo Finance page.
- Data for each stock is stored in a struct and then appended to a slice of structs.
- Finally, the program writes the collected data into a CSV file.

## Limitations
- The program relies on the HTML structure of Yahoo Finance pages. Any changes to the structure may break the scraping functionality.
- It only supports scraping data from Yahoo Finance and may not work for other financial websites without modifications.


