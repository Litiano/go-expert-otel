package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func RequestWithTimeout(ctx context.Context, timeout time.Duration, method, url string, body io.Reader, requestId string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	//req.Header.Set("X-Request-Id", requestId)

	resp, err := http.DefaultClient.Do(req)
	if resp != nil && resp.StatusCode != http.StatusOK {
		fmt.Println("Error: ", err)
		return resp, errors.New("Request error with status code: " + resp.Status)
	}

	return resp, err
}
