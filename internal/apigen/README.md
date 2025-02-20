# RHOAS API Generator Tool

Internal tools for generating API clients

## Fetch OpenAPI file

Install:

```shell
go get -u github.com/redhat-developer/app-services-sdk-go/internal/apigen/cmd/api-fetch@latest
```

Usage:

```shell
api-fetch --download-url="https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml" --token="$GITHUB_ACCESS_TOKEN"
```

### Generate API client

Install:

```shell
go get -u github.com/redhat-developer/app-services-sdk-go/internal/apigen/cmd/api-generate@latest
```

Usage:

```shell
api-generate --client-id="kafka-mgmt/v1" --generator=go --templates-dir="./scripts/templates" --repo-metadata ".config/api-client-metadata.json"
```

### Generate all API clients

Install:

```shell
go get -u github.com/redhat-developer/app-services-sdk-go/internal/apigen/cmd/api-generate-all@latest
```

Usage:

```shell
api-generate-all --repo-metadata=.config/api-client-metadata.json --generator go --templates-dir=./scripts/templates
```

### Local development

Execute command to install all  versions of the commandline directly from source code.

```shell
make install
```