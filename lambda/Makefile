build:
	GOOS=linux go build -ldflags="-s -w" -o main main.go && zip deployment.zip main

deploy:
	aws lambda create-function \
	--region ap-southeast-1 \
	--function-name DiscoverCards \
	--zip-file fileb://./deployment.zip \
	--runtime go1.x \
	--role $(AWS_ROLE) \
	--handler main

redeploy:
	aws lambda update-function-code \
	--region ap-southeast-1 \
	--function-name DiscoverCards \
	--zip-file fileb://./deployment.zip