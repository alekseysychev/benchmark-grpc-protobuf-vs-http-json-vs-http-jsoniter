package benchmarks

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	httpjson "github.com/alekseysychev/benchmark-grpc-protobuf-vs-http-json-vs-http-jsoniter/http-json"
	httpjsoniter "github.com/alekseysychev/benchmark-grpc-protobuf-vs-http-json-vs-http-jsoniter/http-jsoniter"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	go httpjsoniter.Start()
	time.Sleep(time.Second)
}

func BenchmarkHTTPJSONIter(b *testing.B) {
	client := &http.Client{}

	for n := 0; n < b.N; n++ {
		doPostJsonIter(client, b)
	}
}

func doPostJsonIter(client *http.Client, b *testing.B) {
	u := &httpjson.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	}
	buf := new(bytes.Buffer)
	jsoniter.NewEncoder(buf).Encode(u)

	resp, err := client.Post("http://127.0.0.1:60002/", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target httpjson.Response
	decodeErr := jsoniter.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.User == nil || target.User.ID != "1000000" {
		b.Fatalf("http response is wrong: %v", resp)
	}
}
