package handlers

import (
	"html/template"
	//"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
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

//LoadTemplate LoadTemplate
func (h *Six910Handler) LoadTemplate() {
	h.ActiveTemplateName = h.TemplateService.GetActiveTemplateName()
	h.Log.Debug("ActiveTemplateName: ", h.ActiveTemplateName)

	//var tperr error
	// tp, tperr := template.ParseFiles("./static/templates/"+h.ActiveTemplateName+"/index.html", "./static/templates/"+h.ActiveTemplateName+"/header.html",
	// 	"./static/templates/"+h.ActiveTemplateName+"/footer.html", "./static/templates/"+h.ActiveTemplateName+"/navbar.html",
	// 	"./static/templates/"+h.ActiveTemplateName+"/contact.html")

	tp, tperr := template.ParseFiles(
		h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/index.html",
		h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/header.html",
		h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/navBar.html",
		h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/product.html",
	)

	//h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/footer.html", h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/navbar.html",
	//h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/contact.html", h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/viewContent.html",
	//h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/blogs.html", h.ActiveTemplateLocation+"/"+h.ActiveTemplateName+"/archivedBlogs.html",

	h.Log.Debug("template error: ", tperr)
	h.Templates = template.Must(tp, tperr)
}
