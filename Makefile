build:
	@go build -o out/breakout main.go
	@cp -r ./res ./out/

run:
	@./out/breakout
