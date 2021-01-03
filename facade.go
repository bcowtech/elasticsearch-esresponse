package esresponse

import (
	"io"
)

func AsSearchResult(responseBody io.Reader) (*SearchResult, error) {
	var result = new(SearchResult)

	err := decode(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AsIndexResponse(responseBody io.Reader) (*IndexResponse, error) {
	var result = new(IndexResponse)

	err := decode(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
