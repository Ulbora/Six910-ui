//Package usersrv ...
package usersrv

import (
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

//MockOauth2UserService MockOauth2UserService
type MockOauth2UserService struct {
	Token    string
	ClientID string
	//APIKey   string
	//UserID string
	//Hashed string
	Host     string
	UserHost string
	Proxy    px.Proxy
	Log      *lg.Logger

	MockUpdateUserResponse *UserResponse
	MockUser         *User
	MockUserCode     int
}

//UpdateUser UpdateUser
func (u *MockOauth2UserService) UpdateUser(user UpdateUser) *UserResponse {
	return u.MockUpdateUserResponse
}

//GetUser GetUser
func (u *MockOauth2UserService) GetUser(username string, clientID string) (*User, int) {
	return u.MockUser, u.MockUserCode
}

//SetToken SetToken
func (u *MockOauth2UserService) SetToken(token string) {
	u.Token = token
}

//GetNew GetNew
func (u *MockOauth2UserService) GetNew() UserService {
	return u
}
