package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

// small first letter means private
// big first letter means public
type BoardGameAtlas struct {
	//"member"
	clientId string
}

// Go Lang can map objcet to struct variables
type Game struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	YearPublished uint   `json:"year_published"`
	Description   string `json:"description"`
	Url           string `json:"official_url"`
	ImageUrl      string `json:"image_url"`
	RulesUrl      string `json:"rules_url"`
}

type SearchResults struct {
	Games []Game `json:"games"`
	Count uint   `json:"count"`
}

// Functions as a constructor
func New(clientId string) BoardGameAtlas {

	return BoardGameAtlas{clientId: clientId}
}

// A receiver, function name must be Upper case, to declare public. Lower case to declare private.
// Method in BoardGameAtlas
func (b BoardGameAtlas) Search(ctx context.Context, query string, limit uint, skip uint) (*SearchResults, error) {
	//Create HTTPClient.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)

	//Check if there is any error
	if nil != err {
		// returns an error object
		return nil, fmt.Errorf("Unable to create HTTP Client: %v", err)
	}

	//Get the query string object
	qs := req.URL.Query()

	//populate URL with query params
	qs.Add("name", query)
	qs.Add("limit", fmt.Sprintf("%d", limit))
	qs.Add("skip", strconv.Itoa(int(skip)))
	qs.Add("client_id", b.clientId)

	//Encode query params, add it back to the request
	req.URL.RawQuery = qs.Encode()

	//fmt.Printf("URL = %s,\n", req.URL.String())
	//make the call
	resp, err := http.DefaultClient.Do(req)

	if nil != err {
		return nil, fmt.Errorf("Unable to call HTTP Error Message: %v", err)
	}

	if resp.StatusCode >= 400 {

		return nil, fmt.Errorf("Unable to call HTTP Error Message: %v", resp.Status)

	}

	var results SearchResults
	//Deserialize the json
	if err := json.NewDecoder(resp.Body).Decode(&results); nil != err {

		return nil, fmt.Errorf("Cannot deserialize JSON payload: %v", err)
	}

	return &results, nil

}
