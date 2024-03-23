package main

import (
	sess "library/Security"
	"time"
)

func main() {

	asd := sess.SessionConfig{
		Expiration:   time.Second * 1000,
		CookieSecure: true,
	}
	sess.SessionCreate(asd)
	sess.SessionMiddleware(asd)
}
