package dataverse_service

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint16 `json:"expires_in"`
}
