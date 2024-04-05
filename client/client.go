package client

type BlockScoutAPIClient struct {
	// URL API in the format "https://instance_base_url/api"
	URL string
	// URL APIv2 in the format "https://instance_base_url/api/v2/"
	URLv2 string
}

func (client *BlockScoutAPIClient) SetBlockScoutApiUrl(url string) {
	client.URL = url
}

func (client *BlockScoutAPIClient) SetBlockScoutApiUrlV2(url string) {
	client.URLv2 = url
}
