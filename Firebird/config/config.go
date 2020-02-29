package config

// API KEY
const (
	// TODO: replace with your own AccessKey and Secret Key
	ACCESS_KEY string = "*" // huobi申请的apiKey  必填
	SECRET_KEY string = "*" // huobi申请的secretKey  必填

	//HOST_NAME string = "api.huobi.so"

	Local_IP string = "192.168.1.104" //本地IP地址 Your Local IP  选填

	ENABLE_PRIVATE_SIGNATURE bool = false

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// replace with your own PrivateKey from privatekey.pem
	PRIVATE_KEY_PRIME_256 string = ``

	// default symbol sync count
	SYMBOL_SYNC_COUNT = 3 // include yesterday and before yesterday.
)

var hostName string = "api.huobiasia.vip"

func GetHostName() string  {
	return hostName
}

func SetHostName(host string)  {
	if "" != host {
		hostName = host
	}
}

// API请求地址, 不要带最后的/
func GetMarketUrl() string {
	return "https://" + hostName
}

func GetTradeUrl() string {
	return "https://" + hostName
}

func GetWsUrl() string {
	return "wss://" + hostName + "/ws"
}

func GetWsOrderUrl() string {
	return "wss://" + hostName + "/ws/v1"
}