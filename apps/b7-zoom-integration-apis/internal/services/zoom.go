package services

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// ZoomService is the service for Zoom
type ZoomService struct {
	Client          *resty.Client
	GrantType       string `json:"grant_type"`
	AccountID       string `json:"account_id"`
	Authorization   string `json:"authorization"`
	RequestTokenURL string `json:"request_token_url"`
	AccessToken     string `json:"access_token"`
	ApiURL          string `json:"api_url"`
}

// NewZoomService creates a new ZoomService
func NewZoomService() *ZoomService {
	return &ZoomService{
		Client:          resty.New(),
		GrantType:       "account_credentials",
		AccountID:       os.Getenv("ZOOM_ACCOUNT_ID"),
		Authorization:   os.Getenv("ZOOM_AUTHORIZATION"),
		RequestTokenURL: os.Getenv("ZOOM_REQUEST_TOKEN_URL"),
	}
}

// GetTokenOutput is the result of the GetToken
type GetTokenOutput struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	ApiURL      string `json:"api_url"`
}

// GetToken gets the Zoom token
func (z *ZoomService) GetToken() (*ZoomService, error) {
	var responseBody *GetTokenOutput

	resp, err := z.Client.R().
		SetHeaders(map[string]string{
			"Authorization": "Basic " + z.Authorization,
		}).
		SetQueryParams(map[string]string{
			"grant_type": z.GrantType,
			"account_id": z.AccountID,
		}).
		SetResult(&responseBody).
		Post(z.RequestTokenURL)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, err
	}

	z.AccessToken = responseBody.AccessToken
	z.ApiURL = responseBody.ApiURL

	return z, nil
}

// MeOutput is the result of the Me
type MeOutput struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	RoleName  string `json:"role_name"`
}

// Me is the current user
func (z *ZoomService) Me() (*MeOutput, error) {
	var responseBody *MeOutput

	resp, err := z.Client.R().
		SetHeaders(map[string]string{
			"Authorization": "Bearer " + z.AccessToken,
		}).
		SetResult(&responseBody).
		Get(z.ApiURL + "/v2/users/me")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to fetched current user data")
	}

	return responseBody, nil
}

// CreateMeetingInput is the input for the CreateMeeting
type CreateMeetingInput struct {
	Agenda    *string `json:"agenda"`
	Type      *int    `json:"type"`
	StartTime *string `json:"start_time"`
	Duration  *int    `json:"duration"`
}

// CreateMeetingOutput is the output for the CreateMeeting
type CreateMeetingOutput struct {
	ID string `json:"id"`
}

// CreateMeeting creates a new Zoom meeting
func (z *ZoomService) CreateMeeting(input *CreateMeetingInput) (*CreateMeetingOutput, error) {
	var output *CreateMeetingOutput

	resp, err := z.Client.R().
		SetHeaders(map[string]string{
			"Authorization": "Bearer " + z.AccessToken,
		}).
		SetBody(input).
		SetResult(&output).
		Post(z.ApiURL + "/v2/users/me/meetings")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to create meeting")
	}

	return output, nil
}
