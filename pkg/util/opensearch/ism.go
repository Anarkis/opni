package opensearch

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

type ISMApi struct {
	*opensearch.Client
}

func generateISMPath(name string) strings.Builder {
	var path strings.Builder
	path.Grow(1 + len("_plugins") + 1 + len("_ism") + 1 + len("policies") + 1 + len(name))
	path.WriteString("/")
	path.WriteString("_plugins")
	path.WriteString("/")
	path.WriteString("_ism")
	path.WriteString("/")
	path.WriteString("policies")
	path.WriteString("/")
	path.WriteString(name)
	return path
}

func (c *ISMApi) GetISM(ctx context.Context, name string) (*opensearchapi.Response, error) {
	method := "GET"
	path := generateISMPath(name)

	req, err := http.NewRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := c.Perform(req)
	if err != nil {
		return nil, err
	}

	return &opensearchapi.Response{StatusCode: res.StatusCode, Body: res.Body, Header: res.Header}, nil
}

func (c *ISMApi) CreateISM(ctx context.Context, name string, body io.Reader) (*opensearchapi.Response, error) {
	method := "PUT"
	path := generateISMPath(name)

	req, err := http.NewRequest(method, path.String(), body)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	req.Header.Add(headerContentType, jsonContentHeader)

	res, err := c.Perform(req)
	if err != nil {
		return nil, err
	}

	return &opensearchapi.Response{StatusCode: res.StatusCode, Body: res.Body, Header: res.Header}, nil
}

func (c *ISMApi) UpdateISM(ctx context.Context, name string, body io.Reader, seqNo int, primaryTerm int) (*opensearchapi.Response, error) {
	method := "PUT"

	path := generateISMPath(name)

	req, err := http.NewRequest(method, path.String(), body)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add(headerContentType, jsonContentHeader)

	query := req.URL.Query()
	query.Set("if_seq_no", fmt.Sprint(seqNo))
	query.Set("if_primary_term", fmt.Sprint(primaryTerm))
	req.URL.RawQuery = query.Encode()

	res, err := c.Perform(req)
	if err != nil {
		return nil, err
	}

	return &opensearchapi.Response{StatusCode: res.StatusCode, Body: res.Body, Header: res.Header}, nil
}
