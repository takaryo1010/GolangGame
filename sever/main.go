package main

import (
    "log"
    "net/http"
)

func main() {
    // ファイルサーバーをカレントディレクトリで起動
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    // サーバーをポート8080で起動
    log.Println("サーバーを開始します... ポート: 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
