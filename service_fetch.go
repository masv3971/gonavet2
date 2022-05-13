package gonavet2

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/masv3971/gonavet2/navettypes"
)

// FetchService holds Fetch object
type fetchService struct {
	client   *Client
	endpoint string
}

func (s *fetchService) requestBody(sokvillkor []navettypes.Sokvillkorsidentitet) (navettypes.HamtaRequest, error) {
	hamtaReq := navettypes.HamtaRequest{
		Bestallning: navettypes.Bestallning{
			BestallningsIdentitet: s.client.orderID,
			Organisationsnummer:   s.client.orgNumber,
		},
		Sokvillkor: sokvillkor,
	}

	if err := validate(hamtaReq); err != nil {
		return navettypes.HamtaRequest{}, err
	}
	return hamtaReq, nil
}

func (s *fetchService) post(ctx context.Context, sokvillkor []navettypes.Sokvillkorsidentitet, data interface{}) (*navettypes.OkHamtaResponse, *resty.Response, error) {
	if data != nil {
		if err := validate(data); err != nil {
			return nil, nil, err
		}
	}

	body, err := s.requestBody(sokvillkor)
	if err != nil {
		return nil, nil, err
	}

	result := &navettypes.OkHamtaResponse{}
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

type RequestFetchNIN struct {
	// NINs can consists of personnummer and/or samordningsnummer
	NINs []string `json:"nin" validate:"request, min=1"`
}

func (s *fetchService) NINs(ctx context.Context, data *RequestFetchNIN) (*navettypes.OkHamtaResponse, *resty.Response, error) {
	sokvillkor := []navettypes.Sokvillkorsidentitet{}

	for _, nin := range data.NINs {
		sokvillkor = append(sokvillkor, navettypes.Sokvillkorsidentitet{
			Identitetsbeteckning: nin,
		})
	}

	return s.post(ctx, sokvillkor, data)
}
