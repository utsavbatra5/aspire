package constants

const (
	ServerPortKey                          = "server.port"
	DBConfigKey                            = "db"
	ResearchAdvisoryTokenConfigKey         = "research.advisory.token"
	InstaTradeTokenConfigKey               = "marketplace.instatrade.token"
	ResearchAdvisoryCronTimeConfigKey      = "cron.advisory.cronTime"
	ScripMasterLastUpdatedCronTime         = "cron.scripMaster.lastUpdatedCronTime"
	AmoTimeCronTimeKey                     = "cron.amoTime"
	LocalIndicesCronTimeKey                = "cron.localIndicesCronTime"
	UpdateNiftySensexScripCacheCronTimeKey = "cron.updateNiftySensexScripCacheCronTime"
	ServerUrlsConfigKey                    = "serverUrls"

	CacheAdvisoryGetCallCountConfigKey                   = "cache.advisory.getCallCount"
	CacheAdvisoryGetActivePicksConfigKey                 = "cache.advisory.getActivePicks"
	CacheAdvisoryGetBulletPointsConfigKey                = "cache.advisory.getBulletPoints"
	CacheMarketplaceGetFutureBuiltUpHeatMapConfigKey     = "cache.marketplace.getFutureBuiltUpHeatMap"
	CacheMarketplaceGetSectorDataConfigKey               = "cache.marketplace.getSectorData"
	CacheMarketplaceGetFuturesMarketsIndicatorsConfigKey = "cache.marketplace.getFuturesMarketsIndicators"
	CacheMarketplaceGetSectorPerformanceConfigKey        = "cache.marketplace.getSectorPerformance"
	CacheMarketplaceGetMarketTrendConfigKey              = "cache.marketplace.getMarketTrend"
	CacheMarketplaceGetIndicesTickerConfigKey            = "cache.marketplace.getIndicesTicker"
	CacheMarketplaceGetBannedScriptConfigKey             = "cache.marketplace.getBannedScript"
	CacheMarketplaceGetMarketMoversByMostConfigKey       = "cache.marketplace.getMarketMoversByMost"
	CacheMarketplaceGetSimilarStocksConfigKey            = "cache.marketplace.getSimilarStocks"
	CacheMarketplaceGetHighestBuiltupStrikeConfigKey     = "cache.marketplace.getHighestBuiltupStrike"
	CacheStockDetailsGetScripTrendConfigKey              = "cache.stockdetails.getScripTrend"
	CacheStockDetailsGetTodayPriceSummaryConfigKey       = "cache.stockdetails.getTodayPriceSummary"
	CacheStockDetailsGetSNLTermDriversConfigKey          = "cache.stockdetails.getSNLTermDrivers"
	CacheStockDetailsGetPriceSummaryAndUpdatesConfigKey  = "cache.stockdetails.getPriceSummaryAndUpdates"
	CacheStockDetailsGetFundamentalRatiosConfigKey       = "cache.stockdetails.getFundamentalRatios"
	CacheStockDetailsGetTechnicalAnalysisConfigKey       = "cache.stockdetails.getTechnicalAnalysis"
	CacheStockDetailsGetExchangeNewsConfigKey            = "cache.stockdetails.getExchangeNews"
	CacheStockDetailsGetStockShareholderConfigKey        = "cache.stockdetails.getStockShareholder"
	CacheStockDetailsGetValuationScorecardConfigKey      = "cache.stockdetails.getValuationScorecard"
	CacheStockDetailsGetScripMarketTrendConfigKey        = "cache.stockdetails.getScripMarketTrend"
	CacheStockDetailsGetStockFundamentalDetailsConfigKey = "cache.stockdetails.getStockFundamentalDetails"
	CacheStockDetailsGetQualityScorecardConfigKey        = "cache.stockdetails.getQualityScorecard"
	CacheScripMasterGetGreeksOptionsConfigKey            = "cache.scripmaster.getGreeksOptions"
	CacheScripMasterGetDerivativeDetailsConfigKey        = "cache.scripmaster.getDerivativeDetails"
	CacheHomepagePortfolioAdvisoryConfigKey              = "cache.homepage.portfolioAdvisory"
	CacheHomepageMarketUpdatesConfigKey                  = "cache.homepage.getMarketUpdates"
	CacheHomepageTemplateAppVersionConfigKey             = "cache.homepage.templateAppVersion"
	CacheHomepageTermsConfigKey                          = "cache.homepage.terms"
	CacheHomepageYoutubeConfigKey                        = "cache.homepage.youtube"

	DefaultCacheStalenessFactor   = 5
	CacheStalenessFactorKey       = "cache.staleFactor"
	MaxExpiryDelayInMillis        = "cache.maxExpiryDelayInMs"
	DefaultMaxExpiryDelayInMillis = 60000
	DefaultCacheDuration          = "2m"

	AMOOpenTimeConfigKey  = "time.amo.openTime"
	AMOCloseTimeConfigKey = "time.amo.closeTime"

	HomepageRegisteredInvestedUserTemplateID = "homepage.template.registeredInvestedUserTemplateID"
	HomepageRegisteredNotInvestedTemplateID  = "homepage.template.registeredNotInvestedTemplateID"
	HomepageNonRegisteredUserTemplateID      = "homepage.template.nonRegisteredUserTemplateID"
	HomepageYoutubeChannelID                 = "homepage.youtube.channelID"
	HomepageYoutubeMaxResults                = "homepage.youtube.maxResults"
	HomepageYoutubeKey                       = "homepage.youtube.key"
	HomepageYoutubeIconTypeKey               = "homepage.youtube.iconType"

	DBNameAwsChartConfigKey = "dbname.awschart"

	AWSRegionConfigKey        = "aws.region"
	AWSIDConfigKey            = "aws.id"
	AWSSecretConfigKey        = "aws.secret"
	AWSBucketConfigKey        = "aws.bucket"
	AWSHomepageTermsConfigKey = "aws.homepageTerms"

	EnvironmentUatKey        = "uat"
	EnvironmentCugKey        = "cug"
	EnvironmentProductionKey = "prod"

	LoggerConfig               = "logger"
	IsHttpDebugEnabledKey      = "logger.enablehttpdebug"
	IsAPIDebugEnabledKey       = "logger.enableapidebug"
	IsRequestLoggingEnabledKey = "logger.enablerequestlog"
	JWTSigningConfigKey        = "jwt.signingKey"
	JWTABMASigningConfigKey    = "jwt.abmasigningKey"
	JWTEnabledConfigKey        = "jwt.enabled"

	HomepageUrl                                = "/kaizen/v1/homepage"
	UserID                                     = "clientCode"
	ClientIDKey                                = "clientID"
	CronTimeDefault                            = "30 8 * * *"
	ScripMasterLastUpdatedCronTimeDefault      = "*/5 * * * *"
	AmoTimeDefault                             = "*/30 * * * *"
	LocalIndicesCronTimeDefault                = "*/10 * * * *"
	UpdateNiftySensexScripCacheCronTimeDefault = "* */8 * * *"
)
