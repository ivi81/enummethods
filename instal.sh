VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse HEAD)
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
APP_NAME="github.com/ivi81/enummethods"

echo "Building version: $VERSION"
echo "Commit: $GIT_COMMIT"
echo "Branch: $GIT_BRANCH"

go install -ldflags "-X $APP_NAME/version.Version=$VERSION \
                     -X $APP_NAME/version.BuildTime=$BUILD_TIME \
                     -X $APP_NAME/version.GitCommit=$GIT_COMMIT \
                     -X $APP_NAME/version.GitBranch=$GIT_BRANCH" .