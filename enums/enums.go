package enums

const (
	Mongo    = "MONGO"
	Inmemory = "INMEMORY"
)

type REPOSITORY_TYPE string

const (
	GITHUB     = REPOSITORY_TYPE("GITHUB")
	BIT_BUCKET = REPOSITORY_TYPE("BIT_BUCKET")
)

type COMPANY_UPDATE_OPTION string

const (
	APPEND_APPLICATION = COMPANY_UPDATE_OPTION("APPEND_APPLICATION")
	APPEND_REPOSITORY  = COMPANY_UPDATE_OPTION("APPEND_REPOSITORY")
)

type COMPANY_STATUS string

const (
	ACTIVE   = COMPANY_STATUS("ACTIVE")
	INACTIVE = COMPANY_STATUS("INACTIVE")
)
