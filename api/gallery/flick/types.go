package flick

type UploadParams struct {
	Path        string `json:"path"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
}
