package constants

import "time"

var (
	MySQLDriverName      = "mysql"
	PostgresDriverName   = "postgres"
	SQLServerDrivernName = "mssql"

	ApplicationName = "bbe-discovery"
	CounterKey      = "key"
	RequestIDKey    = "requestID"
	RequestQueryKey = "RequestQuery"
	RequestBodyKey  = "RequestBody"
	ARQPrefix       = "ARQ"
	ARQDelimiter    = "========"
	TimeFormat      = "2006-01-02T15:04:05.000Z"
	DeviceType      = "x-devicetype"
	OperatingSystem = "x-operatingsystem"

	NonTradingAccessTokenHeader               = "AccessToken"
	TestMode                                  = "test"
	ReleaseMode                               = "release"
	EnvProdValue                              = "prod"
	HostnameKey                               = "HOSTNAME"
	UserKey                                   = "USER"
	AWSRegionKey                              = "AWS_REGION"
	AWSRegionDefaultValue                     = "ap-south-1"
	AWSAccessKeyID                            = "AWS_ACCESS_KEY_ID"
	AWSSecretAccessKey                        = "AWS_SECRET_ACCESS_KEY"
	AWSBucket                                 = "AWS_BUCKET"
	BaseConfigPathKey                         = "base-config-path"
	BaseConfigPathDefaultValue                = "conf"
	BaseConfigPathUsage                       = "path to folder that stores your configurations"
	MaxDuration                 time.Duration = 1<<63 - 1
	CacheTimestampLayout                      = "20060102150405"
	CacheDataSeperator                        = "|^|"
	ScripMasterTimeFormat                     = "2006-01-02 15:04:05"
)
