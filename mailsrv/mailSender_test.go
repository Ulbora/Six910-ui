package mailsrv

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
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ml "github.com/Ulbora/go-mail-sender"
)

func TestCmsService_SendMail(t *testing.T) {
	var ci Six910MailService

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ci.Log = &l

	var ms ml.MockSecureSender
	ms.MockSuccess = true
	ci.MailSender = ms.GetNew()

	s := ci.GetNew()

	var m ml.Mailer
	m.Subject = "test"
	m.Body = "this is a test"
	m.Recipients = []string{"tester@tester.com"}
	m.SenderAddress = "somedude@sender.com"

	suc := s.SendMail(&m)
	if !suc {
		t.Fail()
	}

}
