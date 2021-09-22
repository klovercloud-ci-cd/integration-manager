package service

type HttpClient interface {
	Get(url string,header map[string]string)(error, []byte)
	Post(url string, header map[string]string, body []byte) error
}
