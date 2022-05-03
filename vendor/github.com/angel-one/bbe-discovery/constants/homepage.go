package constants

const (
	RegisteredUser                   = "Registered User"
	NonRegisteredUser                = "Non Registered User"
	RegisteredNotInvestedTemplateID  = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14"
	NonRegisteredUserTemplateID      = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12"
	RegisteredInvestedUserTemplateID = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13"

	HomePagePostgres        = "HomePage"
	UserActionWidget        = "USER_ACTION"
	KYCWidget               = "USER_ACTION_2"
	MarketsWidget           = "MARKETS_WIDGET"
	AdvisoryWidget          = "ADVISORY_LAYOUT"
	PortfolioWidget         = "PORTFOLIO_LAYOUT"
	PortfolioAdvisoryWidget = "PORTFOLIO_ADVISORY"
	ExternalAppWidget       = "EXTERNAL_APP"

	CallCategoryIntraday  = "Intraday"
	CallCategoryShortTerm = "ShortTerm"
	CallCategoryLongTerm  = "LongTerm"

	NonRegisteredID         = "NonRegistered"
	RegisteredInvestedID    = "RegisteredInvested"
	RegisteredNonInvestedID = "RegisteredNonInvested"

	YoutubeSearchOrder = "date"
	YoutubeSearchPart  = "snippet"

	YoutubeVideoPart         = "contentDetails"
	YoutubeDurationSeperator = "PT"

	YoutubeAngelChannel         = "UCHzAyAwFbp2KVwvZ_fJuH0Q"
	HomepageYoutubeDataCacheKey = "youtube_data"
)

//AMXConstants TO accept
var AMXConstants = map[string]string{
	"USERID":    "#USERID",
	"SESSIONID": "#SESSIONID",
	"SEGMENT":   "#SEGMENT",
	"EXCHANGE":  "#EXCHANGE",
	"PRODUCT":   "#PRODUCT",
}

//AMX API Req Constants ...
var AMXAPI = map[string]string{
	"RMS_LIMIT_REQ": "203|" + AMXConstants["USERID"] + "|" + AMXConstants["SESSIONID"] + "|" + AMXConstants["USERID"] + "|" + AMXConstants["USERID"] + "|" + AMXConstants["SEGMENT"] + "|" + AMXConstants["EXCHANGE"] + "|" + AMXConstants["PRODUCT"] + "|",
}
