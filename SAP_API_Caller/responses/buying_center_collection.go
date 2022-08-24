package responses

type BuyingCenterCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID            string `json:"ObjectID"`
			BuyingCenterID      string `json:"BuyingCenterID"`
			ObjectUUID          string `json:"ObjectUUID"`
			ObjectTypeCode      string `json:"ObjectTypeCode"`
			CustomerID          string `json:"CustomerID"`
			OpportunityID       string `json:"OpportunityID"`
			Name                string `json:"Name"`
			EntityLastChangedOn string `json:"EntityLastChangedOn"`
			ETag                string `json:"ETag"`
		} `json:"results"`
	} `json:"d"`
}
