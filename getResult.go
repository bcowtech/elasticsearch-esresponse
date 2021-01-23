// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package esresponse

import "encoding/json"

type GetResult struct {
	Index       string                 `json:"_index"`   // index meta field
	Type        string                 `json:"_type"`    // type meta field
	Id          string                 `json:"_id"`      // id meta field
	Uid         string                 `json:"_uid"`     // uid meta field (see MapperService.java for all meta fields)
	Routing     string                 `json:"_routing"` // routing meta field
	Parent      string                 `json:"_parent"`  // parent meta field
	Version     *int64                 `json:"_version"` // version number, when Version is set to true in SearchService
	SeqNo       *int64                 `json:"_seq_no"`
	PrimaryTerm *int64                 `json:"_primary_term"`
	Source      json.RawMessage        `json:"_source,omitempty"`
	Found       bool                   `json:"found,omitempty"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
	//Error     string                 `json:"error,omitempty"` // used only in MultiGet
	// TODO double-check that MultiGet now returns details error information
	Error *ErrorDetails `json:"error,omitempty"` // only used in MultiGet
}
