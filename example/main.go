package main

import (
	"context"
	"fmt"

	"github.com/masv3971/gonavet2"
)

func main() {
	navet, err := gonavet2.New(&gonavet2.Config{
		BaseURL:        "",
		ProxyURL:       "",
		CertificatePEM: []byte{},
		PrivateKeyPEM:  []byte{},
		OrgNumber:      "",
		OrderID:        "",
		Token: gonavet2.OauthConfig{
			Scope:        "",
			BaseURL:      "",
			ClientID:     "",
			ClientSecret: "",
			GrantType:    "",
		},
		APIGWClientID:     "",
		APIGWClientSecret: "",
	})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	person, req, err := navet.Search.Name(ctx, &gonavet2.RequestSearchName{
		GivenName: "testName",
		Surname:   "testSurname",
		City:      "testCity",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("person", person, "req", req)
}
