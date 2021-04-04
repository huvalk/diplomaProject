package globalVars

import "os"

var (
	ENV                  = os.Getenv("ENV")
	JWT_SECRET           = os.Getenv("JWT_SECRET")
	FRONTEND_URI         = os.Getenv("FRONTEND_URI")
	CLIENT_ID            = os.Getenv("CLIENT_ID")
	BACKEND_URI          = os.Getenv("BACKEND_URI")
	STATE                = os.Getenv("STATE")
	CLIENT_SECRET        = os.Getenv("CLIENT_SECRET")
	POSTGRES_USER        = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD    = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB          = os.Getenv("POSTGRES_DB")
	TEAMUP_BUCKET_ID     = os.Getenv("TEAMUP_BUCKET_ID")
	TEAMUP_BUCKET_SECRET = os.Getenv("TEAMUP_BUCKET_SECRET")
	TEAMUP_BUCKET_NAME   = os.Getenv("TEAMUP_BUCKET_NAME")
)
