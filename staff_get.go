package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/gocarina/gocsv"
)

func (c *Client) NewStaffGetRequest() StaffGetRequest {
	return StaffGetRequest{
		client:      c,
		queryParams: c.NewStaffGetQueryParams(),
		pathParams:  c.NewStaffGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewStaffGetRequestBody(),
	}
}

type StaffGetRequest struct {
	client      *Client
	queryParams *StaffGetQueryParams
	pathParams  *StaffGetPathParams
	method      string
	headers     http.Header
	requestBody StaffGetRequestBody
}

func (c *Client) NewStaffGetQueryParams() *StaffGetQueryParams {
	return &StaffGetQueryParams{}
}

type StaffGetQueryParams struct {
}

func (p StaffGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *StaffGetRequest) QueryParams() *StaffGetQueryParams {
	return r.queryParams
}

func (c *Client) NewStaffGetPathParams() *StaffGetPathParams {
	return &StaffGetPathParams{}
}

type StaffGetPathParams struct {
}

func (p *StaffGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *StaffGetRequest) PathParams() *StaffGetPathParams {
	return r.pathParams
}

func (r *StaffGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *StaffGetRequest) Method() string {
	return r.method
}

func (s *Client) NewStaffGetRequestBody() StaffGetRequestBody {
	return StaffGetRequestBody{}
}

type StaffGetRequestBody struct {
}

func (r *StaffGetRequest) RequestBody() *StaffGetRequestBody {
	return &r.requestBody
}

func (r *StaffGetRequest) SetRequestBody(body StaffGetRequestBody) {
	r.requestBody = body
}

func (r *StaffGetRequest) NewResponseBody() *StaffGetResponseBody {
	return &StaffGetResponseBody{}
}

type StaffGetResponseBody struct {
	Staff
}

func (r *StaffGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("index2.php?option=com_webservices&controller=csv&method=hours.board.fetch&key={{.key}}&element=staff", r.PathParams())
}

func (r *StaffGetRequest) Do() (StaffGetResponseBody, error) {
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

	err = gocsv.Unmarshal(resp.Body, &responseBody.Staff)
	return *responseBody, err
}
