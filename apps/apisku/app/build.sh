
VERSION=3
AWS_VERSION=$VERSION

AWS_ENDPOINT="301261400751.dkr.ecr.us-east-1.amazonaws.com"
APPNAME=api-sku
echo "iniciando build $APPNAME com versão $VERSION"
DOCKER_IMAGE="$APPNAME:$VERSION"

./gradlew build

echo $DOCKER_IMAGE

export DOCKER_IMAGE=$DOCKER_IMAGE
docker build -t ${DOCKER_IMAGE} .

aws ecr get-login-password --region us-east-1 \
  | docker login --username AWS --password-stdin "$AWS_ENDPOINT"

docker images
docker tag ${DOCKER_IMAGE} "$AWS_ENDPOINT/$APPNAME:$AWS_VERSION"
docker push "$AWS_ENDPOINT/$APPNAME:$AWS_VERSION"

