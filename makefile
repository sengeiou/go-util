mock:
	mockgen -source interface/http.go -destination mock/http_mock.go -package mock
	mockgen -source interface/redis.go -destination mock/redis_mock.go -package mock
	mockgen -source interface/repo.go -destination mock/repo_mock.go -package mock