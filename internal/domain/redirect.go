package domain

import "time"

type Redirect struct {
	Url        string
	Expiration time.Time
}
