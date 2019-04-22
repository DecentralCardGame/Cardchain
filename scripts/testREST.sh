#!/bin/sh

colorMagenta='\033[0;35m'
colorCyan='\033[0;36m'
colorGreen='\033[0;32m'
colorRed='\033[0;31m'
colorNormal='\033[0m'

# set working directory to script directory
cd "$(dirname "$0")"

# read the json for these REST tests
JSON=$(cat testREST.json)

# iterate over all prevaluations
printf "${colorCyan}executing prevaluation commands ${colorNormal}\n"

length=$(echo ${JSON} | jq -r '.prevaluate' | jq length)

for i in `seq 0 $((length-1))`;
  do
    cmd=$(echo ${JSON} | jq -r ".prevaluate["${i}"]")
    printf "${colorMagenta}${cmd}${colorNormal} \n"
    eval $cmd
  done


#URL="https://localhost:1317/auth/accounts/"
printf "${colorCyan}querying rest-server ${colorNormal}"

while [ -z "$RES" ]; do
  sleep 1

  RES=$(${curlcmd} ${URL}${AliceAddress});
  RES2=$(${curlcmd} ${URL}${BobAddress});

  printf '.'
done

printf '\n'

AliceSequence=$(echo $RES | jq -r  '.value.sequence')
AliceAccountNumber=$(echo $RES | jq -r  '.value.account_number')
BobSequence=$(echo $RES2 | jq -r  '.value.sequence')
BobAccountNumber=$(echo $RES2 | jq -r  '.value.account_number')

#echo $AliceSequence" - "$AliceAccountNumber" - "$BobSequence" - "$BobAccountNumber

# iterate over all TestCases
length=$(echo ${JSON} | jq -r '.TestCases' | jq length)

for i in `seq 0 $((length-1))`;
  do
    # get the ith case
    currentCase=$(echo ${JSON} | jq -r ".TestCases["${i}"]")

    printf "${colorCyan}running test: "$(echo ${currentCase} | jq -r ".name")"${colorNormal}\n"

    # evaluate all .evaluate entries
    length=$(echo ${currentCase} | jq -r '.evaluate' | jq length)

    exp1='.evaluate['
    exp2=']'

    for i in `seq 0 $((length-1))`;
      do
        cmd[${i}]=$(echo ${currentCase} | jq -r "$exp1$i$exp2")
        printf "${colorMagenta}${cmd[${i}]}${colorNormal}\n"
        result[${i}]=$(eval ${cmd[${i}]})
      done

    # fill in %s entries
    currJSON=$(printf "${currentCase}" "${result[@]}")

    REQ=$(echo ${currJSON} | jq -c '.reqBody')
    Route=$(echo ${currJSON} | jq -r '.route')

    #echo $REQ
    #echo $Route

    cmd="curl -XPOST -s --insecure ${Route} --data-binary ${REQ}"
    echo $cmd

    RES=$($cmd)

    echo $RES

    #compare with expected
    cmd="json validate --document-json='${RES}' --schema-json='$(echo ${currJSON} | jq -r '.expected')'"
    echo $cmd

    echo -e "${colorRed}"
    eval $cmd
    if [ "$?" = "0" ]; then
    	echo -e "${colorGreen}Test passed"
    else
    	echo -e "${?}"
    fi
    echo -e "${colorNormal}"

  done

# iterate over all postvaluations
printf "${colorCyan}executing postvaluation commands ${colorNormal}\n"
length=$(echo ${JSON} | jq -r '.postvaluate' | jq length)

for i in `seq 0 $((length-1))`;
  do
    cmd=$(echo ${JSON} | jq -r ".postvaluate["${i}"]")
    echo -e "${colorMagenta}${cmd}${colorNormal}"
    eval $cmd
  done
