// auth-password - Username/Password Authentication
// Copyright 2015 Dean Troyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keystonev3

import (
	"encoding/json"
	"errors"
)

type OSDomain struct {
  Name string `json:"name"`
}

type OSUser struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Domain OSDomain `json:"domain"`
}

type OSPassword struct {
	User OSUser `json:"user"`
}

type OSProject struct {
	Name string `json:"name"`
	Domain OSDomain `json:"domain"`
}

type OSScope struct {
	Project OSProject `json:"project"`
}

type OSIdentity struct {
    Methods []string `json:"methods"`
    Password OSPassword `json:"password"`
}

type OSAuth struct {
    Identity OSIdentity `json:"identity"`
    Scope    OSScope    `json:"scope"`
}

type UserPassV3 struct {
	Auth OSAuth `json:"auth"`
}

func NewUserPassV3(ao AuthOpts) (upv3 *UserPassV3, err error) {
	// Validate incoming values
	if ao.AuthUrl == "" {
		err = errors.New("AuthUrl required")
		return nil, err
	}
	if ao.Username == "" {
		err = errors.New("Username required")
		return nil, err
	}
	if ao.Password == "" {
		err = errors.New("Password required")
		return nil, err
	}
	if ao.Project == "" {
		err = errors.New("Project required")
		return nil, err
	}
	upv3 = &UserPassV3{
		Auth: OSAuth{
			Identity: OSIdentity{
                        	Methods: []string{"password"},
				Password: OSPassword{
					User: OSUser{
                                    		Name: ao.Username,
				    		Password: ao.Password,
						Domain: OSDomain{
							Name: "default",
						},
                                	},
				},
			},
			Scope: OSScope{
				Project: OSProject{
					Name: ao.Project,
					Domain: OSDomain{
						Name: "default",
					},
				},
			},
		},
	}
	return upv3, nil
}

// Produce JSON output
func (s *UserPassV3) JSON() []byte {
	reqAuth, err := json.Marshal(s)
	if err != nil {
		// Return an empty structure
		reqAuth = []byte{'{', '}'}
	}
	return reqAuth
}

