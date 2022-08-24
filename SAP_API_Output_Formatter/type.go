package sap_api_output_formatter

type BuyingCenter struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	APISchema        string `json:"api_schema"`
	BuyingCenterCode string `json:"buying_center_code"`
	Deleted          bool   `json:"deleted"`
}

type BuyingCenterCollection struct {
	ObjectID            string `json:"ObjectID"`
	BuyingCenterID      string `json:"BuyingCenterID"`
	ObjectUUID          string `json:"ObjectUUID"`
	ObjectTypeCode      string `json:"ObjectTypeCode"`
	CustomerID          string `json:"CustomerID"`
	OpportunityID       string `json:"OpportunityID"`
	Name                string `json:"Name"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
	ETag                string `json:"ETag"`
}
