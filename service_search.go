package gonavet2

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/masv3971/gonavet2/navettypes"
)

type searchService struct {
	client   *Client
	endpoint string
}

func (s *searchService) requestBody(navettypes.Sokvillkor) (navettypes.SokRequest, error) {
	req := navettypes.SokRequest{
		Bestallning: navettypes.Bestallning{
			BestallningsIdentitet: s.client.orderID,
			Organisationsnummer:   s.client.orgNumber,
		},
		Sokvillkor: navettypes.Sokvillkor{},
	}

	if err := validate(req); err != nil {
		return navettypes.SokRequest{}, err
	}
	return req, nil
}

func (s *searchService) post(ctx context.Context, sokvillkor navettypes.Sokvillkor, data interface{}) (*navettypes.OkSokResponse, *resty.Response, error) {
	if data != nil {
		if err := validate(data); err != nil {
			return nil, nil, err
		}
	}

	body, err := s.requestBody(sokvillkor)
	if err != nil {
		return nil, nil, err
	}

	result := &navettypes.OkSokResponse{}
	req, err := s.client.req(ctx, body, result)
	if err != nil {
		return nil, nil, err
	}

	resp, err := req.Post(s.endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type RequestSearchName struct {
	GivenName string `json:"given_name" validate:"required_without_all=Surename City"`
	Surname   string `json:"surename" validate:"required_without_all=GivenName City"`
	City      string `json:"city" validate:"required_without_all=GivenName Surename"`
}

func (s *searchService) Name(ctx context.Context, data *RequestSearchName) (*navettypes.OkSokResponse, *resty.Response, error) {
	sokvillkor := navettypes.Sokvillkor{
		Adress: navettypes.SokvillkorAdress{
			Gatuadress: navettypes.SokvillkorGatuadress{},
			Postnummer: navettypes.SokvillkorPostnummer{},
			Postort:    data.City,
		},
		Fodelsetid: navettypes.SokvillkorFodelsetid{},
		Namn: navettypes.SokvillkorNamn{
			Fornamn: navettypes.SokvillkorFornamn{
				Namndelar: []string{data.GivenName},
			},
			MellanEllerEfternamn: navettypes.SokvillkorMellanEllerEfternamn{
				Namndelar: []string{data.Surname},
			},
		},
	}

	return s.post(ctx, sokvillkor, data)
}

// RequestSearchFodelsetid input type for search date of birth
type RequestSearchFodelsetid struct {
	From navettypes.SokvillkorFodelsetidDatum `json:"from" validate:"required"`
	To   navettypes.SokvillkorFodelsetidDatum `json:"to" validate:"required"`
}

func (s *searchService) Fodelsetid(ctx context.Context, data *RequestSearchFodelsetid) (*navettypes.OkSokResponse, *resty.Response, error) {
	sokvillkor := navettypes.Sokvillkor{
		Adress: navettypes.SokvillkorAdress{},
		Fodelsetid: navettypes.SokvillkorFodelsetid{
			Fran: data.From,
			Till: data.To,
		},
	}

	return s.post(ctx, sokvillkor, data)
}

type RequestSearchPostal struct {
	StreetAddress string `json:"street_address"`
	PostalNumber  []int  `json:"postal_number" validate:"omitempty,eq=2"`
	City          string `json:"city"`
}

// Postal search person with postal address attributes
func (s *searchService) Postal(ctx context.Context, data *RequestSearchPostal) (*navettypes.OkSokResponse, *resty.Response, error) {
	sokvillkor := navettypes.Sokvillkor{
		Adress: navettypes.SokvillkorAdress{
			Gatuadress: navettypes.SokvillkorGatuadress{
				SammansattVarde: data.StreetAddress,
			},
			Postnummer: navettypes.SokvillkorPostnummer{
				Fran: data.PostalNumber[0],
				Till: data.PostalNumber[1],
			},
			Postort: data.City,
		},
	}

	return s.post(ctx, sokvillkor, data)

}

// All is able to use all types in navettypes.Sokvillkor, this should not be necessary but included just in case.
// Do not hesitate to include a method for a special search pattern if not found here, in the name of cleaner and more readable code.
func (s *searchService) All(ctx context.Context, data navettypes.Sokvillkor) (*navettypes.OkSokResponse, *resty.Response, error) {
	if err := validate(data); err != nil {
		return nil, nil, err
	}

	return s.post(ctx, data, data)
}
