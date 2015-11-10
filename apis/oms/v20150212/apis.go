// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package oms

import (
	"errors"
	"fmt"

	"git.oschina.net/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
	_ = make(openapi.M)
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
	Product = "Oms"
	Style   = "RPC"
	Version = "2015-02-12"
)

// GetProductDefine version 2015-02-12
//
// required parameters:
//  name: DataType, type: string
//  name: ProductName, type: string
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetProductDefine(DataType, ProductName string, optional openapi.M) (*GetProductDefineResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "GetProductDefine")
	args.Query.Set("DataType", DataType)
	args.Query.Set("ProductName", ProductName)
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
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(GetProductDefineResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetProductDefineResponse represents the response of the api GetProductDefine.
type GetProductDefineResponse struct {
	DataType    string
	ProductList struct {
		Product []struct {
			ProductName string
			TypeList    struct {
				Type []struct {
					DataType string
				}
			}
		}
	}
	ProductName string
}

// GetUserData version 2015-02-12
//
// required parameters:
//  name: DataType, type: string
//  name: EndTime, type: string
//  name: ProductName, type: string
//  name: StartTime, type: string
//
// optional parameters:
//  name: MaxResult, type: int, min value: 1, max value: 200
//  name: NextToken, type: string
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) GetUserData(DataType, EndTime, ProductName, StartTime string, optional openapi.M) (*GetUserDataResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "GetUserData")
	args.Query.Set("DataType", DataType)
	args.Query.Set("EndTime", EndTime)
	args.Query.Set("ProductName", ProductName)
	args.Query.Set("StartTime", StartTime)
	if v, ok := optional["MaxResult"]; ok {
		if MaxResult, ok := v.(int); ok {
			if MaxResult < 1 {
				return nil, errors.New("MaxResult must be equal or greater than 1")
			}
			if MaxResult > 200 {
				return nil, errors.New("MaxResult must be equal or less than 200")
			}
			args.Query.Set("MaxResult", fmt.Sprint(MaxResult))
		} else {
			return nil, errors.New("MaxResult must be int")
		}
	}
	if v, ok := optional["NextToken"]; ok {
		if NextToken, ok := v.(string); ok {
			args.Query.Set("NextToken", NextToken)
		} else {
			return nil, errors.New("NextToken must be string")
		}
	}
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
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
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(GetUserDataResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserDataResponse represents the response of the api GetUserData.
type GetUserDataResponse struct {
	DataList struct {
		Data []struct {
			DataItems struct {
				DataItem []struct {
					Name  string
					Value string
				}
			}
		}
	}
	DataType    string
	NextToken   string
	ProductName string
}