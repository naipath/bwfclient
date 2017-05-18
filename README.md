# bwfclient

bwfclient is a golang library that wraps the bwf interface

Example request: 
```go
var bwfClient = bwfclient.New()
resp, err := bwfClient.Request(40000, 0)

fmt.Print(resp.MaxTeLenenObvInkomen.Tienjaarsrente.KoopsomBedr)
```