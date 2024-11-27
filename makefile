## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i ./css/input.css -o ./css/output.css --watch