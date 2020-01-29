package config

// API KEY
const (
	// TODO: replace with your own AccessKey and Secret Key
	ACCESS_KEY string = "*" // huobi申请的apiKey  必填
	SECRET_KEY string = "*" // huobi申请的secretKey  必填

	// API请求地址, 不要带最后的/
	MARKET_URL   string = "https://api.huobi.so"
	TRADE_URL    string = "https://api.huobi.so"
	WS_URL       string = "wss://api.huobi.so/ws"
	WS_ORDER_URL string = "wss://api.huobi.so/ws/v1"
	Local_IP     string = "192.168.1.104" //本地IP地址 Your Local IP  选填

	//replace with real URLs and HostName
	HOST_NAME string = "api.hbdm.com"

	ENABLE_PRIVATE_SIGNATURE bool = false

	// generated the key by: openssl ecparam -name prime256v1 -genkey -noout -out privatekey.pem
	// only required when Private Signature is enabled
	// replace with your own PrivateKey from privatekey.pem
	PRIVATE_KEY_PRIME_256 string = ``

	// default symbol sync count
	SYMBOL_SYNC_COUNT = 3 // include yesterday and before yesterday.
)
