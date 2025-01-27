package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-nostradamus/utils"
)

func (c *Client) NewCreateSalesRequest() CreateSalesRequest {
	r := CreateSalesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateSalesRequest struct {
	client      *Client
	queryParams *CreateSalesRequestQueryParams
	pathParams  *CreateSalesRequestPathParams
	method      string
	headers     http.Header
	requestBody CreateSalesRequestBody
}

func (r CreateSalesRequest) NewQueryParams() *CreateSalesRequestQueryParams {
	return &CreateSalesRequestQueryParams{}
}

type CreateSalesRequestQueryParams struct{}

func (p CreateSalesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateSalesRequest) QueryParams() *CreateSalesRequestQueryParams {
	return r.queryParams
}

func (r CreateSalesRequest) NewPathParams() *CreateSalesRequestPathParams {
	return &CreateSalesRequestPathParams{}
}

type CreateSalesRequestPathParams struct{}

func (p *CreateSalesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateSalesRequest) PathParams() *CreateSalesRequestPathParams {
	return r.pathParams
}

func (r *CreateSalesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CreateSalesRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateSalesRequest) Method() string {
	return r.method
}

func (r CreateSalesRequest) NewRequestBody() CreateSalesRequestBody {
	return CreateSalesRequestBody{}
}

type CreateSalesRequestBody Sales

func (r *CreateSalesRequest) RequestBody() *CreateSalesRequestBody {
	return &r.requestBody
}

func (r *CreateSalesRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *CreateSalesRequest) SetRequestBody(body CreateSalesRequestBody) {
	r.requestBody = body
}

func (r *CreateSalesRequest) NewResponseBody() *CreateSalesRequestResponseBody {
	return &CreateSalesRequestResponseBody{}
}

type CreateSalesRequestResponseBody struct{}

func (r *CreateSalesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CreateSalesRequest) Do() (CreateSalesRequestResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), *r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// // Set Token
	// err = r.client.InitToken(req)
	// if err != nil {
	// 	return *r.NewResponseBody(), err
	// }

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
