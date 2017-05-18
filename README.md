# bwfclient

bwfclient is a golang library that wraps the bwf interface

Example request: 
```go
var bwfClient = bwfclient.New()
resp, err := bwfClient.Request(
    BwfRequest{
		AanvragerBrutoJaarinkomenBedr: 0,
		PartnerBrutoJaarinkomenBedr:   0,
	}
)

fmt.Print(resp.MaxTeLenenObvInkomen.Tienjaarsrente.KoopsomBedr)
```