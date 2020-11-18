package handlers

import (
	"encoding/xml"
	"time"
	//"net/http"
	"strconv"
	//m "github.com/Ulbora/Six910-ui/managers"
	//api "github.com/Ulbora/Six910API-Go"
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

//SiteMapValues SiteMapValues
type SiteMapValues struct {
	Domain        string
	ProductIDList *[]int64
}

//SiteMap SiteMap
type SiteMap struct {
	Xmlns   string   `xml:"xmlns,attr"`
	XMLName xml.Name `xml:"urlset"`
	URLList *[]SiteMapURL
}

//SiteMapURL SiteMapURL
type SiteMapURL struct {
	XMLName      xml.Name `xml:"url"`
	Loc          string   `xml:"loc"`
	LastModified string   `xml:"lastmod"`
	ChangeFreq   string   `xml:"changefreq"`
	Priority     string   `xml:"priority"`
}

func (h *Six910Handler) generateSiteMap(v *SiteMapValues) []byte {
	var sm SiteMap
	sm.Xmlns = "http://www.google.com/schemas/sitemap/0.9"
	var ulst []SiteMapURL
	today := time.Now()
	m := today.Month()
	d := today.Day()
	y := today.Year()
	lastMod := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

	var u SiteMapURL
	u.Loc = v.Domain
	u.LastModified = lastMod
	u.ChangeFreq = "daily"
	u.Priority = "1"
	ulst = append(ulst, u)

	for i, id := range *v.ProductIDList {
		if i < 49998 {
			var up SiteMapURL
			up.Loc = v.Domain + "/viewProduct/" + strconv.FormatInt(id, 10)
			up.LastModified = lastMod
			up.ChangeFreq = "weekly"
			up.Priority = "0.8"
			ulst = append(ulst, up)
		}
	}

	sm.URLList = &ulst

	out, _ := xml.MarshalIndent(sm, " ", "  ")
	// h.Log.Debug("site map \n", string(out))

	return out

}
