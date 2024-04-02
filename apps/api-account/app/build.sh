APPNAME="api-account"

VERSION=6
AWS_VERSION=$VERSION
DOCKER_IMAGE="$APPNAME:$VERSION"
AWS_ENDPOINT="301261400751.dkr.ecr.us-east-1.amazonaws.com"


echo "iniciando build $APPNAME com versão $VERSION"

echo $DOCKER_IMAGE
go mod tidy
go mod vendor
GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static'" -o appbin ./app.go

echo "iniciando build $APPNAME com versão $VERSION"

echo $DOCKER_IMAGE

export DOCKER_IMAGE=$DOCKER_IMAGE
docker build -t ${DOCKER_IMAGE} .

aws ecr get-login-password --region us-east-1 \
  | docker login --username AWS --password-stdin "$AWS_ENDPOINT"

docker tag ${DOCKER_IMAGE} "$AWS_ENDPOINT/$APPNAME:$AWS_VERSION"
docker push "$AWS_ENDPOINT/$APPNAME:$AWS_VERSION"


kubectl apply -f ../