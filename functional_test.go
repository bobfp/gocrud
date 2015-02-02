package main

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWebServer(t *testing.T) {
	Convey("Joe GETs /", t, func() {
		response, _ := http.Get("http://localhost:8888/")
		defer response.Body.Close()

		Convey("Joe receives a status 200", func() {
			So(response.Status, ShouldEqual, "200 OK")
		})
	})
}
