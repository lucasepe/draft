#!/bin/bash

#export DRAFT_ICONS_PATH=../icons
SRC_DIR=../examples
DPI=110
EXE=./dist/draft_linux_amd64/draft

## declare an array of files
declare -a arr=("$SRC_DIR/clients.yml" 
                "$SRC_DIR/networking.yml"
                "$SRC_DIR/compute.yml"
                "$SRC_DIR/database.yml"
                "$SRC_DIR/storage.yml"
                "$SRC_DIR/security.yml"
                "$SRC_DIR/connections.yml" )

## now loop through the above array
for i in "${arr[@]}"
do
   # grab the filename without extension
   filename=$(basename -- "$i")
   # run draft...run!
   "$EXE" -impl -verbose "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}.png"
done

"$EXE" -impl "$SRC_DIR/s3-upload-presigned-url.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/s3-upload-presigned-url.png"
"$EXE" -impl "$SRC_DIR/backend-for-frontend.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/backend-for-frontend.png"
"$EXE" -impl "$SRC_DIR/cognito-custom-auth-flow.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/cognito-custom-auth-flow.png"
"$EXE" -impl "$SRC_DIR/token-manager.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/token-manager.png"
"$EXE" -impl "$SRC_DIR/token-manager-google.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/token-manager-google.png"
"$EXE" -impl "$SRC_DIR/auth0-custom-db-connection-with-jwt.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/auth0-custom-db-connection-with-jwt.png"
