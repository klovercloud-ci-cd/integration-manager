package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION mongo as db
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

// REPOSITORY_TYPE repository types[may be any git]
type REPOSITORY_TYPE string

const (
	// GITHUB github as repository
	GITHUB = REPOSITORY_TYPE("GITHUB")
	// BIT_BUCKET bitbucket as repository
	BIT_BUCKET = REPOSITORY_TYPE("BIT_BUCKET")
)

// REPOSITORY_UPDATE_OPTION company update options
type REPOSITORY_UPDATE_OPTION string

const (
	// APPEND_REPOSITORY company update option to append repository
	APPEND_REPOSITORY = REPOSITORY_UPDATE_OPTION("APPEND_REPOSITORY")
	// SOFT_DELETE_REPOSITORY company update option to soft delete repository
	SOFT_DELETE_REPOSITORY = REPOSITORY_UPDATE_OPTION("SOFT_DELETE_REPOSITORY")
	// DELETE_REPOSITORY company update option to delete repository
	DELETE_REPOSITORY = REPOSITORY_UPDATE_OPTION("DELETE_REPOSITORY")
)

// APPLICATION_UPDATE_OPTION company update options
type APPLICATION_UPDATE_OPTION string

const (
	// APPEND_APPLICATION company update option to append application
	APPEND_APPLICATION = APPLICATION_UPDATE_OPTION("APPEND_APPLICATION")
	// SOFT_DELETE_APPLICATION company update option to soft delete application
	SOFT_DELETE_APPLICATION = APPLICATION_UPDATE_OPTION("SOFT_DELETE_APPLICATION")
	// DELETE_APPLICATION company update option to delete application
	DELETE_APPLICATION = APPLICATION_UPDATE_OPTION("DELETE_APPLICATION")
)

// PIPELINE_GET_OPTION pipeline get options
type PIPELINE_GET_OPTION string

const (
	// GET_PIPELINE_FOR_VALIDATION pipeline get option for validation
	GET_PIPELINE_FOR_VALIDATION = PIPELINE_GET_OPTION("GET_PIPELINE_FOR_VALIDATION")
)

// COMPANY_STATUS company status options
type COMPANY_STATUS string

const (
	// ACTIVE company status for active company
	ACTIVE = COMPANY_STATUS("ACTIVE")
	// INACTIVE company status for inactive company
	INACTIVE = COMPANY_STATUS("INACTIVE")
)

// STEP_TYPE steps type
type STEP_TYPE string

const (
	// BUILD step that builds image from source code
	BUILD = STEP_TYPE("BUILD")
	// DEPLOY step that deploys workloads and others to cluster
	DEPLOY = STEP_TYPE("DEPLOY")
	// INTERMEDIARY step that runs custom jobs
	INTERMEDIARY = STEP_TYPE("INTERMEDIARY")
	// JENKINS_JOB step that runs jenkins jobs
	JENKINS_JOB = STEP_TYPE("JENKINS_JOB")
)

// GITHUB_URL gitbhub url for different operations
type GITHUB_URL string

const (
	// GITHUB_RAW_CONTENT_BASE_URL gitbhub url for raw content
	GITHUB_RAW_CONTENT_BASE_URL = "https://raw.githubusercontent.com/"
	// GITHUB_BASE_URL gitbhub base url
	GITHUB_BASE_URL = "https://github.com/"
	// GITHUB_API_BASE_URL gitbhub base url for api access
	GITHUB_API_BASE_URL = "https://api.github.com/"
)

// BITBUCKET_URL bitbucket url for different operations
type BITBUCKET_URL string

const (
	// BITBUCKET_RAW_CONTENT_BASE_URL bitbucket url for raw content
	BITBUCKET_RAW_CONTENT_BASE_URL = "https://bitbucket.org/api/2.0/repositories/"
	// BITBUCKET_BASE_URL bitbucket base url
	BITBUCKET_BASE_URL = "https://bitbucket.org/"
	// BITBUCKET_API_BASE_URL bitbucket base url for api access
	BITBUCKET_API_BASE_URL = "https://api.bitbucket.org/2.0/"
)

const (
	// PIPELINE_FILE_NAME pipeline containing file name
	PIPELINE_FILE_NAME = "pipeline"
)

// TRIGGER pipeline trigger options
type TRIGGER string

const (
	// AUTO pipeline trigger options is auto
	AUTO = TRIGGER("AUTO")
	// MANUAL pipeline trigger options is MANUAL
	MANUAL = TRIGGER("MANUAL")
)

// PARAMS pipeline parameters
type PARAMS string

const (
	// REVISION resource revision key for  pipeline step param
	REVISION = PARAMS("revision")

	// ALLOWED_BRANCHES allowed branches for this pipeline
	ALLOWED_BRANCHES= PARAMS("allowed_branches")

	// IMAGE resource image key for  pipeline step param
	IMAGE = PARAMS("images")

	// STORAGE resource storage key for  pipeline step param
	STORAGE = PARAMS("storage")

	// ACCESS resource access mode key for  pipeline step param
	ACCESS = PARAMS("access_mode")

	// BUILD_TYPE resource build type key for  pipeline step param
	BUILD_TYPE = PARAMS("build_type")

	// REPOSITORY_TYPE_PARAM resource repository type
	REPOSITORY_TYPE_PARAM = PARAMS("repository_type")
)

// VARIABLES pipeline parameters variables
type VARIABLES string

const  (
    BRANCH= VARIABLES("$BRANCH")
)


type ACCESS_MODE string

const (
	// READ_WRITE_ONCE access mode for read and write once
	READ_WRITE_ONCE = ACCESS_MODE("ReadWriteOnce")

	// READ_WRITE_MANY access mode for read write many
	READ_WRITE_MANY = ACCESS_MODE("ReadWriteMany")

	// READ_ONLY_MANY access mode for read only many
	READ_ONLY_MANY = ACCESS_MODE("ReadOnlyMany")

	// READ_WRITE_ONCE_POD access mode for read write once pod
	READ_WRITE_ONCE_POD = ACCESS_MODE("ReadWriteOncePod")
)

// PIPELINE_FILE_BASE_DIRECTORY pipeline file base directory
const PIPELINE_FILE_BASE_DIRECTORY = "klovercloud/pipeline"

// PIPELINE_DESCRIPTORS_BASE_DIRECTORY pipeline descriptors base directory
const PIPELINE_DESCRIPTORS_BASE_DIRECTORY = "klovercloud/pipeline/configs"

// GITHUB_EVENT git web hook event options
type GITHUB_EVENT string

const (
	// GITHUB_PUSH_EVENT git web hook push event option
	GITHUB_PUSH_EVENT = GITHUB_EVENT("push")
	// GITHUB_RELEASE_EVENT git web hook release event option
	GITHUB_RELEASE_EVENT = GITHUB_EVENT("release")
	// GITHUB_DELETE_EVENT git web hook delete event option
	GITHUB_DELETE_EVENT = GITHUB_EVENT("delete")
	// WEBHOOK_ENABLE git web hook enable event option
	WEBHOOK_EANBLE = GITHUB_EVENT("enable")
	// WEBHOOK_DISABLE git web hook disable event option
	WEBHOOK_DISABLE = GITHUB_EVENT("disable")
)

// BITBUCKET_EVENT git web hook event options
type BITBUCKET_EVENT string

const (
	// BITBUCKET_PUSH_EVENT git web hook push event option
	BITBUCKET_PUSH_EVENT = GITHUB_EVENT("push")
)
