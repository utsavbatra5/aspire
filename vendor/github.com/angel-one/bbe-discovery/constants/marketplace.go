package constants

const (
	MarketplaceDataPath      = "/api/FNOAPI/GetData"
	MarketplaceGroupDataPath = "/api/FNOAPI/GetGroupData"
	IndicesTickerPath        = "/MobileService/AEMobile.svc/getlocalindices?responsetype=json"
	SimilarStocks            = "/stocks/getPeersDetails/"
	ProjectDateFormat        = "2006-01-02 15:04:05"
	SelectLocalIndicesQuery  = `SELECT DISTINCT V.co_code,V.exchange,V.sc_code,symbol=V.cm_symbol,V.co_name,
CurrValue=ISNULL(B.price,N.LTP),PrevValue=ISNULL(B.prev_close,N.prev_close),
LUT=ISNULL(B.upd_time,N.tr_date)
FROM ANGEL_WMS.dbo.vw_stock_mst V WITH (NOLOCK)
LEFT OUTER JOIN FSS_EQUITY.dbo.bi_bseindices_latestprices B WITH (NOLOCK) ON V.sc_code = B.sc_code
LEFT OUTER JOIN FSS_EQUITY.dbo.bi_nselatest_latestprices N WITH (NOLOCK) ON V.symbol = N.symbol
WHERE V.co_code IN (20558,21553,21554,21555,21556,21557,21558,21559,23998,25112,25113,25114,26567,26568,28155,35223,
20559,22113,24484,26753,26768,26769,26771,27176,28203,28253,30183,40939,41356,41383)
AND V.sc_code <> '21335'`
	SelectNiftySensexScripDataQuery = `exec GET_NIFTY_SENSEX_SCRIPS`
	FutureBuiltUpHeatMapAPIName     = "GetFutureBuiltUpHeatMap"
	SectorDataAPIName               = "GetSectorData"
	FuturesMarketsIndicatorsAPIName = "GetFuturesMarketsIndicators"
	SectorPerformanceAPIName        = "GetSectorPerformance"
	InstaTradeToken                 = "FGjASChztd6rca+LbunqsbL0B9bV0F9brGpvuuz+8kW7GTRivHyGxCwki6IFDtnBW8T0SKc4We+KfqTJ3eS3ztb/5ye5u4sTeV2CSslUGUl3XNe37rjx+efb8exg5rK1d2gq3Ew2OIH+m3AyyIvyPmBCD7K4rXi69mTV3O9cG9WxW+tgIzCkxTSCHKdWC3Dp+O+pGeI8cOmIBIUOWfUZeYQVI8N1LUF7bQmEARe/CaOfNSCyMJOvn1y4MF1LT86YuhU/MN0InvUFYIkXuu95EfY8mptsHYDEt7V4Qv7DWenpdzkV97+7oxKLOL39e2sxtvmg0UCYOtOWSSfR72WYbw=="

	LocalIndicesCacheKey      = "LOCAL_INDICES"
	NiftySensexScripCacheKey  = "NIFTY_SENSEX_SCRIPT"
	CommonTradeRedisNamespace = "common-trade"
	LTPRedisFetchConstant     = 100
	AMXuserIDConfigKey        = "marketplace.amx.userID"
	AMXpasswordConfigKey      = "marketplace.amx.pass"
	MostActiveVolume          = "MOST_ACT_VOLUMN"
	MostActiveValue           = "MOST_ACT_VALUE"
	BSEExchangeConstant       = "BSE"
	NSEExchangeConstant       = "NSE"
)

var MarketTrendPayload = []map[string]interface{}{
	{
		"name": "Symbol", "value": "NIFTY",
	},
	{
		"name": "TimeFrame", "value": "I900",
	},
	{
		"name": "ExchCode", "value": "NFO",
	},
	{
		"name": "Series", "value": "FUTIDX",
	},
	{
		"name": "CallPut", "value": "",
	},
	{
		"name": "StrikePrice", "value": "0",
	},
	{
		"name": "ExpiryDate", "value": "",
	},
}
