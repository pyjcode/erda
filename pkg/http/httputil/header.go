// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httputil

import (
	"fmt"
	"net/http"
	"strings"
)

// dice 公共 Header
const (
	UserHeader                 = "User-ID"
	OrgHeader                  = "Org-ID"
	InternalHeader             = "Internal-Client"        // 内部服务间调用时使用
	InternalActionHeader       = "Internal-Action-Client" // action calls the api header
	RequestIDHeader            = "RequestID"
	UserInfoDesensitizedHeader = "Openapi-Userinfo-Desensitized"
	LangHeader                 = "lang"

	ClientIDHeader   = "Client-ID"
	ClientNameHeader = "Client-Name"
)

func HeaderContains[Header http.Header | []String | []string | ~string, String ~string](header Header, s String) bool {
	switch i := (interface{})(header).(type) {
	case http.Header:
		for _, vv := range i {
			for _, v := range vv {
				if strings.Contains(v, string(s)) {
					return true
				}
			}
		}
	case []String:
		for _, v := range i {
			if strings.Contains(string(v), string(s)) {
				return true
			}
		}
	case []string:
		for _, v := range i {
			if strings.Contains(string(v), string(s)) {
				return true
			}
		}
	default:
		return strings.Contains(fmt.Sprintf("%s", i), string(s))
	}
	return false
}
