// auth-token - Token Authentication
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
	"time"
)

// Identity Response Types

type AuthToken struct {
	Token Token `json:"token"`
}

type AuthDomain struct {
	ID string   `json:"id"`
        Name string `json:"name"`
}

type AuthUser struct {
        ID string   `json:"id"`
	Name string `json:"name"`
	Domain AuthDomain `json:"domain"`
}

type Token struct {
        Methods []string  `json:"methods"`
	Extra   interface{} `json:"extra"`
	Expires time.Time `json:"expires_at"`
	Issued  time.Time `json:"issued_at"`
	User    AuthUser  `json:"user"`
        AuditIDs []string `json:"audit_ids"`
        
}

func (s AuthToken) GetToken() string {
	return s.Token.AuditIDs[0]
}

func (s AuthToken) GetExpiration() time.Time {
	return s.Token.Expires
}
