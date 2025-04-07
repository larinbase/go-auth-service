package dto

type TokenCoupleResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenCoupleResponse(accessToken string, refreshToken string) *TokenCoupleResponse {
	return &TokenCoupleResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
