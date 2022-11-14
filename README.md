### Bintable API Golang


Just contribution to Bintable project with golang API library.


```go

import ""github.com/lexterl33t/bintable_api_go"
```

You must create bintable account to this link <a href="https://bintable.com/get-api">Bintable</a>


Set API KEY

```go
New() *BintableConfig

```

```go

type LookupResponse_t struct {
	Result  int                  `json:"result"`
	Message string               `json:"message"`
	Data    LookupDataResponse_t `json:"data"`
}

type LookupDataResponse_t struct {
	Card    LookupDataCardResponse_t    `json:"card"`
	Country LookupDataCountryResponse_t `json:"country"`
	Bank    LookupDataBankResponse_t    `json:"bank"`
}

type LookupDataCardResponse_t struct {
	Scheme    string `json:"scheme"`
	Type      string `json:"type"`
	Category  string `json:"category"`
	Length    int    `json:"length"`
	CheckLuhn int    `json:"checluhn"`
	CVVlen    int    `json:"cvvlen"`
}

type LookupDataCountryResponse_t struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Flag         string `json:"flag"`
	Currency     string `json:"currency"`
	CurrencyCode string `json:"currency_code"`
}

type LookupDataBankResponse_t struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type Bintable interface {
	New(key string) *BintableConfig
}

type BintableConfig struct {
	Key      string
	Endpoint string
}
```


```go

bin := bintable.New("API_KEY")

```

```go
res, err := bin.Lookup("403244")
if err != nil {
  log.Fatalln(err)
}

fmt.Println(res)

```

