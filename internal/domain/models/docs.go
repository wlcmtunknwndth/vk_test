package models

type UserTDocument struct {
	Url  string
	Text string
}

type TDocument struct {
	Url            string
	PubDate        uint64
	FetchTime      uint64
	Text           string
	FirstFetchTime uint64
}
