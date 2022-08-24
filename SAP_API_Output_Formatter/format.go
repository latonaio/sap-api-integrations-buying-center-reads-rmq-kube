package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-buying-center-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBuyingCenterCollection(raw []byte, l *logger.Logger) ([]BuyingCenterCollection, error) {
	pm := &responses.BuyingCenterCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BuyingCenterCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	buyingCenterCollection := make([]BuyingCenterCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		buyingCenterCollection = append(buyingCenterCollection, BuyingCenterCollection{
			ObjectID:            data.ObjectID,
			BuyingCenterID:      data.BuyingCenterID,
			ObjectUUID:          data.ObjectUUID,
			ObjectTypeCode:      data.ObjectTypeCode,
			CustomerID:          data.CustomerID,
			OpportunityID:       data.OpportunityID,
			Name:                data.Name,
			EntityLastChangedOn: data.EntityLastChangedOn,
			ETag:                data.ETag,
		})
	}

	return buyingCenterCollection, nil
}