package nostradamus

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"

	"github.com/alecthomas/template"
	"github.com/gocarina/gocsv"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-nostradamus/" + libraryVersion
	mediaType      = "text/html"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "{{.customer}}.nostradamus.nu",
		Path:   "",
	}
)

type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	customer string
	key      string

	// User agent for client
	userAgent string

	mediaType string
	charset   string
}

func NewClient(httpClient *http.Client, customer string, key string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http: httpClient,
	}

	client.SetCustomer(customer)
	client.SetKey(key)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)

	return client
}

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Customer() string {
	return c.customer
}

func (c *Client) SetCustomer(customer string) {
	c.customer = customer
}

func (c *Client) Key() string {
	return c.key
}

func (c *Client) SetKey(key string) {
	c.key = key
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) GetEndpointURL(relative string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	relativeURL, err := url.Parse(relative)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Host = strings.Replace(clientURL.Host, "{{.customer}}", c.Customer(), -1)
	clientURL.Path = path.Join(clientURL.Path, relativeURL.Path)
	clientURL.RawQuery = strings.Replace(clientURL.RawQuery, "{{.key}}", c.Key(), -1)
	relativeURL.RawQuery = strings.Replace(relativeURL.RawQuery, "{{.key}}", c.Key(), -1)

	query := url.Values{}
	for k, v := range clientURL.Query() {
		query[k] = append(query[k], v...)
	}
	for k, v := range relativeURL.Query() {
		query[k] = append(query[k], v...)
	}
	clientURL.RawQuery = query.Encode()

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	// if body != nil {
	// 	err := json.NewEncoder(buf).Encode(body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	// req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	// req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	// don't close body io.Reader: it will be used later by the csv parser
	// defer func() {
	// 	if rerr := httpResp.Body.Close(); err == nil {
	// 		err = rerr
	// 	}
	// }()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	// err = gocsv.Unmarshal(httpResp.Body, &responseBody)
	// if err != nil {
	// 	return httpResp, err
	// }

	return httpResp, nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}

type ErrorResponse struct {
	error

	// HTTP response that caused this error
	Response *http.Response `json:"-"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprint("err")
}

func init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.FieldsPerRecord = -1
		return r
	})
}
