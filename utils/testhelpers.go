package utils

import "net/http"

func SetRestCountriesBaseURL(url string) func() {
	original := restCountriesBase
	restCountriesBase = url
	return func() { restCountriesBase = original }
}


func SetOpenTripMapBaseURL(url string) func() {
	original := openTripMapBase
	openTripMapBase = url
	return func() { openTripMapBase = original }
}


func SetHTTPClient(c *http.Client) func() {
	original := httpClient
	httpClient = c
	return func() { httpClient = original }
}

func SetWishlistFile(path string) func() {
	original := wishlistFile
	wishlistFile = path
	return func() { wishlistFile = original }
}

func SetUserFile(path string) func() {
	original := userFile
	userFile = path
	return func() { userFile = original }
}