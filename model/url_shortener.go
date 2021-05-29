package model

import "time"

// UrlShortener represent the url_shortener model
type UrlShortener struct {
	Slug		string		`json:"shortCode"`
	URL			string		`json:"url"`
	CreatedAt	time.Time	`json:"startDate"`
	LastSeen	time.Time	`json:"lastSeenDate"`
	Hits		int			`json:"redirectCount"`
}

type ShortenReq struct {
	Slug	string	`json:"shortCode"`
	URL		string	`json:"url"`
}

type ShortenResp struct {
	Slug	string	`json:"shortCode"`
}

type GetShortCodeResp struct {
	Location	string	`json:"location"`
}