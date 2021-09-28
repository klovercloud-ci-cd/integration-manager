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
	APPEND_APPLICATION      = COMPANY_UPDATE_OPTION("APPEND_APPLICATION")
	APPEND_REPOSITORY       = COMPANY_UPDATE_OPTION("APPEND_REPOSITORY")
	SOFT_DELETE_APPLICATION = COMPANY_UPDATE_OPTION("SOFT_DELETE_APPLICATION")
	DELETE_APPLICATION      = COMPANY_UPDATE_OPTION("DELETE_APPLICATION")
	SOFT_DELETE_REPOSITORY  = COMPANY_UPDATE_OPTION("SOFT_DELETE_REPOSITORY")
	DELETE_REPOSITORY       = COMPANY_UPDATE_OPTION("DELETE_REPOSITORY")
)

type COMPANY_STATUS string

const (
	ACTIVE   = COMPANY_STATUS("ACTIVE")
	INACTIVE = COMPANY_STATUS("INACTIVE")
)

//
//type PIPELINE_RESOURCE_TYPE string
//
//const (
//	GIT         = PIPELINE_RESOURCE_TYPE("git")
//	IMAGE       = PIPELINE_RESOURCE_TYPE("image")
//	DEPLOYMENT  = PIPELINE_RESOURCE_TYPE("deployment")
//	STATEFULSET = PIPELINE_RESOURCE_TYPE("statefulset")
//	DAEMONSET   = PIPELINE_RESOURCE_TYPE("daemonset")
//	POD         = PIPELINE_RESOURCE_TYPE("pod")
//	REPLICASET  = PIPELINE_RESOURCE_TYPE("replicaset")
//)

type STEP_TYPE string

const (
	BUILD  = STEP_TYPE("BUILD")
	DEPLOY = STEP_TYPE("DEPLOY")
)

type PIPELINE_PURGING string

const (
	PIPELINE_PURGING_ENABLE  = PIPELINE_PURGING("ENABLE")
	PIPELINE_PURGING_DISABLE = PIPELINE_PURGING("DISABLE")
)

type GITHUB_URL string

const (
	GITHUB_RAW_CONTENT_BASE_URL = "https://raw.githubusercontent.com/"
	GITHUB_BASE_URL             = "https://github.com/"
	GITHUB_API_BASE_URL         = "https://api.github.com/"
)

type GIT_FILE_NAME string

const (
	PIPELINE_FILE_NAME = "pipeline"
)

type TRIGGER string

const (
	AUTO   = TRIGGER("AUTO")
	MANUAL = TRIGGER("MANUAL")
)

type PARAMS string

const (
	_TYPE = PARAMS("type")
	ENV   = PARAMS("env")
)

const PIPELINE_FILE_BASE_DIRECTORY = "klovercloud/pipeline"
const PIPELINE_DESCRIPTORS_BASE_DIRECTORY = "klovercloud/pipeline/descriptors"
