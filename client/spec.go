package client

type Account struct {
	ID            string   `json:"id"`
	AccountName   string   `json:"account_name,omitempty"`
	DefaultRegion string   `json:"default_region,omitempty"`
	Regions       []string `json:"regions,omitempty"`
}

type Spec struct {
	Regions           []string           `json:"regions,omitempty"`
	Accounts          []Account          `json:"accounts"`
	AWSDebug          bool               `json:"aws_debug,omitempty"`
	MaxRetries        *int               `json:"max_retries,omitempty"`
	MaxBackoff        *int               `json:"max_backoff,omitempty"`
	ServiceRateLimits map[string]float64 `json:"service_rate_limits,omitempty"`
}
