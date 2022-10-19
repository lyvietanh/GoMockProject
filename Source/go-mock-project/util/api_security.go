package util

type SecuredHeader struct {
	AccessToken string `header:"access-token"`
}
