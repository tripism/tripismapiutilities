package tripismapiutilities

// ImageFormatFromMimeType returns the image format for a give image mime type
func ImageFormatFromMimeType(mimetype string) string {
	switch mimetype {
	case "image/gif":
		return "gif"
	case "image/jpg":
		return "jpeg"
	case "image/jpeg":
		return "jpeg"
	case "image/png":
		return "png"
	case "image/svg+xml":
		return "svg"
	case "image/x-icon":
		return "ico"
	}
	return "jpeg"
}
