package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

const API = "https://api.pokemontcg.io/v1/cards"

var (
	ErrorBackend = errors.New("Something went wrong")
)

type Request struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Types    string `json:"types"`
	Name     string `json:"name"`
}

func (r *Request) QueryParams() url.Values {
	q := url.Values{}
	if r.Page > 0 {
		q.Add("page", strconv.Itoa(r.Page))
	}

	if r.PageSize > 0 {
		q.Add("pageSize", strconv.Itoa(r.PageSize))
	} else {
		q.Add("pageSize", strconv.Itoa(10))
	}

	if r.Types != "" {
		q.Add("types", r.Types)
	}

	if r.Name != "" {
		q.Add("name", r.Name)
	}

	return q
}

type CardsResponse struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	NationalPokedexNumber int      `json:"nationalPokedexNumber"`
	ImageURL              string   `json:"imageUrl"`
	ImageURLHiRes         string   `json:"imageUrlHiRes"`
	Types                 []string `json:"types"`
	Supertype             string   `json:"supertype"`
	Subtype               string   `json:"subtype"`
	Hp                    string   `json:"hp"`
	RetreatCost           []string `json:"retreatCost"`
	ConvertedRetreatCost  int      `json:"convertedRetreatCost"`
	Number                string   `json:"number"`
	Artist                string   `json:"artist"`
	Rarity                string   `json:"rarity"`
	Series                string   `json:"series"`
	Set                   string   `json:"set"`
	SetCode               string   `json:"setCode"`
	Attacks               []struct {
		Cost                []string `json:"cost"`
		Name                string   `json:"name"`
		Text                string   `json:"text"`
		Damage              string   `json:"damage"`
		ConvertedEnergyCost int      `json:"convertedEnergyCost"`
	} `json:"attacks"`
	Resistances []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"resistances"`
	Weaknesses []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"weaknesses"`
}

func Handler(ctx context.Context, request Request) ([]Card, error) {
	url := fmt.Sprintf(API)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []Card{}, ErrorBackend
	}

	q := request.QueryParams()
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return []Card{}, ErrorBackend
	}
	defer resp.Body.Close()

	var data CardsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Card{}, ErrorBackend
	}

	return data.Cards, nil
}

func main() {
	lambda.Start(Handler)
}
