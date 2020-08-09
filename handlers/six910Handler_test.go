package handlers

import (
	"net/http"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

func TestSix910Handler_getSession(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	ses, suc := h.getSession(r)
	if ses == nil || !suc {
		t.Fail()
	}
}
