// Copyright (c) 2014 Hewlett-Packard Development Company, L.P.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package user_test

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"git.openstack.org/openstack/golang-client.git/user/v2"
	"git.openstack.org/openstack/golang-client.git/openstack"
	"git.openstack.org/openstack/golang-client.git/testUtil"
)

var tokn = "eaaafd18-0fed-4b3a-81b4-663c99ec1cbb"

func TestListUsers(t *testing.T) {
	anon := func(userService *user.Service) {
		users, err := userService.Users()
		if err != nil {
			t.Error(err)
		}

		if len(users) != 2 {
			t.Error(errors.New("Incorrect number of users found"))
		}
		expectedUser := user.Response{
			Name:            "admin",
			UserName:        "admin",
			Email:           "aa@aa.aa",
			ID:              "bec3cab5-4722-40b9-a78a-3489218e22fe",
			Enabled:         false}
		// Verify first one matches expected values
		testUtil.Equals(t, expectedUser, users[0])
	}

	testUserServiceAction(t, "users", sampleUsersData, anon)
}

func TestGetUser(t *testing.T) {
	anon := func(userService *user.Service) {
		one_user, err := userService.GetUserByName("admin")
		if err != nil {
			t.Error(err)
		}

		expectedUser := user.Response{
			Name:            "admin",
			UserName:        "admin",
			Email:           "aa@aa.aa",
			ID:              "bec3cab5-4722-40b9-a78a-3489218e22fe",
			Enabled:         false}
		// Verify first one matches expected values
		testUtil.Equals(t, expectedUser, one_user)
	}

	testUserServiceAction(t, "users?name=admin", sampleUserData, anon)
}

func testUserServiceAction(t *testing.T, uriEndsWith string, testData string, userServiceAction func(*user.Service)) {
    anon := func(req *http.Request) {
			reqURL := req.URL.String()
			if !strings.HasSuffix(reqURL, uriEndsWith) {
				t.Error(errors.New("Incorrect url created, expected:" + uriEndsWith + " at the end, actual url:" + reqURL))
			}
	}
	apiServer := testUtil.CreateGetJSONTestRequestServer(t, tokn, testData, anon)
	defer apiServer.Close()

	auth := openstack.AuthToken{
			Access: openstack.AccessType{
					Token: openstack.Token{
						    ID: tokn,
					},
			},
	}
	sess, _ := openstack.NewSession(http.DefaultClient, auth, nil)
	userService := user.Service{
					Session: *sess,
					URL:     apiServer.URL,
	}
	userServiceAction(&userService)
}

var sampleUsersData = `{
    "users":[
	    {
		    "name":"admin",
		    "username":"admin",
			"email":"aa@aa.aa",
			"id":"bec3cab5-4722-40b9-a78a-3489218e22fe",
			"enaled":false
		},
	    {
		    "name":"demo",
		    "username":"demo",
			"email":"aa@aa.aa",
			"id":"bec3cab5-4722-40b9-a78a-3489218e22fa",
			"enaled":false
		}
	]
}`


var sampleUserData = `{
	"name":"admin",
	"username":"admin",
	"email":"aa@aa.aa",
	"id":"bec3cab5-4722-40b9-a78a-3489218e22fe",
	"enaled":false
}`
