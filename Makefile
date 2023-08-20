clean:
	go mod tidy
run-server:
	go run server/main.go
run-client:
	cd sveltekit-client && pnpm dev