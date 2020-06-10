#!/bin/bash

SRC_DIR=../examples
DPI=120
EXE=./dist/draft_linux_amd64/draft

## declare an array of files
declare -a arr=("$SRC_DIR/client.yml" 
                "$SRC_DIR/service.yml"
                "$SRC_DIR/broker.yml"
                "$SRC_DIR/gateway.yml"
                "$SRC_DIR/queue.yml"
                "$SRC_DIR/function.yml"
                "$SRC_DIR/database.yml"
                "$SRC_DIR/storage.yml"
                "$SRC_DIR/balancer.yml"
                "$SRC_DIR/cdn.yml"
                "$SRC_DIR/dns.yml"
                "$SRC_DIR/custom_image.yml"
                "$SRC_DIR/message-bus-pattern.yml"
                "$SRC_DIR/aws-cognito-custom-auth-flow.yml"
                "$SRC_DIR/s3-upload-presigned-url.yml"
                "$SRC_DIR/backend-for-frontend.yml" )

## now loop through the above array
for i in "${arr[@]}"
do
   # grab the filename without extension
   filename=$(basename -- "$i")
   # run draft...run!
   "$EXE" "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}.png"
done