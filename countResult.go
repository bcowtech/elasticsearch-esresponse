// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package esresponse

import (
	"net/http"
)

// CountResult is the result of a count in Elasticsearch.
type CountResult struct {
	Header          http.Header          `json:"-"`
	Count    		int64                `json:"count,omitempty"`             		// count 
	Error           *ErrorDetails        `json:"error,omitempty"`        // only used in MultiGet
	Shards          *ShardsInfo          `json:"_shards,omitempty"`      // shard information
	Status          int                  `json:"status,omitempty"`       // used in MultiSearch
}

// TotalCount is a convenience function to return the number of Count for
// a Count result. The return value might not be accurate, unless
func (r *CountResult) TotalCount() int64 {
	if r != nil  {
		return r.Count
	}
	return 0
}
