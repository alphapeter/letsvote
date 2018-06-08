package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"mime"
	"net/http"
	"strings"
)

type EmbeddedHandler struct {
	data            map[string][]byte
	defaultResource string
}

func CreateHandler(data map[string][]byte, defaultResource string) EmbeddedHandler {
	return EmbeddedHandler{
		data:            data,
		defaultResource: defaultResource,
	}
}

func (h *EmbeddedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	resource := strings.TrimLeft(r.RequestURI, "/")
	if len(resource) == 0 {
		resource = h.defaultResource
	}
	acceptsGzip := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")

	d, found := getData(h.data, resource, acceptsGzip)

	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	if acceptsGzip {
		w.Header().Set("Content-Encoding", "gzip")
	}

	contentType := getContentType(resource)
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(d)
	return
}

func getData(data map[string][]byte, resource string, acceptsGzip bool) ([]byte, bool) {
	_, found := data[resource]
	if !found {
		resource = strings.TrimRight(resource, "/") + "/index.html"
	}

	d, found := data[resource]

	if found && acceptsGzip && !strings.HasSuffix(strings.ToLower(resource), ".gz") {
		gzResource := resource + ".gz"
		if _, found := data[gzResource]; !found {
			var b bytes.Buffer
			w, _ := gzip.NewWriterLevel(&b, gzip.BestCompression)
			w.Write(d)
			w.Close()
			data[gzResource] = b.Bytes()
			fmt.Println("compress file", gzResource)
		}
		d, found := data[gzResource]
		return d, found
	}

	return d, found
}

func getContentType(resourceName string) string {
	lastDot := strings.LastIndex(resourceName, ".")
	if lastDot > -1 {
		extension := resourceName[lastDot:]
		return mime.TypeByExtension(strings.ToLower(extension))
	}
	return mime.TypeByExtension("application/octet-stream")
}
