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

func AsCountResult(responseBody io.Reader) (*CountResult, error) {
	var result = new(CountResult)

	err := decode(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AsGetResult(responseBody io.Reader) (*GetResult, error) {
	var result = new(GetResult)

	err := decode(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AsDeleteResponse(responseBody io.Reader) (*DeleteResponse, error) {
	var result = new(DeleteResponse)

	err := decode(responseBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
