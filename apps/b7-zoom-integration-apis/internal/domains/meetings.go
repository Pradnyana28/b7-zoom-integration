package domains

import (
	"b7-zoom-integration/apps/b7-zoom-integration-apis/internal/services"
)

// ListMeetings list the created meetings
func ListMeetings() []interface{} {
	return []interface{}{}
}

// CreateMeetingInputDTO is the input for creating a meeting
type CreateMeetingInputDTO struct {
	Agenda    string `json:"agenda" binding:"required"`
	Type      int    `json:"type" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	Duration  int    `json:"duration" binding:"required"`
}

// CreateMeeting creates a new meeting
func CreateMeeting(input *CreateMeetingInputDTO) (*services.CreateMeetingOutput, error) {
	zoomService := services.NewZoomService()
	if _, err := zoomService.GetToken(); err != nil {
		return nil, err
	}

	resp, err := zoomService.CreateMeeting(&services.CreateMeetingInput{
		Agenda:    &input.Agenda,
		Type:      &input.Type,
		StartTime: &input.StartTime,
		Duration:  &input.Duration,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
