package domain

type AuthURL struct {
	ServiceId string
	Url       string
	ForceAuth int
	Prefix    int
}

func (au *AuthURL) IsForeLogin() bool {
	return au.ForceAuth == 1
}

func (au *AuthURL) IsMustMatchPrefix() bool {
	return au.Prefix == 1
}
