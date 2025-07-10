package tripismapiutilities

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/globalsign/mgo/bson"
)

// GetAvatarImageDataFromURL retrieves the avatar image data from the avatarURL specified
func GetAvatarImageDataFromURL(avatarURL string) (bson.Binary, string, error) {
	var noData bson.Binary
	// Parse the remote URL
	u, err := url.Parse(avatarURL)
	if err != nil {
		return noData, "", err
	}
	// Make a GET request to the URL
	res, err := HTTPClient.Get(u.String())
	if err != nil {
		return noData, "", err
	}
	defer res.Body.Close()
	// Did the remote server respond with a 200 OK?
	if res.StatusCode != http.StatusOK {
		// Was the response a redirection?
		if res.StatusCode == http.StatusMovedPermanently || res.StatusCode == http.StatusTemporaryRedirect || res.StatusCode == http.StatusFound {
			loc := res.Header.Get("location")
			if len(loc) > 0 {
				// Recursively call method to allow following redirect
				recdata, rectype, err := GetAvatarImageDataFromURL(loc)
				if err != nil {
					return noData, "", err
				}
				return recdata, rectype, nil
			}
			return noData, "", errors.New("unable to follow redirect from remote server")
		}
		// Return error to calling function and handle default image handler
		return noData, "", errors.New("received " + res.Status + "back from server")
	}

	// We read all the bytes of the image
	// Types: data []byte
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return noData, "", errors.New("unable to read content of response")
	}
	// Work out what the content type returned was from the Content-Type header
	contentType := res.Header.Get("Content-Type")
	imgformat := "gif"
	switch contentType {
	case "image/gif":
		imgformat = "gif"
	case "image/jpg":
		imgformat = "jpeg"
	case "image/jpeg":
		imgformat = "jpeg"
	case "image/png":
		imgformat = "png"
	default:
		// Couldn't determine mime type - try and extract from URL end of path
		if strings.HasSuffix(u.String(), ".gif") {
			imgformat = "gif"
		} else if strings.HasSuffix(u.String(), ".jpg") {
			imgformat = "jpeg"
		} else if strings.HasSuffix(u.String(), ".jpeg") {
			imgformat = "jpeg"
		} else if strings.HasSuffix(u.String(), ".png") {
			imgformat = "png"
		}
	}

	var imageData bson.Binary
	imageData.Kind = 0x00
	imageData.Data = data

	return imageData, imgformat, nil
}
