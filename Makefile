dev:
	go run .

watch-css:
	npx tailwindcss -i ./styles.css -o ./public/styles.css --watch

build-css:
	npx postcss styles.css -o public/styles.css