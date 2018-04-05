#!/bin/bash
set -e


GometalinterVariable=(
           "aligncheck"
           "deadcode"
           "dupl"
           "errcheck"
           "gas"
           "goconst"
           "gocyclo"
           "golint"
           "gosimple"
           "ineffassign"
           "misspell"
           "safesql"
           "staticcheck"
           "structcheck"
           "unconvert"
           "unparam"
           "unused"
           "varcheck"
)


Directory=(
            "adapters"
            "controllers"
            "helper"
            "databases"
            "models"
            "structs"
            "structs/api"
          )

arrayGometalinterVariable=${#GometalinterVariable[@]}
arrayDirectory=${#Directory[@]}


go get -u gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install

for ((k=0; k<${arrayDirectory}; k++));
do
  for ((i=1; i<${arrayGometalinterVariable}; i++));
  do
        if [ "${Directory[$k]}" == "controllers" ] || [ "${Directory[$k]}" == "structs" ] || [ "${Directory[$k]}" == "structs/api" ] 
          then
          if [ "${GometalinterVariable[$i]}" != "gocyclo" ] && [ "${GometalinterVariable[$i]}" != "lll" ] && [ "${GometalinterVariable[$i]}" != "dupl" ] && [ "${GometalinterVariable[$i]}" != "goconst" ]
            then
            echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
            gometalinter.v1 -j 1 --disable-all  --exclude=_test --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
          fi
        else
          echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
          gometalinter.v1 -j 1 --disable-all  --exclude=_test --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
        fi

        sleep 1
        wait

  done
done
