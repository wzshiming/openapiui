IMAGE := swaggerapi/swagger-editor
PWD   := $(shell pwd)

.PHONY: swaggereditor
swaggereditor:
	docker pull ${IMAGE}
	docker run --rm -v ${PWD}/swaggereditor/:/swaggereditor/ ${IMAGE} cp -r /usr/share/nginx/html/. /swaggereditor
	rm -f ./swaggereditor/**/*.map ./swaggereditor/**/.*.png
	sed -i '' -e "s#SwaggerEditorBundle({#SwaggerEditorBundle({url: './openapi.json',#" ./swaggereditor/index.html
