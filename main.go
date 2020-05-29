package main

/*
 Six910-ui is a shopping cart and E-commerce system.
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

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// just the start of Six910-ui Server
	// This is the storefront for Six910-ui.
	fmt.Println("Six910 (six nine ten) UI is running on port 3001!")
	router := mux.NewRouter()
	http.ListenAndServe(":3001", router)
}

// go mod init github.com/Ulbora/Six910-ui
