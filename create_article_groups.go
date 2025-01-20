package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-nostradamus/utils"
)

func (c *Client) NewCreateArticleGroupsRequest() CreateArticleGroupsRequest {
	r := CreateArticleGroupsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateArticleGroupsRequest struct {
	client      *Client
	queryParams *CreateArticleGroupsRequestQueryParams
	pathParams  *CreateArticleGroupsRequestPathParams
	method      string
	headers     http.Header
	requestBody CreateArticleGroupsRequestBody
}

func (r CreateArticleGroupsRequest) NewQueryParams() *CreateArticleGroupsRequestQueryParams {
	return &CreateArticleGroupsRequestQueryParams{}
}

type CreateArticleGroupsRequestQueryParams struct{}

func (p CreateArticleGroupsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateArticleGroupsRequest) QueryParams() *CreateArticleGroupsRequestQueryParams {
	return r.queryParams
}

func (r CreateArticleGroupsRequest) NewPathParams() *CreateArticleGroupsRequestPathParams {
	return &CreateArticleGroupsRequestPathParams{}
}

type CreateArticleGroupsRequestPathParams struct{}

func (p *CreateArticleGroupsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateArticleGroupsRequest) PathParams() *CreateArticleGroupsRequestPathParams {
	return r.pathParams
}

func (r *CreateArticleGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CreateArticleGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateArticleGroupsRequest) Method() string {
	return r.method
}

func (r CreateArticleGroupsRequest) NewRequestBody() CreateArticleGroupsRequestBody {
	return CreateArticleGroupsRequestBody{}
}

type CreateArticleGroupsRequestBody ArticleGroups

func (r *CreateArticleGroupsRequest) RequestBody() *CreateArticleGroupsRequestBody {
	return &r.requestBody
}

func (r *CreateArticleGroupsRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *CreateArticleGroupsRequest) SetRequestBody(body CreateArticleGroupsRequestBody) {
	r.requestBody = body
}

func (r *CreateArticleGroupsRequest) NewResponseBody() *CreateArticleGroupsRequestResponseBody {
	return &CreateArticleGroupsRequestResponseBody{}
}

type CreateArticleGroupsRequestResponseBody struct {
	Function struct {
		Test string `xml:"test"`
	} `xml:"function"`
}

func (r *CreateArticleGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CreateArticleGroupsRequest) Do() (CreateArticleGroupsRequestResponseBody, error) {
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
