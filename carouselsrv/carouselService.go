package carouselsrv

import (
	"encoding/json"
)

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

//Carousel Carousel
type Carousel struct {
	Name    string
	Enabled bool
	Image1  string
	Image2  string
	Image3  string
}

//GetCarousel GetCarousel
func (c *Six910CarouselService) GetCarousel(name string) (bool, *Carousel) {
	var rtn Carousel
	var suc bool
	ec := c.Store.Read(name)
	if *ec != nil {
		err := json.Unmarshal(*ec, &rtn)
		if err == nil {
			suc = true
		}
	}
	return suc, &rtn
}

//UpdateCarousel UpdateCarousel
func (c *Six910CarouselService) UpdateCarousel(car *Carousel) bool {
	var rtn bool
	ec := c.Store.Read(car.Name)
	if *ec != nil {
		var cc Carousel
		err := json.Unmarshal(*ec, &cc)
		c.Log.Debug("found Carousel in update: ", cc)
		if err == nil {
			cc.Enabled = car.Enabled
			cc.Image1 = car.Image1
			cc.Image2 = car.Image2
			cc.Image3 = car.Image3
			suc := c.Store.Save(car.Name, cc)
			rtn = suc
		}
	}
	return rtn
}

//GetNew GetNew
func (c *Six910CarouselService) GetNew() CarouselService {
	return c
}
