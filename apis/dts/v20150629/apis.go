// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package dts

import (
	"errors"
	"fmt"

	"git.oschina.net/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
)

type API struct {
	openapi.Service
}

func New(accessKeyId, accessKeySecret string) API {
	return API{openapi.NewService(accessKeyId, accessKeySecret)}
}

func NewParams() openapi.Params {
	args := openapi.NewParams()
	args.Product = Product
	args.Style = Style
	args.Version = Version
	return args
}

const (
	Product = "Dts"
	Style   = "RPC"
	Version = "2015-06-29"
)

// DrcGuidRouteApi version 2015-06-29
//
// required parameters:
//  name: guid, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
func (api API) DrcGuidRouteApi(guid string, optional openapi.M) (*DrcGuidRouteApiResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DrcGuidRouteApi")
	args.Query.Set("guid", guid)
	args.Scheme = "http"
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}

	result := new(DrcGuidRouteApiResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DrcGuidRouteApiResponse represents the response of the api DrcGuidRouteApi.
type DrcGuidRouteApiResponse struct {
	Data    string `json:"data" xml:"data"`
	Success bool   `json:"success" xml:"success"`
}
