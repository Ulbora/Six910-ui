package tmptsrv

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
	"encoding/json"
	"path/filepath"
	"sync"
)

//Template template
type Template struct {
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	ScreenShot string `json:"screenShot"`
}

var tmu sync.Mutex

//AddTemplateFile AddTemplateFile
func (c *Six910TemplateService) AddTemplateFile(name string, originalFileName string, fileData []byte) bool {
	tmu.Lock()
	defer tmu.Unlock()
	var rtn bool
	c.Log.Debug("template file name in add: ", name)
	var tpl TemplateFile
	tpl.FileData = fileData
	tpl.Name = name
	tpl.OriginalFileName = originalFileName
	//c.Log.Debug("tpl in add: ", tpl)
	rtn = c.ExtractFile(&tpl)
	//var templateName = c.TemplateFilePath + string(filepath.Separator) + name
	//c.Log.Debug("template complete file name in add: ", templateName)
	//err := ioutil.WriteFile(templateName, fileData, 0644)
	return rtn
}

//AddTemplate AddTemplate
func (c *Six910TemplateService) AddTemplate(tpl *Template) bool {
	var rtn bool
	c.Log.Debug("tpl template in add: ", *tpl)
	if tpl != nil && tpl.Name != "" {
		tpl.Active = false
		exis := c.TemplateStore.Read(tpl.Name)
		c.Log.Debug("existing template in add: ", *exis)
		if *exis == nil {
			rtn = c.TemplateStore.Save(tpl.Name, tpl)
			c.Log.Debug("template add suc: ", rtn)
		}
	}
	return rtn
}

//GetActiveTemplateName GetActiveTemplateName
func (c *Six910TemplateService) GetActiveTemplateName() string {
	var rtn string
	res := c.TemplateStore.ReadAll()
	//c.Log.Debug("tpls template get active: ", *res)
	for r := range *res {
		var t Template
		err := json.Unmarshal((*res)[r], &t)
		c.Log.Debug("found template in list: ", t)
		if err == nil && t.Active {
			rtn = t.Name
			break
		}
	}
	return rtn
}

//GetTemplateList GetTemplateList
func (c *Six910TemplateService) GetTemplateList() *[]Template {
	var rtn []Template
	res := c.TemplateStore.ReadAll()
	//c.Log.Debug("tpls template get list: ", *res)
	for r := range *res {
		var t Template
		err := json.Unmarshal((*res)[r], &t)
		c.Log.Debug("found template in list: ", t)
		if err == nil {
			t.ScreenShot = ".." + string(filepath.Separator) + "templates" + string(filepath.Separator) + t.Name + string(filepath.Separator) + "screenshot.png"
			rtn = append(rtn, t)
		}
	}
	return &rtn
}

//ActivateTemplate ActivateTemplate
func (c *Six910TemplateService) ActivateTemplate(name string) bool {
	var rtn bool
	res := c.TemplateStore.ReadAll()
	//c.Log.Debug("tpls template get list: ", *res)
	for r := range *res {
		var t Template
		err := json.Unmarshal((*res)[r], &t)
		c.Log.Debug("found template in list in activate: ", t)
		if err == nil {
			t.Active = false
			c.TemplateStore.Save(t.Name, t)
		}
	}
	etpl := c.TemplateStore.Read(name)
	if *etpl != nil {
		var t Template
		err := json.Unmarshal(*etpl, &t)
		c.Log.Debug("found template in activate: ", t)
		if err == nil {
			t.Active = true
			rtn = c.TemplateStore.Save(t.Name, t)
			c.Log.Debug("template activate suc: ", rtn)
		}
	}
	return rtn
}

//DeleteTemplate DeleteTemplate
func (c *Six910TemplateService) DeleteTemplate(name string) bool {
	var rtn bool
	detpl := c.TemplateStore.Read(name)
	if *detpl != nil {
		var t Template
		err := json.Unmarshal(*detpl, &t)
		c.Log.Debug("found template in delete: ", t)
		if err == nil {
			if !t.Active {
				rtn = c.TemplateStore.Delete(name)
				c.Log.Debug("template delete suc: ", rtn)
			}
		}
	}
	return rtn
}
