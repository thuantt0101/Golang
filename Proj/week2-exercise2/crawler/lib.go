package crawler

import "time"

type Data struct {
	Title         string
	PublishedDate time.Time
	Author        string
	Content       string
}
