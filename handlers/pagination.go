package handlers

import (
	"strconv"
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

//Pagination Pagination
type Pagination struct {
	PrevLink string
	NextLink string
	Pages    *[]Pageinate
}

//Pageinate Pageinate
type Pageinate struct {
	PageCount int
	Active    string
	Link      string
	First     bool
	Last      bool
}

func (h *Six910Handler) doPagination(vpstart int64, recLen int, pageSize int, linkBase string) *Pagination {
	var pagination Pagination
	pstrt := 0
	if vpstart > 0 {
		pstrt = int(vpstart) - pageSize
	}
	pend := pstrt + pageSize
	pagination.PrevLink = linkBase + "/" + strconv.Itoa(pstrt) + "/" + strconv.Itoa(pend)

	h.Log.Debug("pagination.PrevLink ", pagination.PrevLink)

	nstrt := int(vpstart)
	h.Log.Debug("nstrt ", nstrt)
	h.Log.Debug("recLen ", recLen)
	h.Log.Debug("pageSize ", pageSize)

	if recLen < pageSize {
		nstrt = int(vpstart) - pageSize
		if nstrt < 0 {
			nstrt = 0
		}
	} else {
		nstrt = int(vpstart) + pageSize
	}

	nend := nstrt + pageSize
	pagination.NextLink = linkBase + "/" + strconv.Itoa(nstrt) + "/" + strconv.Itoa(nend)

	h.Log.Debug("pagination.NextLink ", pagination.NextLink)

	var plst []Pageinate
	pcount := (int(vpstart) / pageSize) + 1
	h.Log.Debug("pcount ", pcount)
	for i := 1; i < pcount+2; i++ {
		var pn Pageinate
		pn.PageCount = i
		strt := (i - 1) * pageSize
		end := i * pageSize

		pn.Link = linkBase + "/" + strconv.Itoa(strt) + "/" + strconv.Itoa(end)

		h.Log.Debug("pn.Link ", pn.Link)

		h.Log.Debug("i ", i)
		h.Log.Debug("pcount ", pcount)
		if i == pcount {
			pn.Active = "active"
		}
		if i == 1 {
			pn.First = true
		} else if i == pcount+1 {
			pn.Last = true
		}

		plst = append(plst, pn)

		if recLen < pageSize && i == pcount {
			break
		}

	}
	h.Log.Debug("plst ", plst)
	pagination.Pages = &plst
	return &pagination
}
