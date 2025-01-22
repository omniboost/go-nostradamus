package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-nostradamus/utils"
)

func (c *Client) NewCreateArticlesRequest() CreateArticlesRequest {
	r := CreateArticlesRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateArticlesRequest struct {
	client      *Client
	queryParams *CreateArticlesRequestQueryParams
	pathParams  *CreateArticlesRequestPathParams
	method      string
	headers     http.Header
	requestBody CreateArticlesRequestBody
}

func (r CreateArticlesRequest) NewQueryParams() *CreateArticlesRequestQueryParams {
	return &CreateArticlesRequestQueryParams{}
}

type CreateArticlesRequestQueryParams struct{}

func (p CreateArticlesRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateArticlesRequest) QueryParams() *CreateArticlesRequestQueryParams {
	return r.queryParams
}

func (r CreateArticlesRequest) NewPathParams() *CreateArticlesRequestPathParams {
	return &CreateArticlesRequestPathParams{}
}

type CreateArticlesRequestPathParams struct{}

func (p *CreateArticlesRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateArticlesRequest) PathParams() *CreateArticlesRequestPathParams {
	return r.pathParams
}

func (r *CreateArticlesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CreateArticlesRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateArticlesRequest) Method() string {
	return r.method
}

func (r CreateArticlesRequest) NewRequestBody() CreateArticlesRequestBody {
	return CreateArticlesRequestBody{}
}

type CreateArticlesRequestBody CreateArticlesDto

func (r *CreateArticlesRequest) RequestBody() *CreateArticlesRequestBody {
	return &r.requestBody
}

func (r *CreateArticlesRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *CreateArticlesRequest) SetRequestBody(body CreateArticlesRequestBody) {
	r.requestBody = body
}

func (r *CreateArticlesRequest) NewResponseBody() *CreateArticlesRequestResponseBody {
	return &CreateArticlesRequestResponseBody{}
}

type CreateArticlesRequestResponseBody struct {
	Function struct {
		Test string `xml:"test"`
	} `xml:"function"`
}

func (r *CreateArticlesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CreateArticlesRequest) Do() (CreateArticlesRequestResponseBody, error) {
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
