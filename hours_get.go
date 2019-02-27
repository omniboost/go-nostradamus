package nostradamus

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gocarina/gocsv"
)

func (c *Client) NewHoursGetRequest() HoursGetRequest {
	return HoursGetRequest{
		client:      c,
		queryParams: c.NewHoursGetQueryParams(),
		pathParams:  c.NewHoursGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewHoursGetRequestBody(),
	}
}

type HoursGetRequest struct {
	client      *Client
	queryParams *HoursGetQueryParams
	pathParams  *HoursGetPathParams
	method      string
	headers     http.Header
	requestBody HoursGetRequestBody
}

func (c *Client) NewHoursGetQueryParams() *HoursGetQueryParams {
	return &HoursGetQueryParams{}
}

type HoursGetQueryParams struct {
	Start time.Time `schema:"start"`
	End   time.Time `schema:"end"`
}

func (p HoursGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *HoursGetRequest) QueryParams() *HoursGetQueryParams {
	return r.queryParams
}

func (c *Client) NewHoursGetPathParams() *HoursGetPathParams {
	return &HoursGetPathParams{}
}

type HoursGetPathParams struct {
}

func (p *HoursGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *HoursGetRequest) PathParams() *HoursGetPathParams {
	return r.pathParams
}

func (r *HoursGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *HoursGetRequest) Method() string {
	return r.method
}

func (s *Client) NewHoursGetRequestBody() HoursGetRequestBody {
	return HoursGetRequestBody{}
}

type HoursGetRequestBody struct {
}

func (r *HoursGetRequest) RequestBody() *HoursGetRequestBody {
	return &r.requestBody
}

func (r *HoursGetRequest) SetRequestBody(body HoursGetRequestBody) {
	r.requestBody = body
}

func (r *HoursGetRequest) NewResponseBody() *HoursGetResponseBody {
	return &HoursGetResponseBody{}
}

type HoursGetResponseBody struct {
	Hours
}

func (r *HoursGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("index2.php?option=com_webservices&controller=csv&method=hours.board.fetch&key={{.key}}&element=hours", r.PathParams())
}

func (r *HoursGetRequest) Do() (HoursGetResponseBody, error) {
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

	err = gocsv.Unmarshal(resp.Body, &responseBody.Hours)
	return *responseBody, err
}
