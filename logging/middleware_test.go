package logging

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	. "github.com/smartystreets/goconvey/convey"
)

type dummyLogFile struct {
	Log []byte
	Wg  *sync.WaitGroup
}

func (d *dummyLogFile) Write(b []byte) (n int, err error) {
	d.Log = b
	d.Wg.Done()
	return
}

type dummyBody struct {
	Test bool `json:"test"`
}

func TestAccessLogger(t *testing.T) {

	Convey("Given a request for access log", t, func() {

		r := gin.New()

		wg := &sync.WaitGroup{}
		wg.Add(1)
		d := &dummyLogFile{Wg: wg}
		r.Use(AccessLogger(d))
		r.GET("/test", func(c *gin.Context) {
			c.JSON(200, map[string]bool{"test": true})
		})

		req, _ := http.NewRequest("GET", "/test?test=true", nil)
		req.Header.Set("X-Forwarded-For", "127.0.0.1")
		req.Header.Set("User-Agent", "testAgent")
		req.Header.Set("Referer", "testReferer")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		wg.Wait()

		var al AccessLog
		err := json.Unmarshal(d.Log, &al)
		So(err, ShouldBeNil)
		t, err := time.Parse(time.RFC3339, al.Date)
		So(err, ShouldBeNil)

		So(al.ClientIP, ShouldEqual, req.Header.Get("X-Forwarded-For"))
		So(t, ShouldHappenBefore, time.Now())
		So(al.Method, ShouldEqual, req.Method)
		So(al.RequestURI, ShouldEqual, req.URL.RequestURI())
		So(al.Referer, ShouldEqual, req.Referer())
		So(al.HTTPVersion, ShouldEqual, req.Proto)
		So(al.Size, ShouldBeGreaterThan, 0)
		So(al.Status, ShouldEqual, 200)
		So(al.UserAgent, ShouldEqual, req.Header.Get("User-Agent"))
		So(al.Latency, ShouldBeGreaterThan, 0)
	})

}

func TestActivityLogger(t *testing.T) {

	Convey("Given a request for activity log", t, func() {

		r := gin.New()

		wg := &sync.WaitGroup{}
		d := &dummyLogFile{Wg: wg}
		r.Use(ActivityLogger(d, func(_ *gin.Context) (interface{}, error) {
			return "testUserID", nil
		}))

		Convey("Given GET request", func() {

			r.GET("/test", func(c *gin.Context) {
				c.JSON(200, map[string]bool{"test": true})
				wg.Done()
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			req.Header.Set("X-Forwarded-For", "127.0.0.1")
			req.Header.Set("User-Agent", "testAgent")

			wg.Add(1)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			wg.Wait()

			So(d.Log, ShouldBeNil)
		})

		Convey("Given POST request", func() {

			r.POST("/test", func(c *gin.Context) {
				c.JSON(200, map[string]bool{"test": true})
			})

			body := &dummyBody{true}
			b, err := json.Marshal(&body)
			So(err, ShouldBeNil)

			req, _ := http.NewRequest("POST", "/test?test=true", bytes.NewReader(b))
			req.Header.Set("X-Forwarded-For", "127.0.0.1")
			req.Header.Set("User-Agent", "testAgent")

			wg.Add(1)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			wg.Wait()

			var al ActivityLog
			err = json.Unmarshal(d.Log, &al)
			So(err, ShouldBeNil)
			t, err := time.Parse(time.RFC3339, al.Date)
			So(err, ShouldBeNil)

			So(al.ClientIP, ShouldEqual, req.Header.Get("X-Forwarded-For"))
			So(t, ShouldHappenBefore, time.Now())
			So(al.Method, ShouldEqual, req.Method)
			So(al.RequestURI, ShouldEqual, req.URL.RequestURI())
			So(al.Referer, ShouldBeEmpty)
			So(al.HTTPVersion, ShouldEqual, req.Proto)
			So(al.Size, ShouldBeGreaterThan, 0)
			So(al.Status, ShouldEqual, 200)
			So(al.UserAgent, ShouldEqual, req.Header.Get("User-Agent"))
			So(al.Latency, ShouldBeGreaterThan, 0)
			So(al.RequestBody["test"], ShouldBeTrue)
			So(al.Extra, ShouldEqual, "testUserID")
		})

	})

}
