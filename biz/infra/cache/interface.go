package cache

type IProxyCache interface {
	GetTargetBaseUrl(key string) (string, bool)
	SetTargetBaseUrl(key string, targetBaseUrl string)
	InvalidateByToken(token string) error
}
