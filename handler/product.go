package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/tirthankarkundu17/ecommerce-price-checker/auth"
	api "github.com/tirthankarkundu17/ecommerce-price-checker/integrations/amazon"
	"github.com/tirthankarkundu17/ecommerce-price-checker/model"
	config "github.com/tirthankarkundu17/ecommerce-price-checker/utils"
)

func callAndUpdateChannel(amazon api.AmazonPriceChecker, url string, responses chan<- model.Product) {
	maxRetries := 3
	retry := 0
	product := model.Product{}
	for retry < maxRetries {
		product, _ = amazon.GetProduct(url)
		fmt.Println(product)
		if product.Name != "" {
			break
		}
		fmt.Println("Retrying")
		time.Sleep(time.Second * 5)
		retry += 1
	}

	responses <- product
}

func (h *Server) FetchProductsHandler(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	p := &model.UserProduct{}
	userProducts, _ := p.GetUserProduct(h.DB, userId)

	c := config.GetConf("integrations/amazon/config.yml")

	amazon := &api.AmazonPriceChecker{
		PriceSelector:  c.Price,
		NameSelector:   c.Name,
		ImagesSelector: c.Images,
		RatingSelector: c.Rating,
	}

	responses := make(chan model.Product)

	products := make([]model.Product, 0)
	for _, p := range *userProducts {
		go callAndUpdateChannel(*amazon, p.URL, responses)
	}

	// Merge the responses together
	for i := 0; i < len(*userProducts); i++ {
		response := <-responses
		products = append(products, response)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *Server) CreateUserProductHandler(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractTokenID(r)
	if err != nil {
		http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
		return
	}

	var p model.UserProduct

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.UserId = userId

	p.SaveUserProduct(h.DB)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
