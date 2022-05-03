package constants

// config names
const (
	RedisConfig                               = "redis"
	RedisBrokenJourneyConfigKey               = ""
	RedisBrokenJourneySentinalClientConfigKey = ""
	BrokenJourneyMySql                        = ""
	RedisAdditionalClientDetailsConfigKey     = ""
	CloudMasterDBMySql                        = ""
	RedisMTFEligibilityModelConfigKey         = ""
	RedisScripmasterClientConfigKey           = "ScripmasterClient"
	RedisGreeksClientConfigKey                = "GreeksClient"
	RedisGreeksSentinalClientConfigKey        = "GreeksSentinalClient"
)

const HTMLMSG1 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p><span>41%</span> of new users have done their first<br>investment today. Grab the <span>exciting offer</span><br> and place the SmartSauda!!</p></body></html> "
const HTMLMSG2 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p>Grab <span>offer</span> on your first funding.</p></body></html>  "
const HTMLMSG3 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p>Now you are <span>Investment ready!</span><br> Want to explore investment opportunities?</p></body></html>  "
const HTMLMSG4 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p>Now you are <span>Investment ready!</span></p></body></html>  "
const HTMLMSG5 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p><span>41%</span> of the new users have done<br> their first trade today.</p></body></html> "
const HTMLMSG6 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p>Be <span>Investment ready</span> by adding funds now.</p></body></html> "
const HTMLMSG7 = "<!DOCTYPE html><html><head><link href='https://fonts.googleapis.com/css?family=Roboto' rel='stylesheet'><style>body {font-family: 'Roboto';}p {line-height: 1.6;text-align: center;font-size: 14px;color: #00000099;}span {font-size: 16px;font-weight: bold;color: #01579B;}</style></head><body><p><span>41%</span> of the new users had become<br> <span>Investment ready</span> by adding funds today.</p></body></html>  "

const MSG1 = "41% of the new users have done their first trade today. Grab an exciting offer and place the SmartSauda!!"
const MSG2 = "Grab offer on your first funding."
const MSG3 = "Now you are investment ready! Want to explore investment opportunities?"
const MSG4 = "Now you are investment ready!"
const MSG5 = "41% of the new users have done their first trade today."
const MSG6 = "Be Investment ready by Adding fund now"
const MSG7 = "41% of the new users had become investment ready by adding funds today"

const (
	MTFUserTypeOne   = "1"
	MTFUserTypeTwo   = "2"
	MTFUserTypeThree = "3"
	//Key to store whether data is stored for a particular day
	// redisKeyPrefixForMTFDataSet = "IsMTFDataSetFor_"
	RedisKeyPrefixMTF = "MTFELigibility_v1"
)
