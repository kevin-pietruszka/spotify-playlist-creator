package auth

const (
	AUTH_BASE = "https://accounts.spotify.com/"

	RESPONSE_TYPE = "code"
	REDIRECT_URI  = "http://localhost:8888/callback"
	SCOPE         = "playlist-modify-private"
)

type SpotifyAuthClient struct {
	client_id     string
	response_type string
	redirect_uri  string
	scope         string
}

func NewSpotifyAuthClient(client_id string) *SpotifyAuthClient {
	return &SpotifyAuthClient{
		client_id:     client_id,
		response_type: RESPONSE_TYPE,
		redirect_uri:  REDIRECT_URI,
		scope:         SCOPE,
	}
}
