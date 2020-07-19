package contsrv

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

import (
	"sync"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

//Service Service
type Service interface {
	AddContent(content *Content) *Response
	UpdateContent(content *Content) *Response
	GetContent(name string) (bool, *Content)
	GetContentList(published bool) *[]Content
	DeleteContent(name string) *Response

	SendCaptchaCall(cap Captcha) *CaptchaResponse

	SaveHits()

	HitCheck()
}

//CmsService service
type CmsService struct {
	Store              ds.JSONDatastore
	ContentStorePath   string
	Log                *lg.Logger
	CaptchaHost        string
	MockCaptcha        bool
	MockCaptchaSuccess bool
	MockCaptchaCode    int
	HitTotal           int
	ContentHits        map[string]int64
	HitLimit           int
	hitmu              sync.Mutex
}

//GetNew GetNew
func (c *CmsService) GetNew() Service {
	var cs Service
	c.ContentHits = make(map[string]int64)
	cs = c
	return cs
}
