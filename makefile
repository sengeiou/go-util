mock:
	mockgen -source interface/http.go -destination mock/http_mock.go -package mock
	mockgen -source interface/cache.go -destination mock/cache_mock.go -package mock