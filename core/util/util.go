package util

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"appengine"
)

const (
	EnvDevelopment = "development" // 開発環境
	EnvProduction  = "production"  // 本番環境
	EnvStaging     = "staging"     // ステージング環境
)

func Log(r *http.Request, format string, v ...interface{}) {
	if IsDebug() {
		defer log.Printf("[INFO] "+format, v...)
	} else {
		defer appengine.NewContext(r).Infof(format, v...)
	}
}

func ErrorLog(r *http.Request, format string, v ...interface{}) {
	if IsDebug() {
		defer log.Printf("[ERR] "+format, v...)
	} else {
		defer appengine.NewContext(r).Errorf(format, v...)
	}
}

func Dump(r *http.Request, val ...interface{}) {
	for _, v := range val {
		Log(r, "%#v", v)
	}
}

func IsDebug() bool {
	return appengine.IsDevAppServer()
}

func Module(r *http.Request) string {
	return appengine.ModuleName(appengine.NewContext(r))
}

func Version(r *http.Request) string {
	return strings.Split(appengine.VersionID(appengine.NewContext(r)), ".")[0]
}

func Env(r *http.Request) string {
	if IsDebug() {
		return EnvDevelopment
	}
	if strings.HasSuffix(appengine.AppID(appengine.NewContext(r)), "staging") {
		return EnvStaging
	}
	return EnvProduction
}

func IsDevelopment(r *http.Request) bool {
	return Env(r) == EnvDevelopment
}

func IsStaging(r *http.Request) bool {
	return Env(r) == EnvStaging
}

func IsProduction(r *http.Request) bool {
	return Env(r) == EnvProduction
}

func RemoteAddr(r *http.Request) string {
	return r.RemoteAddr
}

func Method(r *http.Request) string {
	return strings.ToUpper(r.Method)
}

// SSL/TLS アクセスかどうかを簡易判別する (証明書が正しいかまでは確認しない)
func IsSecure(r *http.Request) bool {
	return r.TLS != nil
}

func ServerName(r *http.Request) string {
	return appengine.ServerSoftware()
}

func RequestServerName(r *http.Request) string {
	if IsDebug() {
		return appengine.DefaultVersionHostname(appengine.NewContext(r))
	}

	if u, err := url.Parse(r.URL.String()); err != nil {
		return ""
	} else {
		return u.Host
	}
}

func RequestURI(r *http.Request) string {
	return r.RequestURI
	//return r.URL.String()
}

func RequestBody(r *http.Request) *[]byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.Bytes()
	return &body
}

func Info(r *http.Request) *map[string]interface{} {
	ev := map[string]interface{}{}
	for _, e := range os.Environ() {
		p := strings.Split(e, "=")
		ev[p[0]] = p[1]
	}
	h := map[string]interface{}{}
	for k, v := range r.Header {
		h[k] = v[0]
	}
	ac := appengine.NewContext(r)
	return &map[string]interface{}{
		"debug":                 IsDebug(),
		"environment":           Env(r),
		"environment_variables": ev,
		"server":                ServerName(r),
		"host":                  RequestServerName(r),
		"remote_addr":           RemoteAddr(r),
		"method":                Method(r),
		"secure":                IsSecure(r),
		"request_uri":           RequestURI(r),
		"body":                  RequestBody(r),
		"data_center":           appengine.Datacenter(),
		"app_id":                appengine.AppID(ac),
		"module":                Module(r),
		"version_id":            Version(r),
		"version_id_full":       appengine.VersionID(ac),
		"request_id":            appengine.RequestID(ac),
		"request_headers":       h,
	}
}
