package imgscale

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type Format struct {
	Prefix    string
	Height    uint
	Ratio     float64
	Thumbnail bool
}

type Config struct {
	Path    string
	Prefix  string
	Formats []*Format
	Exts    []string
	Comment string
}

var supportedExts = map[string]string{"jpg": "image/jpeg", "png": "image/png"}

type Handler struct {
	Config        *Config
	Formats       map[string]*Format
	Path          string
	regexp        *regexp.Regexp
	supportedExts map[string]string
}

func (h *Handler) Match(url string) (bool, *ImageInfo) {
	matches := h.regexp.FindStringSubmatch(url)

	if len(matches) == 0 {
		return false, nil
	}

	return true, h.GetImageInfo(matches[1], matches[2], matches[3])
}

func (h *Handler) GetContentType(ext string) string {
	return h.supportedExts[ext]
}

func (h *Handler) GetFormat(format string) *Format {
	return h.Formats[format]
}

func (h *Handler) GetImageInfo(format, filename, ext string) *ImageInfo {
	f := h.GetFormat(format)
	return &ImageInfo{fmt.Sprintf("%s/%s.%s", h.Config.Path, filename, ext), f.Height, f.Ratio, ext, f.Thumbnail, h.Config.Comment}
}

func (h *Handler) Serve(res http.ResponseWriter, req *http.Request, info *ImageInfo) {
	img, err := GetImage(info)
	defer img.Destroy()
	if err == nil {
		imgData := img.GetImageBlob()
		res.Header().Set("Content-Type", h.GetContentType(info.Ext))
		res.Header().Set("Content-Length", strconv.Itoa(len(imgData)))
		res.Write(imgData)
	} else {
		panic(err)
	}
}