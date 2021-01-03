// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package esresponse

import "errors"

var (
	// ErrNoClient is raised when no Elasticsearch node is available.
	ErrNoClient = errors.New("no Elasticsearch node available")

	// ErrRetry is raised when a request cannot be executed after the configured
	// number of retries.
	ErrRetry = errors.New("cannot connect after several retries")

	// ErrTimeout is raised when a request timed out, e.g. when WaitForStatus
	// didn't return in time.
	ErrTimeout = errors.New("timeout")

	// ErrBulkItemRetry is returned in BulkProcessor from a worker when
	// a response item needs to be retried.
	ErrBulkItemRetry = errors.New("elastic: uncommitted bulk response items")

	// ErrResponseSize is raised if a response body exceeds the given max body size.
	ErrResponseSize = errors.New("elastic: response size too large")
)
