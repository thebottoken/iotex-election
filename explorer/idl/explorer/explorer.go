// Code generated by idl2go from JSON generated by Barrister v0.1.6
package explorer

import (
	"fmt"
	"reflect"

	"github.com/coopernurse/barrister-go"
)

const BarristerVersion string = "0.1.6"
const BarristerChecksum string = "7928b11b3d82a4f9a4284f010dc6463a"
const BarristerDateGenerated int64 = 1549740539248000000

type ChainMeta struct {
	Height          string `json:"height"`
	TotalCandidates int64  `json:"totalCandidates"`
}

type Bucket struct {
	Voter             string `json:"voter"`
	Votes             string `json:"votes"`
	WeightedVotes     string `json:"weightedVotes"`
	RemainingDuration int64  `json:"remainingDuration"`
}

type Candidate struct {
	Name               string `json:"name"`
	Address            string `json:"pubKey"`
	TotalWeightedVotes string `json:"totalWeightedVotes"`
}

type GetCandidatesRequest struct {
	Height string `json:"height"`
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
}

type GetBucketsByCandidateRequest struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
}

type Explorer interface {
	GetMeta() (ChainMeta, error)
	GetCandidates(request GetCandidatesRequest) ([]Candidate, error)
	GetBucketsByCandidate(request GetBucketsByCandidateRequest) ([]Bucket, error)
}

func NewExplorerProxy(c barrister.Client) Explorer {
	return ExplorerProxy{c, barrister.MustParseIdlJson([]byte(IdlJsonRaw))}
}

type ExplorerProxy struct {
	client barrister.Client
	idl    *barrister.Idl
}

func (_p ExplorerProxy) GetMeta() (ChainMeta, error) {
	_res, _err := _p.client.Call("Explorer.getMeta")
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getMeta").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(ChainMeta{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(ChainMeta)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getMeta returned invalid type: %v", _t)
			return ChainMeta{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return ChainMeta{}, _err
}

func (_p ExplorerProxy) GetCandidates(request GetCandidatesRequest) ([]Candidate, error) {
	_res, _err := _p.client.Call("Explorer.getCandidates", request)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getCandidates").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Candidate{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Candidate)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getCandidates returned invalid type: %v", _t)
			return []Candidate{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Candidate{}, _err
}

func (_p ExplorerProxy) GetBucketsByCandidate(request GetBucketsByCandidateRequest) ([]Bucket, error) {
	_res, _err := _p.client.Call("Explorer.getBucketsByCandidate", request)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getBucketsByCandidate").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Bucket{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Bucket)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getBucketsByCandidate returned invalid type: %v", _t)
			return []Bucket{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Bucket{}, _err
}

func NewJSONServer(idl *barrister.Idl, forceASCII bool, explorer Explorer) barrister.Server {
	return NewServer(idl, &barrister.JsonSerializer{forceASCII}, explorer)
}

func NewServer(idl *barrister.Idl, ser barrister.Serializer, explorer Explorer) barrister.Server {
	_svr := barrister.NewServer(idl, ser)
	_svr.AddHandler("Explorer", explorer)
	return _svr
}

var IdlJsonRaw = `[
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "Copyright (c) 2018 IoTeX\nThis is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no\nwarranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent\npermitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache\nLicense 2.0 that can be found in the LICENSE file.",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "To compile this file:\n1. Install the barrister translator (IDL -\u003e JSON)\nyou need to be root (or use sudo)\npip install barrister",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "2. Install barrister-go\ngo get github.com/coopernurse/barrister-go\ngo install github.com/coopernurse/barrister-go/idl2go",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "3. barrister explorer.idl | $GOPATH/bin/idl2go -i -p explorer",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "ChainMeta",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "height",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "totalCandidates",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Bucket",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "voter",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": "hex string"
            },
            {
                "name": "votes",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "weightedVotes",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "remainingDuration",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": "seconds"
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Candidate",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "name",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "pubKey",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": "hex string"
            },
            {
                "name": "totalWeightedVotes",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "GetCandidatesRequest",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "height",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "offset",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "limit",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "GetBucketsByCandidateRequest",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "name",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "height",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "offset",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "limit",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "interface",
        "name": "Explorer",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": [
            {
                "name": "getMeta",
                "comment": "get the blockchain meta data",
                "params": [],
                "returns": {
                    "name": "",
                    "type": "ChainMeta",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getCandidates",
                "comment": "get candidates",
                "params": [
                    {
                        "name": "request",
                        "type": "GetCandidatesRequest",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Candidate",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getBucketsByCandidate",
                "comment": "get buckets by candidate",
                "params": [
                    {
                        "name": "request",
                        "type": "GetBucketsByCandidateRequest",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Bucket",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            }
        ],
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "meta",
        "name": "",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "0.1.6",
        "date_generated": 1549740539248,
        "checksum": "7928b11b3d82a4f9a4284f010dc6463a"
    }
]`
