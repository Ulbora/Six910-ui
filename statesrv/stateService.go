package statesrv

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2020 Ken Williamson
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

//Template Service

import (
	"encoding/json"
	//lg "github.com/Ulbora/Level_Logger"
	//ds "github.com/Ulbora/json-datastore"
)

//GetStateList GetStateList
func (c *Six910StateService) GetStateList(name string) *[]State {
	var rtn []State
	em := c.StateStore.Read(name)
	if *em != nil {
		json.Unmarshal(*em, &rtn)
	}
	return &rtn
}
