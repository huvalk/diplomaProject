package constants

import (
	"time"
)

const (
	CookieName     = "sessionID"
	UserIdKey      = "userID"
	CookieDuration = 365 * 24 * time.Hour
)

const (
	CSRFHeader = "X-CSRF-TOKEN"
	CSRFKey    = "eE%yh?aAH_hYk*5h$DXvTddAGt2eWCt^+TT_4*$ADxz^X$5ue74jmeJT@z^+c_*v"
)

const (
	EventStatusClosed = "Closed"
	EventStatusOpen   = "Open"
)

const (
	PROD = "prod"
	DEV   = "dev"
	LOCAL   = "local"
)
