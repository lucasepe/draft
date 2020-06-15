#!/bin/bash

SRC_DIR=../examples
DPI=120
EXE=./dist/draft_linux_amd64/draft

## declare an array of files
declare -a arr=("$SRC_DIR/cli.yml" 
                "$SRC_DIR/ser.yml"
                "$SRC_DIR/msg.yml"
                "$SRC_DIR/gtw.yml"
                "$SRC_DIR/que.yml"
                "$SRC_DIR/fun.yml"
                "$SRC_DIR/rdb.yml"
                "$SRC_DIR/doc.yml"
                "$SRC_DIR/bst.yml"
                "$SRC_DIR/ost.yml"
                "$SRC_DIR/fst.yml"
                "$SRC_DIR/lba.yml"
                "$SRC_DIR/cdn.yml"
                "$SRC_DIR/dns.yml"
                "$SRC_DIR/waf.yml"
                "$SRC_DIR/kub.yml"
                "$SRC_DIR/mem.yml"
                "$SRC_DIR/system-view.yml"
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

   "$EXE" -impl aws "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}_aws.png"
   "$EXE" -impl gcp "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}_gcp.png"
   "$EXE" -impl azure "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}_azure.png"

done