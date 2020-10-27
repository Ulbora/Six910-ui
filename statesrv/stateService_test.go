package statesrv

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
	"testing"
)

func TestSix910StateService_GetStateList(t *testing.T) {

	var st Six910StateService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	st.StateStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	st.Log = &l

	s := st.GetNew()

	stlst := s.GetStateList("states")
	fmt.Println("states: ", *stlst)
	if len(*stlst) != 51 {
		t.Fail()
	}

}
