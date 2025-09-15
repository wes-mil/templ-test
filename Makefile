dev:
	npx @tailwindcss/cli -i ./components/css/input.css -o ./static/css/output.css -m
	templ generate --watch --cmd "go run ."