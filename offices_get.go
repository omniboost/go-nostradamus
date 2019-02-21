package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/gocarina/gocsv"
)

func (c *Client) NewOfficesGetRequest() OfficesGetRequest {
	return OfficesGetRequest{
		client:      c,
		queryParams: c.NewOfficesGetQueryParams(),
		pathParams:  c.NewOfficesGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewOfficesGetRequestBody(),
	}
}

type OfficesGetRequest struct {
	client      *Client
	queryParams *OfficesGetQueryParams
	pathParams  *OfficesGetPathParams
	method      string
	headers     http.Header
	requestBody OfficesGetRequestBody
}

func (c *Client) NewOfficesGetQueryParams() *OfficesGetQueryParams {
	return &OfficesGetQueryParams{}
}

type OfficesGetQueryParams struct {
}

func (p OfficesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OfficesGetRequest) QueryParams() *OfficesGetQueryParams {
	return r.queryParams
}

func (c *Client) NewOfficesGetPathParams() *OfficesGetPathParams {
	return &OfficesGetPathParams{}
}

type OfficesGetPathParams struct {
}

func (p *OfficesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OfficesGetRequest) PathParams() *OfficesGetPathParams {
	return r.pathParams
}

func (r *OfficesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *OfficesGetRequest) Method() string {
	return r.method
}

func (s *Client) NewOfficesGetRequestBody() OfficesGetRequestBody {
	return OfficesGetRequestBody{}
}

type OfficesGetRequestBody struct {
}

func (r *OfficesGetRequest) RequestBody() *OfficesGetRequestBody {
	return &r.requestBody
}

func (r *OfficesGetRequest) SetRequestBody(body OfficesGetRequestBody) {
	r.requestBody = body
}

func (r *OfficesGetRequest) NewResponseBody() *OfficesGetResponseBody {
	return &OfficesGetResponseBody{}
}

type OfficesGetResponseBody struct {
	Offices
}

func (r *OfficesGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("index2.php?option=com_webservices&controller=csv&method=hours.board.fetch&key={{.key}}&element=offices", r.PathParams())
}

func (r *OfficesGetRequest) Do() (OfficesGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)

	err = gocsv.Unmarshal(resp.Body, &responseBody.Offices)
	return *responseBody, err
}
