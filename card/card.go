package card

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// CardInfo contains the card information.
type CardInfo struct {
	Number  number  `json:"number"`
	Scheme  string  `json:"scheme"`
	Type    string  `json:"type"`
	Brand   string  `json:"brand"`
	Prepaid bool    `json:"prepaid"`
	Country country `json:"country"`
	Bank    bank    `json:"bank"`
}

type number struct {
	Length int  `json:"length"`
	Luhn   bool `json:"luhn"`
}

type country struct {
	Numeric   string `json:"numeric"`
	Alpha2    string `json:"alpha2"`
	Name      string `json:"name"`
	Emoji     string `json:"emoji"`
	Currency  string `json:"currency"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}

type bank struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

// GetInfo gets the information of a card from the bin number.
func GetInfo(bin string) (*CardInfo, error) {
	if !isValidCardNumber(bin) {
		return nil, fmt.Errorf("Card number not valid")
	}

	url := fmt.Sprintf("https://lookup.binlist.net/%s", bin)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Version", "3")

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info CardInfo

	err = json.Unmarshal(body, &info)

	return &info, err
}

// isValidCardNumber returns true if the card bin is valid.
func isValidCardNumber(bin string) bool {
	re := regexp.MustCompile(`^\d{6,19}$`)
	return re.MatchString(bin)
}
