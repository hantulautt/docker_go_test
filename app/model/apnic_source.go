package model

type ApnicSource struct {
	InetNum string
	NetName string
	Desc    string
}

type ApnicSourceResponse struct {
	IpRange     string `json:"ip_range"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
