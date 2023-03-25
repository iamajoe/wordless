package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/iamajoe/wordless/config"
	"github.com/iamajoe/wordless/infrastructure/redis"
)

func Test_reqCreateUrl(t *testing.T) {
	c, err := config.Get(os.Getenv)
	if err != nil {
		t.Fatal(err)
	}

	repos, err := redis.InitRepos(c.RedisHost, c.RedisSecret)
	if err != nil {
		t.Fatal(err)
	}

	type response struct {
		Ok   bool   `json:"ok"`
		Code int    `json:"code"`
		Data string `json:"data,omitempty"`
		Err  string `json:"err,omitempty"`
	}

	type args struct {
		urlValue string
		body     []byte
	}
	type testStruct struct {
		name     string
		args     args
		wantCode int
		wantBody response
		wantUser bool
	}

	tests := []testStruct{
		func() testStruct {
			urlValue := fmt.Sprintf("%d", rand.Intn(10000))
			body := []byte(fmt.Sprintf(`{"url":"%s"}`, urlValue))

			return testStruct{
				name:     "runs",
				args:     args{urlValue, body},
				wantCode: http.StatusOK,
				wantBody: response{
					Ok:   true,
					Code: http.StatusOK,
					Err:  "",
					Data: "something...",
				},
				wantUser: true,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/", bytes.NewBuffer(tt.args.body))
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(reqCreateUrl(repos))
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("wrong status code: got %v want %v", rec.Code, tt.wantCode)
			}

			var res response
			err = json.NewDecoder(rec.Body).Decode(&res)
			if err != nil {
				t.Fatal(err)
			}

			if res.Ok != tt.wantBody.Ok || res.Err != tt.wantBody.Err || res.Code != tt.wantBody.Code {
				t.Errorf("body = %v, want %v", rec.Body.String(), tt.wantBody)
				return
			}

			if (len(res.Data) == 0 && len(tt.wantBody.Data) > 0) || (len(res.Data) > 0 && len(tt.wantBody.Data) == 0) {
				t.Errorf("data = %v, want %v", res.Data, tt.wantBody.Data)
			}

			urls, err := repos.GetUrl().GetByIDs([]string{res.Data})
			if err != nil {
				t.Fatal(err)
			}
			if len(urls) == 0 || len(urls[0].Value) == 0 || urls[0].Value != tt.args.urlValue {
				t.Fatal("url value coming is wrong")
			}
		})
	}
}
