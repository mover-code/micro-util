package umfPay

type Config struct {
	MER_ID      string `json:"mer_id"`
	URL_STR     string `json:"url_str"`
	NOTIFY_URL  string `json:"notify_url"`
	SERVER_IP   string `json:"server_ip"`
	PUBLIC_KEY  string `json:"public_key"`
	PRIVATE_KEY string `json:"private_key"`
}
