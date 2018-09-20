package filelisting

import (
	"github.com/golang/go/src/pkg/strings"
	"io/ioutil"
	"net/http"
	"os"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("Path must start with" + prefix)
	}

	path := request.URL.Path[len(prefix):]
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
