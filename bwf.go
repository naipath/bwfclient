package bwfclient

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	url = "https://www.rabobank.nl/diensten/hypotheken/homecatcher/taxatie/basisWoningfinanciering"
)

// BwfResponse is the response retrieved from Bwf
type BwfResponse struct {
	MaxTeLenenObvInkomen struct {
		Toetsrente struct {
			MaxTeLenenInkomenBedr int `json:"maxTeLenenInkomenBedr"`
			KoopsomBedr           int `json:"koopsomBedr"`
			KostenBedr            int `json:"kostenBedr"`
		} `json:"toetsrente"`
		Tienjaarsrente struct {
			MaxTeLenenInkomenBedr int `json:"maxTeLenenInkomenBedr"`
			KoopsomBedr           int `json:"koopsomBedr"`
			KostenBedr            int `json:"kostenBedr"`
		} `json:"tienjaarsrente"`
	} `json:"maxTeLenenObvInkomen"`
	WoonLasten struct {
		Toetsrente struct {
			BrutoLastMbedr int `json:"brutoLastMbedr"`
			NettoLastMbedr int `json:"nettoLastMbedr"`
		} `json:"toetsrente"`
		Tienjaarsrente struct {
			BrutoLastMbedr int `json:"brutoLastMbedr"`
			NettoLastMbedr int `json:"nettoLastMbedr"`
		} `json:"tienjaarsrente"`
	} `json:"woonLasten"`
	FinancieringObvNettoLast struct {
		Toetsrente struct {
			FinancieringBedr int `json:"financieringBedr"`
			KoopsomBedr      int `json:"koopsomBedr"`
			KostenBedr       int `json:"kostenBedr"`
		} `json:"toetsrente"`
		Tienjaarsrente struct {
			FinancieringBedr int `json:"financieringBedr"`
			KoopsomBedr      int `json:"koopsomBedr"`
			KostenBedr       int `json:"kostenBedr"`
		} `json:"tienjaarsrente"`
	} `json:"financieringObvNettoLast"`
	Toetsrente     float64 `json:"toetsrente"`
	Tienjaarsrente float64 `json:"tienjaarsrente"`
}

// BwfClient to represent all actions possible with bwf
type BwfClient struct {
}

// New creates a BwfClient
func New() *BwfClient {
	return &BwfClient{}
}

// Request retrieves the bwf results
func (b BwfClient) Request(aanvragerBrutoJaarinkomenBedr int, partnerBrutoJaarinkomenBedr int) (BwfResponse, error) {

	resp, err := http.Get(createRequest(aanvragerBrutoJaarinkomenBedr, partnerBrutoJaarinkomenBedr))

	if err != nil {
		return BwfResponse{}, err
	}

	defer resp.Body.Close()

	var data BwfResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return BwfResponse{}, err
	}
	return data, nil
}

func createRequest(aanvragerBrutoJaarinkomenBedr int, partnerBrutoJaarinkomenBedr int) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}

	q := req.URL.Query()
	q.Add("aanvragerBrutoJaarinkomenBedr", strconv.Itoa(aanvragerBrutoJaarinkomenBedr))
	q.Add("partnerBrutoJaarinkomenBedr", strconv.Itoa(partnerBrutoJaarinkomenBedr))
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
