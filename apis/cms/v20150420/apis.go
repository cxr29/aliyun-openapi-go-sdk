// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package cms

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
	Product = "Cms"
	Style   = "RPC"
	Version = "2015-04-20"
)

// DescribeMetricDatum version 2015-04-20
//
// optional parameters:
//  name: Dimensions, type: string
//  name: EndTime, type: string
//  name: GroupName, type: string
//  name: Length, type: int
//  name: MetricName, type: string
//  name: Namespace, type: string
//  name: NextToken, type: string
//  name: Period, type: string
//  name: StartTime, type: string
//  name: Statistics, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
func (api API) DescribeMetricDatum(optional openapi.M) (*DescribeMetricDatumResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "DescribeMetricDatum")
	args.Scheme = "http"
	if v, ok := optional["Dimensions"]; ok {
		if Dimensions, ok := v.(string); ok {
			args.Query.Set("Dimensions", Dimensions)
		} else {
			return nil, errors.New("Dimensions must be string")
		}
	}
	if v, ok := optional["EndTime"]; ok {
		if EndTime, ok := v.(string); ok {
			args.Query.Set("EndTime", EndTime)
		} else {
			return nil, errors.New("EndTime must be string")
		}
	}
	if v, ok := optional["GroupName"]; ok {
		if GroupName, ok := v.(string); ok {
			args.Query.Set("GroupName", GroupName)
		} else {
			return nil, errors.New("GroupName must be string")
		}
	}
	if v, ok := optional["Length"]; ok {
		if Length, ok := v.(int); ok {
			args.Query.Set("Length", fmt.Sprint(Length))
		} else {
			return nil, errors.New("Length must be int")
		}
	}
	if v, ok := optional["MetricName"]; ok {
		if MetricName, ok := v.(string); ok {
			args.Query.Set("MetricName", MetricName)
		} else {
			return nil, errors.New("MetricName must be string")
		}
	}
	if v, ok := optional["Namespace"]; ok {
		if Namespace, ok := v.(string); ok {
			args.Query.Set("Namespace", Namespace)
		} else {
			return nil, errors.New("Namespace must be string")
		}
	}
	if v, ok := optional["NextToken"]; ok {
		if NextToken, ok := v.(string); ok {
			args.Query.Set("NextToken", NextToken)
		} else {
			return nil, errors.New("NextToken must be string")
		}
	}
	if v, ok := optional["Period"]; ok {
		if Period, ok := v.(string); ok {
			args.Query.Set("Period", Period)
		} else {
			return nil, errors.New("Period must be string")
		}
	}
	if v, ok := optional["StartTime"]; ok {
		if StartTime, ok := v.(string); ok {
			args.Query.Set("StartTime", StartTime)
		} else {
			return nil, errors.New("StartTime must be string")
		}
	}
	if v, ok := optional["Statistics"]; ok {
		if Statistics, ok := v.(string); ok {
			args.Query.Set("Statistics", Statistics)
		} else {
			return nil, errors.New("Statistics must be string")
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

	result := new(DescribeMetricDatumResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetricDatumResponse represents the response of the api DescribeMetricDatum.
type DescribeMetricDatumResponse struct {
	Code      string
	Message   string
	NextToken string
}
