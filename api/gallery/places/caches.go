package places

import "os"

func caches() (string, string, string) {
	cachePath := os.Getenv("GALLERY_CACHE_PATH")
	if cachePath == "" {
		panic("Missing GALLERY_CACHE_PATH")
	}

	thumbnails := cachePath + "\\thumbnails"
	previews := cachePath + "\\previews"
	images := cachePath + "\\images"
	return thumbnails, previews, images
}

var ThumbnailsCache, PreviewsCache, ImagesCache = caches()
