//Package usersrv ...
package usersrv

import (
	"bytes"
	"encoding/json"
	"net/http"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//User user
type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Enabled      bool   `json:"enabled"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	RoleID       int64  `json:"roleId"`
	ClientID     int64  `json:"clientId"`
}

//UpdateUser interface
type UpdateUser interface {
	GetType() string
}

//UserPW user
type UserPW struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

//GetType type
func (u *UserPW) GetType() string {
	return "PW"
}

//UserDis user
type UserDis struct {
	Username string `json:"username"`
	Enabled  bool   `json:"enabled"`
	ClientID int64  `json:"clientId"`
}

//GetType type
func (u *UserDis) GetType() string {
	return "DIS"
}

//UserInfo user
type UserInfo struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	RoleID       int64  `json:"roleId"`
	ClientID     int64  `json:"clientId"`
}

//GetType type
func (u *UserInfo) GetType() string {
	return "INFO"
}

//Role user role
type Role struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

//UserResponse resp
type UserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//Oauth2UserService Oauth2UserService
type Oauth2UserService struct {
	Token    string
	ClientID string
	//APIKey   string
	//UserID string
	//Hashed string
	Host     string
	UserHost string
	Proxy    px.Proxy
	Log      *lg.Logger
}

//UserService UserService
type UserService interface {
	UpdateUser(user UpdateUser) *UserResponse
	GetUser(username string, clientID string) (*User, int)
	SetToken(token string)
}

//UpdateUser update
func (u *Oauth2UserService) UpdateUser(user UpdateUser) *UserResponse {
	var rtn = new(UserResponse)
	var upURL = u.UserHost + "/rs/user/update"
	aJSON, err := json.Marshal(user)
	u.Log.Debug("update user: ", err)
	if err == nil {
		req, rErr := http.NewRequest("PUT", upURL, bytes.NewBuffer(aJSON))
		u.Log.Debug("update user req: ", rErr)
		if rErr == nil {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+u.Token)
			req.Header.Set("clientId", u.ClientID)
			//req.Header.Set("apiKey", u.APIKey)
			_, code := u.Proxy.Do(req, &rtn)
			rtn.Code = code
		}
	}
	return rtn
}

// GetUser get
func (u *Oauth2UserService) GetUser(username string, clientID string) (*User, int) {
	var rtn = new(User)
	var code int
	var gURL = u.UserHost + "/rs/user/get/" + username + "/" + clientID
	req, rErr := http.NewRequest("GET", gURL, nil)
	u.Log.Debug("get user req: ", rErr)
	if rErr == nil {
		req.Header.Set("clientId", u.ClientID)
		req.Header.Set("Authorization", "Bearer "+u.Token)
		//req.Header.Set("apiKey", u.APIKey)
		_, code = u.Proxy.Do(req, &rtn)
	}
	return rtn, code
}

//SetToken SetToken
func (u *Oauth2UserService) SetToken(token string) {
	u.Token = token
}

//GetNew GetNew
func (u *Oauth2UserService) GetNew() UserService {
	return u
}
