package domains

import (
	"b7-zoom-integration/apps/b7-zoom-integration-apis/internal/services"
)

// Me is the current user
func Me() (*services.MeOutput, error) {
	zoomService := services.NewZoomService()
	if _, err := zoomService.GetToken(); err != nil {
		return nil, err
	}

	users, err := zoomService.Me()
	if err != nil {
		return nil, err
	}

	return users, nil
}
