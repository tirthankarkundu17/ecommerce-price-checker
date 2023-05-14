package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tirthankarkundu17/ecommerce-price-checker/model"
	"github.com/tirthankarkundu17/ecommerce-price-checker/utils"
)

type AmazonPriceChecker struct {
	PriceSelector  utils.Selector
	NameSelector   utils.Selector
	ImagesSelector utils.Selector
	RatingSelector utils.Selector
}

func (c *AmazonPriceChecker) GetProduct(productURL string) (model.Product, error) {

	response := &model.Product{}

	productURL = utils.RemoveQueryParamsFromUrl(productURL)

	// Send GET request to product page URL

	client := &http.Client{}
	req, _ := http.NewRequest("GET", productURL, nil)
	headers := http.Header{}
	headers.Set("dnt", "1")
	headers.Set("upgrade-insecure-requests", "1")
	headers.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	headers.Set("sec-fetch-site", "same-origin")
	headers.Set("sec-fetch-mode", "navigate")
	headers.Set("sec-fetch-user", "?1")
	headers.Set("sec-fetch-dest", "document")

	referrer, _ := utils.GetDomainHost(productURL)
	headers.Set("referer", referrer)
	headers.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")

	// Set the headers in the request
	req.Header = headers

	urlResponse, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return *response, err
	}
	defer urlResponse.Body.Close()

	// Validate the urlResponse
	if urlResponse.StatusCode != 200 {
		return *response, err
	}

	// Create a new document from the response body
	doc, err := goquery.NewDocumentFromReader(urlResponse.Body)
	if err != nil {
		log.Fatal(err)
		return *response, err
	}

	if err != nil {
		log.Fatal(err)
		return *response, err
	}

	price := doc.Find(c.PriceSelector.Selector).First().Text()
	name := doc.Find(c.NameSelector.Selector).Text()
	images := make([]string, 0)
	i := doc.Find(c.ImagesSelector.Selector)
	i.Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr(c.ImagesSelector.Attribute)
		images = append(images, strings.TrimSpace(link))
	})

	rating := strings.Fields((strings.Split(doc.Find(c.RatingSelector.Selector).Text(), "out")[0]))[0]

	response = &model.Product{
		Name:   strings.TrimSpace(name),
		Price:  strings.TrimSpace(price),
		Rating: strings.TrimSpace(rating),
		Images: images,
		URL:    productURL,
	}

	return *response, nil
}
