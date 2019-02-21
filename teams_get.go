package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/gocarina/gocsv"
)

func (c *Client) NewTeamsGetRequest() TeamsGetRequest {
	return TeamsGetRequest{
		client:      c,
		queryParams: c.NewTeamsGetQueryParams(),
		pathParams:  c.NewTeamsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewTeamsGetRequestBody(),
	}
}

type TeamsGetRequest struct {
	client      *Client
	queryParams *TeamsGetQueryParams
	pathParams  *TeamsGetPathParams
	method      string
	headers     http.Header
	requestBody TeamsGetRequestBody
}

func (c *Client) NewTeamsGetQueryParams() *TeamsGetQueryParams {
	return &TeamsGetQueryParams{}
}

type TeamsGetQueryParams struct {
}

func (p TeamsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TeamsGetRequest) QueryParams() *TeamsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewTeamsGetPathParams() *TeamsGetPathParams {
	return &TeamsGetPathParams{}
}

type TeamsGetPathParams struct {
}

func (p *TeamsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TeamsGetRequest) PathParams() *TeamsGetPathParams {
	return r.pathParams
}

func (r *TeamsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TeamsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewTeamsGetRequestBody() TeamsGetRequestBody {
	return TeamsGetRequestBody{}
}

type TeamsGetRequestBody struct {
}

func (r *TeamsGetRequest) RequestBody() *TeamsGetRequestBody {
	return &r.requestBody
}

func (r *TeamsGetRequest) SetRequestBody(body TeamsGetRequestBody) {
	r.requestBody = body
}

func (r *TeamsGetRequest) NewResponseBody() *TeamsGetResponseBody {
	return &TeamsGetResponseBody{}
}

type TeamsGetResponseBody struct {
	Teams
}

func (r *TeamsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("index2.php?option=com_webservices&controller=csv&method=hours.board.fetch&key={{.key}}&element=teams", r.PathParams())
}

func (r *TeamsGetRequest) Do() (TeamsGetResponseBody, error) {
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

	err = gocsv.Unmarshal(resp.Body, &responseBody.Teams)
	return *responseBody, err
}
