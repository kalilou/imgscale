## imgscale

Middleware/handler for scaling image in golang. Use for serving images in different formats. Can use as middleware which is compitable with http.Handler

* Crop & Scale
* Autorotate
* Watermark


Warning: image processing is very resource consuming. If you use this in production you should put a cache server in front, such as Varnish.

## Dependencies

You need to install ImageMagick first.

### MacPorts
	
	sudo port install ImageMagick

### Homebrew

	brew install imagemagick
	
### Ubuntu/Debian

	sudo apt-get install libmagickwand-dev

### ImageMagick binding

	go get github.com/vanng822/imgscale/imagick

## Install 

	go get github.com/vanng822/imgscale/imgscale


## Example

	// Negroni
	app := negroni.New()
	handler := imgscale.Configure("./config/formats.json")
	// Free C pointers and terminate MagickWand environment
	defer handler.Cleanup()
	n.UseHandler(handler)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	// Martini
	app := martini.Classic()
	handler := imgscale.Configure("./config/formats.json")
	defer handler.Cleanup()
	app.Use(handler.ServeHTTP)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), app)

	// http.Handler
	handler := imgscale.Configure("./config/formats.json")
	defer handler.Cleanup()
	http.Handle("/", handler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), nil)

## GoDoc

[![GoDoc](https://godoc.org/github.com/vanng822/imgscale/imgscale?status.svg)](https://godoc.org/github.com/vanng822/imgscale/imgscale)


## Try it out

### checkout
	
	git clone https://github.com/vanng822/imgscale.git
	

### install dependencies

	go get github.com/gographics/imagick/imagick
	go get github.com/vanng822/imgscale/imgscale
	go get github.com/go-martini/martini
	go get github.com/codegangsta/negroni
	
	
### run application

	go run example/all.go

### browse it
	
	http://127.0.0.1:8080/img/100x100/kth.jpg
	http://127.0.0.1:8081/img/100x100/http://127.0.0.1:8080/img/original/kth.jpg
	http://127.0.0.1:8082/img/100x100/kth.jpg
	