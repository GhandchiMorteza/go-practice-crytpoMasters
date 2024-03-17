package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"test.com/crypto/data"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*data.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := data.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}