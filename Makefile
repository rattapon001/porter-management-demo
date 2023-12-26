
.PHONY: run-job
run-job:
	go run cmd/job/main.go

.PHONY: run-porter
run-porter:
	go run cmd/porter/main.go