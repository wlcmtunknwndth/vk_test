package models

import docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"

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

func ProtocDocToDoc(req *docsv1.TDocument) *TDocument {
	return &TDocument{
		Url:            req.Url,
		PubDate:        req.PubDate,
		FetchTime:      req.FetchTime,
		Text:           req.Text,
		FirstFetchTime: req.FirstFetchTime,
	}
}

func DocToProtocDoc(req *TDocument) *docsv1.TDocument {
	return &docsv1.TDocument{
		Url:            req.Url,
		PubDate:        req.PubDate,
		FetchTime:      req.FetchTime,
		Text:           req.Text,
		FirstFetchTime: req.FirstFetchTime,
	}
}

func UsrDocToProtocUsrDoc(req *UserTDocument) *docsv1.UserTDocument {
	return &docsv1.UserTDocument{
		Url:  req.Url,
		Text: req.Text,
	}
}

func ProtocUsrDocToUsrDoc(req *docsv1.UserTDocument) *UserTDocument {
	return &UserTDocument{
		Url:  req.Url,
		Text: req.Text,
	}
}
