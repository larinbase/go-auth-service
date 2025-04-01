package dto

type JWTResponse struct {
	AccessToken string `json:"access_token"`
}

func NewJWTResponse(accessToken string) *JWTResponse {
	return &JWTResponse{
		AccessToken: accessToken,
	}
}
