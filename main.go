package main

import (
	_ "core/app"
	"core/util/http/responder"

	"net/http"

	"github.com/guregu/kami"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

// ServeFiles is function to serving static multi-files.
// See `"github.com/julienschmidt/httprouter".ServeFiles()`.
func ServeFiles(path string, dir http.FileSystem) {
	kami.Handler().(*httprouter.Router).ServeFiles(path, dir)
}

// ServeFile is function to serving static single-file.
func ServeFile(path, file string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})
}

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	responder.Respond(w, "Hello, "+kami.Param(ctx, "name")+"!")
}

func init() {
	http.Handle("/", kami.Handler())
	kami.Get("/hello/:name", hello)
	ServeFiles("/files/*filepath", http.Dir("static-files"))
	ServeFile("/favicon.ico", "static-files/favicon.ico")
}

/*
TODO: カスタムエラーページ
TODO: appengine のマルチモジュール対応
TODO: マルチモジュール名の別ファイル化と環境別 app.yaml の共通化
TODO: context とサービス層での循環参照問題
TODO: 汎用レスポンダーとテンプレート処理
*/
