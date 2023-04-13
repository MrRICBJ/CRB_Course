package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"io/ioutil"
	"log"
	"time"
)

const dateFormat = "02/01/2006"
const baseURL = "http://www.cbr.ru/scripts/XML_daily.asp"

// Currency is a currency item
type Currency struct {
	ID       string `xml:"ID,attr"`
	NumCode  uint   `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nom      uint   `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

// Result is a result representation
type Result struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Currencies []Currency `xml:"Valute"`
}

func getCurrencyRate(t time.Time, fetch FetchFunction) (Result, error) {
	log.Printf("Fetching the currency rate at %v\n", t.Format("02.01.2006"))
	var result Result
	err := getCurrencies(&result, t, fetch)
	if err != nil {
		return Result{}, err
	}
	return result, nil
}

func getCurrencies(v *Result, t time.Time, fetch FetchFunction) error {
	url := baseURL + "?date_req=" + t.Format(dateFormat)
	resp, err := fetch(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Request error: StatusCode: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("Unknown charset: %s", charset)
		}
	}
	err = decoder.Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
