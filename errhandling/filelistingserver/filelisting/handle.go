package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, error := os.Open(path)
	if error != nil {
		return error
	}
	defer file.Close()

	all, error := ioutil.ReadAll(file)
	if error != nil {
		return error
	}
	writer.Write(all)
	return nil
}
