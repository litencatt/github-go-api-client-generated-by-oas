GENERATOR_VERSION=v6.0.1
openapi-generate-client:
	rm -rf client/github-go/*.go client/github-go/docs/*.md
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:$(GENERATOR_VERSION) generate -i /local/spec/api.github.com.yaml -g go -p packageVersion=0.0.1 -p packageName=github -o /local/client/github-go
	find client/github-go/docs -name "*.md" | xargs sed -i'' -e 's/openapiclient\./github\./g'
	find client/github-go/docs -name "*.md" | xargs sed -i'' -e 's/openapiclient "\.\/openapi"/"github.com\/litencatt\/github-go-api-client-generated-by-oas\/github-go"/g'
	find client/github-go/docs -name "*.md-e" | xargs rm -f
