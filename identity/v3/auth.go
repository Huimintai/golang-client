// auth - Authentication interface
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
	"errors"
	"time"

	"git.openstack.org/openstack/golang-client.git/openstack"
)

// AuthRef is the returned authentication object, maybe v2 or v3
type AuthRef interface {
	GetToken() string
	GetExpiration() time.Time
}

// AuthOpts is the set of credentials used to authenticate to OpenStack
type AuthOpts struct {
	// AuthUrl is always required
	AuthUrl string

	// Domain is ignored for v2 and required for v3 auth
	Domain string

	// Project is optional to get an unscoped token but required for
	// a scoped token, which is required to do pretty much everything
	// except list projects
	Project string

	// Username is required for password auth
	Username string

	// Password is required for password auth
	Password string

	// Token is required for Toekn auth
	Token string
}

func (s *AuthOpts) GetAuthType() (string, error) {
	var auth_type string
	if s.AuthUrl != "" && s.Token != "" {
		auth_type = "token"
	} else if s.Username != "" {
		auth_type = "password"
	}
	return auth_type, nil
}

// Basic auth call
// These args should be an interface??
func DoAuthRequest(authopts AuthOpts) (AuthRef, string, error) {
	var auth = AuthToken{}
        var token = ""

	// Assume passwordv2 for now
	auth_mod, err := NewUserPassV3(authopts)
	if err != nil {
		err = errors.New("Failed to get auth options")
		return nil, "", err
	}

	_, token, err = openstack.PostJSON(authopts.AuthUrl + "/auth/tokens", nil, nil, &auth_mod, &auth)
	if err != nil {
		return nil, "", err
	}

	return auth, token, nil
}
