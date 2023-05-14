# Laptop Scraper

This is a simple web scraper written in Go that extracts information about laptops from the Dell UK website and saves it in a JSON file.

## Requirements

- Go 1.16 or higher
- `github.com/gocolly/colly` package

## Installation

1. Install Go on your system by following the instructions on the [official Go website](https://golang.org/doc/install).
2. Install the `gocolly/colly` package by running the following command in your terminal:

```sh
go get -u github.com/gocolly/colly
```

## Usage

1. Clone or download the source code from this repository.
2. Navigate to the directory where the `main.go` file is located.
3. Run the following command to build the executable file:

```sh
go build
```

4. Run the executable file using the following command:

```sh
./laptop-scraper
```

The scraper will start fetching data from the Dell UK website and save it in a file named `laptops.json` in the same directory as the executable file.
