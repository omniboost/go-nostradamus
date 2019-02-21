package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/gocarina/gocsv"
)

func (c *Client) NewDepartmentsGetRequest() DepartmentsGetRequest {
	return DepartmentsGetRequest{
		client:      c,
		queryParams: c.NewDepartmentsGetQueryParams(),
		pathParams:  c.NewDepartmentsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewDepartmentsGetRequestBody(),
	}
}

type DepartmentsGetRequest struct {
	client      *Client
	queryParams *DepartmentsGetQueryParams
	pathParams  *DepartmentsGetPathParams
	method      string
	headers     http.Header
	requestBody DepartmentsGetRequestBody
}

func (c *Client) NewDepartmentsGetQueryParams() *DepartmentsGetQueryParams {
	return &DepartmentsGetQueryParams{}
}

type DepartmentsGetQueryParams struct {
}

func (p DepartmentsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *DepartmentsGetRequest) QueryParams() *DepartmentsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewDepartmentsGetPathParams() *DepartmentsGetPathParams {
	return &DepartmentsGetPathParams{}
}

type DepartmentsGetPathParams struct {
}

func (p *DepartmentsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DepartmentsGetRequest) PathParams() *DepartmentsGetPathParams {
	return r.pathParams
}

func (r *DepartmentsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DepartmentsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewDepartmentsGetRequestBody() DepartmentsGetRequestBody {
	return DepartmentsGetRequestBody{}
}

type DepartmentsGetRequestBody struct {
}

func (r *DepartmentsGetRequest) RequestBody() *DepartmentsGetRequestBody {
	return &r.requestBody
}

func (r *DepartmentsGetRequest) SetRequestBody(body DepartmentsGetRequestBody) {
	r.requestBody = body
}

func (r *DepartmentsGetRequest) NewResponseBody() *DepartmentsGetResponseBody {
	return &DepartmentsGetResponseBody{}
}

type DepartmentsGetResponseBody struct {
	Departments
}

func (r *DepartmentsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("index2.php?option=com_webservices&controller=csv&method=hours.board.fetch&key={{.key}}&element=departments", r.PathParams())
}

func (r *DepartmentsGetRequest) Do() (DepartmentsGetResponseBody, error) {
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

	err = gocsv.Unmarshal(resp.Body, &responseBody.Departments)
	return *responseBody, err
}
