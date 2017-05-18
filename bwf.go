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

// BwfRequest needed for the request
type BwfRequest struct {
	AanvragerBrutoJaarinkomenBedr int
	PartnerBrutoJaarinkomenBedr   int
	KoopsomBedr                   int
	NettoLastMbedr                int
}

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
func (b BwfClient) Request(request BwfRequest) (BwfResponse, error) {
	resp, err := http.Get(createRequest(request))

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

func createRequest(request BwfRequest) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
	}

	q := req.URL.Query()
	q.Add("aanvragerBrutoJaarinkomenBedr", strconv.Itoa(request.AanvragerBrutoJaarinkomenBedr))
	q.Add("partnerBrutoJaarinkomenBedr", strconv.Itoa(request.PartnerBrutoJaarinkomenBedr))
	q.Add("nettoLastMbedr", strconv.Itoa(request.NettoLastMbedr))
	q.Add("koopsomBedr", strconv.Itoa(request.KoopsomBedr))

	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
