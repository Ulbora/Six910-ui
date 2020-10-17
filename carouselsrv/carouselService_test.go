package carouselsrv

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

func TestSix910CarouselService_GetCarousel(t *testing.T) {

	var cs Six910CarouselService
	var cds ds.DataStore
	cds.Path = "./testFiles"
	cs.Store = cds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	suc, car := s.GetCarousel("testCar")
	fmt.Println("found carousel", *car)
	if !suc {
		t.Fail()
	}
}

func TestSix910CarouselService_UpdateCarousel(t *testing.T) {

	var cs Six910CarouselService
	var cds ds.DataStore
	cds.Path = "./testFiles"
	cs.Store = cds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	var car Carousel
	car.Name = "testCar"
	car.Enabled = true
	car.Image1 = "/img1"
	car.Image2 = "/img2"
	car.Image3 = "/img3"

	suc := s.UpdateCarousel(&car)
	if !suc {
		t.Fail()
	}
}
