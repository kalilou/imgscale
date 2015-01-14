package imgscale

import (
	"github.com/gographics/imagick/imagick"
	"net/http"
)

/*
	ImageProvider implements image fetching and provide the image source
	for the handler to serve to the request. There are 2 providers available
	One for local filesystem imageProviderFile and one for remote file
	using http GET imageProviderHTTP
*/
type ImageProvider interface {
	Fetch(filename string) (*imagick.MagickWand, error)
}

/*
	Validator implements the validation of the filename
	The filename can identified as
	baseurl/<image prefix>/<format prefix>/<filename>
	An example is
	http://127.0.0.1:8080/img/original/kth.jpg
	and filename is kth.jpg
*/
type Validator interface {
	// Name of the image, ie everything after "<prefix>/<format>/"
	Validate(filename string) bool
}

/* 
	Handler implements http.Handler so it can use for many frameworks available
	Handler.ServeHTTP can use similar to http.HandleFunc in case frameworks support
	only this.
	
	Beside that Handler has 2 methods, Handler.SetValidator for setting own validation of the filename
	Handler.SetImageProvider is suitable when you have customized image provider. Default is
	imageProviderFile
*/
type Handler interface {
	// http.HandleFunc
	ServeHTTP(res http.ResponseWriter, req *http.Request)
	// For setting own image provider, such as remote storage
	SetImageProvider(provider ImageProvider)
	// For setting validator of filename/name of the image
	SetValidator(validator Validator)
}

