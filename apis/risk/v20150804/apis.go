// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package risk

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
	Product = "Risk"
	Style   = "RPC"
	Version = "2015-08-04"
)

// FindRisk version 2015-08-04
//
// required parameters:
//  name: CodeType, type: string
//  name: IdType, type: string
//  name: MteeCode, type: string
//  name: UserId, type: string
//
// optional parameters:
//  name: Collina, type: string
//  name: Email, type: string
//  name: Extend, type: string
//  name: Ip, type: string
//  name: Phone, type: string
//  name: Umid, type: string
//  name: UmidToken, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) FindRisk(CodeType, IdType, MteeCode, UserId string, optional openapi.M) (*FindRiskResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "FindRisk")
	args.Query.Set("CodeType", CodeType)
	args.Query.Set("IdType", IdType)
	args.Query.Set("MteeCode", MteeCode)
	args.Query.Set("UserId", UserId)
	if v, ok := optional["Collina"]; ok {
		if Collina, ok := v.(string); ok {
			args.Query.Set("Collina", Collina)
		} else {
			return nil, errors.New("Collina must be string")
		}
	}
	if v, ok := optional["Email"]; ok {
		if Email, ok := v.(string); ok {
			args.Query.Set("Email", Email)
		} else {
			return nil, errors.New("Email must be string")
		}
	}
	if v, ok := optional["Extend"]; ok {
		if Extend, ok := v.(string); ok {
			args.Query.Set("Extend", Extend)
		} else {
			return nil, errors.New("Extend must be string")
		}
	}
	if v, ok := optional["Ip"]; ok {
		if Ip, ok := v.(string); ok {
			args.Query.Set("Ip", Ip)
		} else {
			return nil, errors.New("Ip must be string")
		}
	}
	if v, ok := optional["Phone"]; ok {
		if Phone, ok := v.(string); ok {
			args.Query.Set("Phone", Phone)
		} else {
			return nil, errors.New("Phone must be string")
		}
	}
	if v, ok := optional["Umid"]; ok {
		if Umid, ok := v.(string); ok {
			args.Query.Set("Umid", Umid)
		} else {
			return nil, errors.New("Umid must be string")
		}
	}
	if v, ok := optional["UmidToken"]; ok {
		if UmidToken, ok := v.(string); ok {
			args.Query.Set("UmidToken", UmidToken)
		} else {
			return nil, errors.New("UmidToken must be string")
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

	result := new(FindRiskResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// FindRiskResponse represents the response of the api FindRisk.
type FindRiskResponse struct {
	Code string
	Data struct {
		Action       string
		Message      string
		NoRisk       bool
		Tag          string
		VerifyDetail string
		VerifyType   string
	}
}

// QueryNameList version 2015-08-04
//
// required parameters:
//  name: DataType, type: string
//  name: DataValue, type: string
//  name: QueryLike, type: string
//  name: Type, type: string
//
// optional parameters:
//  name: Extend, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) QueryNameList(DataType, DataValue, QueryLike, Type string, optional openapi.M) (*QueryNameListResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "QueryNameList")
	args.Query.Set("DataType", DataType)
	args.Query.Set("DataValue", DataValue)
	args.Query.Set("QueryLike", QueryLike)
	args.Query.Set("Type", Type)
	if v, ok := optional["Extend"]; ok {
		if Extend, ok := v.(string); ok {
			args.Query.Set("Extend", Extend)
		} else {
			return nil, errors.New("Extend must be string")
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

	result := new(QueryNameListResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryNameListResponse represents the response of the api QueryNameList.
type QueryNameListResponse struct {
	Code          string
	RiskNameLists struct {
		RiskName []struct {
			DataType  string
			DataValue string
			Type      string
		}
	}
}

// SendVerifyCode version 2015-08-04
//
// required parameters:
//  name: ChannelType, type: string
//  name: CodeType, type: string
//  name: IdType, type: string
//  name: MteeCode, type: string
//  name: UserId, type: string
//
// optional parameters:
//  name: BizId, type: string
//  name: EventId, type: string
//  name: Extend, type: string
//  name: MessageParameters, type: string
//  name: MessageReiver, type: string
//  name: RequestId, type: string
//  name: TimeInterval, type: int
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) SendVerifyCode(ChannelType, CodeType, IdType, MteeCode, UserId string, optional openapi.M) (*SendVerifyCodeResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "SendVerifyCode")
	args.Query.Set("ChannelType", ChannelType)
	args.Query.Set("CodeType", CodeType)
	args.Query.Set("IdType", IdType)
	args.Query.Set("MteeCode", MteeCode)
	args.Query.Set("UserId", UserId)
	if v, ok := optional["BizId"]; ok {
		if BizId, ok := v.(string); ok {
			args.Query.Set("BizId", BizId)
		} else {
			return nil, errors.New("BizId must be string")
		}
	}
	if v, ok := optional["EventId"]; ok {
		if EventId, ok := v.(string); ok {
			args.Query.Set("EventId", EventId)
		} else {
			return nil, errors.New("EventId must be string")
		}
	}
	if v, ok := optional["Extend"]; ok {
		if Extend, ok := v.(string); ok {
			args.Query.Set("Extend", Extend)
		} else {
			return nil, errors.New("Extend must be string")
		}
	}
	if v, ok := optional["MessageParameters"]; ok {
		if MessageParameters, ok := v.(string); ok {
			args.Query.Set("MessageParameters", MessageParameters)
		} else {
			return nil, errors.New("MessageParameters must be string")
		}
	}
	if v, ok := optional["MessageReiver"]; ok {
		if MessageReiver, ok := v.(string); ok {
			args.Query.Set("MessageReiver", MessageReiver)
		} else {
			return nil, errors.New("MessageReiver must be string")
		}
	}
	if v, ok := optional["RequestId"]; ok {
		if RequestId, ok := v.(string); ok {
			args.Query.Set("RequestId", RequestId)
		} else {
			return nil, errors.New("RequestId must be string")
		}
	}
	if v, ok := optional["TimeInterval"]; ok {
		if TimeInterval, ok := v.(int); ok {
			args.Query.Set("TimeInterval", fmt.Sprint(TimeInterval))
		} else {
			return nil, errors.New("TimeInterval must be int")
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

	result := new(SendVerifyCodeResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// SendVerifyCodeResponse represents the response of the api SendVerifyCode.
type SendVerifyCodeResponse struct {
	Code string
}

// ValidateVerifyCode version 2015-08-04
//
// required parameters:
//  name: ChannelType, type: string
//  name: CodeType, type: string
//  name: IdType, type: string
//  name: MteeCode, type: string
//  name: UserId, type: string
//  name: VerifyCode, type: string
//
// optional parameters:
//  name: Collina, type: string
//  name: Extend, type: string
//  name: Ip, type: string
//  name: RequestId, type: string
//  name: UmidToken, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) ValidateVerifyCode(ChannelType, CodeType, IdType, MteeCode, UserId, VerifyCode string, optional openapi.M) (*ValidateVerifyCodeResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "ValidateVerifyCode")
	args.Query.Set("ChannelType", ChannelType)
	args.Query.Set("CodeType", CodeType)
	args.Query.Set("IdType", IdType)
	args.Query.Set("MteeCode", MteeCode)
	args.Query.Set("UserId", UserId)
	args.Query.Set("VerifyCode", VerifyCode)
	if v, ok := optional["Collina"]; ok {
		if Collina, ok := v.(string); ok {
			args.Query.Set("Collina", Collina)
		} else {
			return nil, errors.New("Collina must be string")
		}
	}
	if v, ok := optional["Extend"]; ok {
		if Extend, ok := v.(string); ok {
			args.Query.Set("Extend", Extend)
		} else {
			return nil, errors.New("Extend must be string")
		}
	}
	if v, ok := optional["Ip"]; ok {
		if Ip, ok := v.(string); ok {
			args.Query.Set("Ip", Ip)
		} else {
			return nil, errors.New("Ip must be string")
		}
	}
	if v, ok := optional["RequestId"]; ok {
		if RequestId, ok := v.(string); ok {
			args.Query.Set("RequestId", RequestId)
		} else {
			return nil, errors.New("RequestId must be string")
		}
	}
	if v, ok := optional["UmidToken"]; ok {
		if UmidToken, ok := v.(string); ok {
			args.Query.Set("UmidToken", UmidToken)
		} else {
			return nil, errors.New("UmidToken must be string")
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

	result := new(ValidateVerifyCodeResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// ValidateVerifyCodeResponse represents the response of the api ValidateVerifyCode.
type ValidateVerifyCodeResponse struct {
	Code string
}

// WriteUssc version 2015-08-04
//
// required parameters:
//  name: AppKey, type: string
//  name: ChannelType, type: string
//  name: CodeType, type: string
//  name: IdType, type: string
//  name: MteeCode, type: string
//  name: Sign, type: string
//  name: SignTime, type: string
//  name: UserId, type: string
//  name: VerifyResult, type: string
//
// optional parameters:
//  name: Collina, type: string
//  name: Extend, type: string
//  name: Ip, type: string
//  name: UmidToken, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) WriteUssc(AppKey, ChannelType, CodeType, IdType, MteeCode, Sign, SignTime, UserId, VerifyResult string, optional openapi.M) (*WriteUsscResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "WriteUssc")
	args.Query.Set("AppKey", AppKey)
	args.Query.Set("ChannelType", ChannelType)
	args.Query.Set("CodeType", CodeType)
	args.Query.Set("IdType", IdType)
	args.Query.Set("MteeCode", MteeCode)
	args.Query.Set("Sign", Sign)
	args.Query.Set("SignTime", SignTime)
	args.Query.Set("UserId", UserId)
	args.Query.Set("VerifyResult", VerifyResult)
	if v, ok := optional["Collina"]; ok {
		if Collina, ok := v.(string); ok {
			args.Query.Set("Collina", Collina)
		} else {
			return nil, errors.New("Collina must be string")
		}
	}
	if v, ok := optional["Extend"]; ok {
		if Extend, ok := v.(string); ok {
			args.Query.Set("Extend", Extend)
		} else {
			return nil, errors.New("Extend must be string")
		}
	}
	if v, ok := optional["Ip"]; ok {
		if Ip, ok := v.(string); ok {
			args.Query.Set("Ip", Ip)
		} else {
			return nil, errors.New("Ip must be string")
		}
	}
	if v, ok := optional["UmidToken"]; ok {
		if UmidToken, ok := v.(string); ok {
			args.Query.Set("UmidToken", UmidToken)
		} else {
			return nil, errors.New("UmidToken must be string")
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

	result := new(WriteUsscResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// WriteUsscResponse represents the response of the api WriteUssc.
type WriteUsscResponse struct {
	Code string
}
