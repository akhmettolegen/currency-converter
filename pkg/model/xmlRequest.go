package model

import "encoding/xml"

type Item struct {
	Title 		Currency `xml:"title"`
	Description float64  `xml:"description"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Item	[]Item 	 `xml:"item"`
}

type XMLRates struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}
