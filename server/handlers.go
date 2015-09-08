package server

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/pics/log"
)

func Form(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<html>
<head>
    <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
      <input type="file" name="file" />
      <input type="submit" value="upload" />
</form>
</body>
</html>`))
}

func Upload(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(32 << 20)

	uploadedFile, handler, err := req.FormFile("file")
	if log.PrintError("pics upload", err) != nil {
		fmt.Fprintf(w, "ERROR")
		return
	}
	defer uploadedFile.Close()

	err = os.MkdirAll(viper.GetString("upload"), 0700)
	if log.PrintError("pics upload", err) != nil {
		fmt.Fprintf(w, "ERROR")
		return
	}
	serverFile, err := os.OpenFile(
		viper.GetString("upload")+handler.Filename,
		os.O_WRONLY|os.O_CREATE,
		0600,
	)
	if log.PrintError("pics upload", err) != nil {
		fmt.Fprintf(w, "ERROR")
		return
	}
	defer serverFile.Close()

	if !Check(uploadedFile) {
		fmt.Fprintf(w, "ERROR")
		return
	}
	// io.Copy(serverFile, uploadedFile)
	fmt.Fprintf(w, "OK")
}

func All(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  files, _ := ioutil.ReadDir("static/pics")
	// Subtract 1 beacuse index starts at 0
	last := len(files) - 1
	fmt.Fprintf(w, "[")
  for index, filename := range files {
		fmt.Fprintf(w, "%q", filename.Name())
		if index < last {
			fmt.Fprintf(w, ", ")
		}
  }
	fmt.Fprintf(w, "]")
}
