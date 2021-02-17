package httprq

import (
	"fmt"
	"log"

	resty "github.com/go-resty/resty/v2"
)

// Request ...
type Request struct {
	client *resty.Client
}

// NewHTTPRq ...
func NewHTTPRq() *Request {
	return &Request{
		client: resty.New(),
	}
}

// PostRq ...
func (r Request) PostRq(url, body string) {
	rs, err := r.client.R().
		EnableTrace().
		SetBody(body).
		Post(url)

	if err != nil {
		log.Fatal("Error while sending POST request!\n", err)
	}
	defer rs.RawBody().Close()

	printHTTPInfo(rs)
}

// GetRq ...
func (r Request) GetRq(url string) {
	rs, err := r.client.R().
		EnableTrace().Get(url)

	if err != nil {
		log.Fatal("Error while sending GET request!\n", err)
	}
	defer rs.RawBody().Close()

	printHTTPInfo(rs)
}

func printHTTPInfo(resp *resty.Response) {
	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  StatusCode :", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  ReceivedAt :", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()

	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}
