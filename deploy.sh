#!/usr/bin/env bash
DEP=alexa-flashbriefing-deployment.zip
S3Bucket=cc-go-service
funcName=alexa-flashbriefing-dev

golint ./... && go fmt ./... && go test ./... && GOOS=linux go build -o main cmd/main.go \
&& zip ${DEP} main && rm main && aws s3 cp ${DEP} s3://${S3Bucket}/ \
&& aws lambda update-function-code --function-name ${funcName} --s3-bucket ${S3Bucket} --s3-key ${DEP}