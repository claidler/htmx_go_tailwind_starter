dev:
	go run .

build:
	go build -o ./tmp/ .

watch-css:
	npx tailwindcss -i ./styles.css -o ./public/styles.css --watch
