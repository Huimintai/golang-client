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

package keystonev3

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"git.openstack.org/openstack/golang-client.git/openstack"
	"git.openstack.org/openstack/golang-client.git/util"
)

type Service struct {
	Session openstack.Session
	Client  http.Client
	URL     string
}

type OSLinks struct {
	Next		string `json:"next"`
	Previous	string `json:"previous"`
	Self		string `json:"self"`
}

type OSUsers struct {
	DomainID        string `json:"domain_id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	ID              string `json:"id"`
	Enabled         bool   `json:"enabled"`
        Links		OSLinks `json:"links"`
}

type Response struct {
	Links           OSLinks `json:"links"`
	Users		[]OSUsers `json:"users"`
}

func (userService Service) Users() ([]OSUsers, error) {
	userContainer := Response{}
	err := userService.queryUsers(&userContainer, "")

	return userContainer.Users, err
}

func (userService Service) GetUserByName(name string) ([]OSUsers, error) {
	userContainer := Response{}
	err := userService.queryUsers(&userContainer, name)

	return userContainer.Users, err
}

func (userService Service) queryUsers(userResponseContainer interface{}, name string) error {
	urlPostFix := "/users"

	reqURL, err := buildQueryURL(userService, name, urlPostFix)
	if err != nil {
		return err
	}

	var headers http.Header = http.Header{}
	headers.Set("Accept", "application/json")
	resp, err := userService.Session.Get(reqURL.String(), nil, &headers)
	if err != nil {
		return err
	}

	err = util.CheckHTTPResponseStatusCode(resp)
	if err != nil {
		return err
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("aaa")
	}
	if err = json.Unmarshal(rbody, &userResponseContainer); err != nil {
		return err
	}
	return nil
}

func buildQueryURL(userService Service, name string, userPartialURL string) (*url.URL, error) {
	reqURL, err := url.Parse(userService.URL)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	if name != "" {
		values.Set("name", name)
	}
	if len(values) > 0 {
		reqURL.RawQuery = values.Encode()
	}
	reqURL.Path += userPartialURL

	return reqURL, nil
}

