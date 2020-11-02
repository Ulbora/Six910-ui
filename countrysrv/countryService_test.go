package countrysrv

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
	"testing"
)

func TestSix910CountryService_GetCountryList(t *testing.T) {

	var st Six910CountryService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	st.CountryStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	st.Log = &l

	s := st.GetNew()

	stlst := s.GetCountryList("countries")
	fmt.Println("Country list: ", *stlst)
	if len(*stlst) != 1 {
		t.Fail()
	}

}
