package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
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

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := fmt.Sprintf(API)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: ErrorBackend.Error(), StatusCode: 400}, nil
	}

	page, _ := strconv.Atoi(request.QueryStringParameters["page"])
	pageSize, _ := strconv.Atoi(request.QueryStringParameters["pageSize"])
	//Get the path parameter that was sent
	q := Request{
		Name:     request.QueryStringParameters["name"],
		Types:    request.QueryStringParameters["types"],
		Page:     page,
		PageSize: pageSize,
	}
	req.URL.RawQuery = q.QueryParams().Encode()

	resp, err := client.Do(req)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: ErrorBackend.Error(), StatusCode: 400}, nil
	}
	defer resp.Body.Close()

	var data CardsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return events.APIGatewayProxyResponse{Body: ErrorBackend.Error(), StatusCode: 400}, nil
	}

	response, err := json.Marshal(&data.Cards)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 404}, nil
	}

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",    // (* or a specific host)
			"Access-Control-Allow-Credentials": "true", // Required for cookies, authorization headers with HTTPS
		},
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
