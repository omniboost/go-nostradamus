package nostradamus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-nostradamus/utils"
)

func (c *Client) NewCreateArticleSubGroupsRequest() CreateArticleSubGroupsRequest {
	r := CreateArticleSubGroupsRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreateArticleSubGroupsRequest struct {
	client      *Client
	queryParams *CreateArticleSubGroupsRequestQueryParams
	pathParams  *CreateArticleSubGroupsRequestPathParams
	method      string
	headers     http.Header
	requestBody CreateArticleSubGroupsRequestBody
}

func (r CreateArticleSubGroupsRequest) NewQueryParams() *CreateArticleSubGroupsRequestQueryParams {
	return &CreateArticleSubGroupsRequestQueryParams{}
}

type CreateArticleSubGroupsRequestQueryParams struct{}

func (p CreateArticleSubGroupsRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreateArticleSubGroupsRequest) QueryParams() *CreateArticleSubGroupsRequestQueryParams {
	return r.queryParams
}

func (r CreateArticleSubGroupsRequest) NewPathParams() *CreateArticleSubGroupsRequestPathParams {
	return &CreateArticleSubGroupsRequestPathParams{}
}

type CreateArticleSubGroupsRequestPathParams struct{}

func (p *CreateArticleSubGroupsRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CreateArticleSubGroupsRequest) PathParams() *CreateArticleSubGroupsRequestPathParams {
	return r.pathParams
}

func (r *CreateArticleSubGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CreateArticleSubGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreateArticleSubGroupsRequest) Method() string {
	return r.method
}

func (r CreateArticleSubGroupsRequest) NewRequestBody() CreateArticleSubGroupsRequestBody {
	return CreateArticleSubGroupsRequestBody{}
}

type CreateArticleSubGroupsRequestBody ArticleSubGroups

func (r *CreateArticleSubGroupsRequest) RequestBody() *CreateArticleSubGroupsRequestBody {
	return &r.requestBody
}

func (r *CreateArticleSubGroupsRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *CreateArticleSubGroupsRequest) SetRequestBody(body CreateArticleSubGroupsRequestBody) {
	r.requestBody = body
}

func (r *CreateArticleSubGroupsRequest) NewResponseBody() *CreateArticleSubGroupsRequestResponseBody {
	return &CreateArticleSubGroupsRequestResponseBody{}
}

type CreateArticleSubGroupsRequestResponseBody struct {
	Function struct {
		Test string `xml:"test"`
	} `xml:"function"`
}

func (r *CreateArticleSubGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *CreateArticleSubGroupsRequest) Do() (CreateArticleSubGroupsRequestResponseBody, error) {
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
