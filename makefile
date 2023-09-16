test:
	go test $(shell go list ./... | grep -v /database | grep -v /interfaces | grep -v /entities)