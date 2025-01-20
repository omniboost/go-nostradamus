package nostradamus_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	nostradamus "github.com/omniboost/go-nostradamus"
)

var (
	client *nostradamus.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	customer := os.Getenv("CUSTOMER")
	key := os.Getenv("KEY")
	debug := os.Getenv("DEBUG")

	client = nostradamus.NewClient(nil, customer, key)
	if debug != "" {
		client.SetDebug(true)
	}
	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
