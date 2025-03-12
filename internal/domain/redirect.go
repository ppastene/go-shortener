package domain

import "time"

type Rediect struct {
	Url        string
	Expiration time.Time
}
