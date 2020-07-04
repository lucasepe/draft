#!/bin/bash

export DRAFT_ICONS_PATH=../icons
SRC_DIR=../examples
DPI=120
EXE=./dist/draft_linux_amd64/draft

## declare an array of files
declare -a arr=("$SRC_DIR/clients.yml" 
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
                "$SRC_DIR/mem.yml" )

## now loop through the above array
for i in "${arr[@]}"
do
   # grab the filename without extension
   filename=$(basename -- "$i")
   # run draft...run!
   "$EXE" -verbose "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}.png"

   "$EXE" -verbose "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}_aws.png"
   "$EXE" -verbose "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}_google.png"
   "$EXE" -verbose "$i" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/${filename%.*}_azure.png"
done

"$EXE" -verbose "$SRC_DIR/system-view.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/system-view.png"
"$EXE" -verbose "$SRC_DIR/system-view.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/system-view-aws.png"

"$EXE" -verbose "$SRC_DIR/impl-example.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/impl-example.png"
"$EXE" -verbose "$SRC_DIR/impl-example.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/impl-example-aws.png"
"$EXE" -verbose "$SRC_DIR/impl-example.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/impl-example-google.png"
"$EXE" -verbose "$SRC_DIR/impl-example.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/impl-example-azure.png"

"$EXE" -verbose "$SRC_DIR/cognito-custom-auth-flow.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/cognito-custom-auth-flow.png"
"$EXE" -verbose "$SRC_DIR/cognito-custom-auth-flow.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/cognito-custom-auth-flow-aws.png"

"$EXE" -verbose "$SRC_DIR/s3-upload-presigned-url.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/s3-upload-presigned-url.png"
"$EXE" -verbose -impl aws "$SRC_DIR/s3-upload-presigned-url.yml" | dot -Tpng -Gdpi=$DPI > "$SRC_DIR/s3-upload-presigned-url-aws.png"

"$EXE" -verbose "$SRC_DIR/demo.yml" | dot -Tpng > "$SRC_DIR/demo.png"
"$EXE" -verbose "$SRC_DIR/demo.yml" | dot -Tpng > "$SRC_DIR/demo-aws.png"
"$EXE" -verbose "$SRC_DIR/demo.yml" | dot -Tpng > "$SRC_DIR/demo-google.png"
"$EXE" -verbose "$SRC_DIR/demo.yml" | dot -Tpng > "$SRC_DIR/demo-azure.png"