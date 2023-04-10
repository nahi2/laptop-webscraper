package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type laptop struct {
	Name            string `json:"name"`
	Price           string `json:"price"`
	OperatingSystem string `json:"operatingSystem"`
	CPU             string `json:"cpu"`
	GraphicsCard    string `json:"graphicsCard"`
	DisplaySize     string `json:"displaySize"`
}

func main() {
	domain := "https://www.dell.com/en-uk/shop/laptop-computers-2-in-1-pcs/sr/laptops?page"

	c := colly.NewCollector()

	laptops := make([]laptop, 0)

	c.OnHTML("section.ps-top", func(h *colly.HTMLElement) {
		laptop := laptop{
			Name:            h.ChildText(".ps-title"),
			Price:           h.ChildText("div.ps-dell-price.ps-simplified"),
			OperatingSystem: h.ChildText("div.short-specs.ps-dds-font-icon.dds_disc-system"),
			CPU:             h.ChildText("div.short-specs.ps-dds-font-icon.dds_processor"),
			GraphicsCard:    h.ChildText("div.short-specs.ps-dds-font-icon.dds_video-card"),
			DisplaySize:     h.ChildText("div.featured-spec.device-laptop"),
		}

		//fmt.Println(laptop)
		laptops = append(laptops, laptop)
	})

	c.OnHTML("#pagination > button:last-of-type", func(h *colly.HTMLElement) {
		if h.Attr("data-disabled") == "False" {
			nextPage := domain + "=" + h.Attr("data-to-page")
			c.Visit(nextPage)
		}
	})

	c.Visit(domain)

	content, err := json.Marshal(laptops)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("laptops.json", content, 0644)

}
