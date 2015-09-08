package server

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/pics/web"
)

func Run() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)
	mux.HandleFunc("/add", Form)
	mux.HandleFunc("/upload", Upload)
	mux.HandleFunc("/all", All)
	web.Start(
		mux,
		viper.GetString("address"),
		viper.GetString("port"),
		viper.GetString("cert"),
		viper.GetString("key"),
	)
}
