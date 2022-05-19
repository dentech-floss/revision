# revision

Provides information about the service (name, version and git commit) that is currently running in the container. This is based on these three env vars that shall be set in the Dockerfile:

```dockerfile
ARG SERVICE_NAME
ENV SERVICE_NAME=$SERVICE_NAME
ARG SERVICE_VERSION
ENV SERVICE_VERSION=$SERVICE_VERSION
ARG GIT_COMMIT
ENV GIT_COMMIT=$GIT_COMMIT
```

So ensure that they are provided as build args in the build pipeline, for example:

```
env:
  DEPLOYMENT_NAME: appointment-service

...

docker build \
--tag "gcr.io/$PROJECT_ID/$IMAGE:${GITHUB_REF:10}" \
--build-arg SERVICE_NAME="$DEPLOYMENT_NAME" \
--build-arg SERVICE_VERSION="${GITHUB_REF:10}" \
--build-arg GIT_COMMIT="$GITHUB_SHA" \
--file=docker/Dockerfile \
.
```

## Install

```
go get github.com/dentech-floss/revision@v0.1.0
```

## Usage

```go
package example

import (
    "github.com/dentech-floss/revision/pkg/revision"
    "github.com/dentech-floss/telemetry/pkg/telemetry"
)

func main() {

    shutdownTracing := telemetry.SetupTracing(
        ctx,
        &telemetry.TracingConfig{
            ServiceName:           revision.ServiceName,
            ServiceVersion:        revision.ServiceVersion,
            DeploymentEnvironment: gcpConfig.ProjectId,
        },
    )
    defer shutdownTracing()
}

```