EVMSIGN=/Users/zatan007/go/src/evmsign/evmsign

MESSAGE=`cat action-param.json | sed "s/ACTION/useBeerCoupon/g" | sed "s/TOKEN_ID/1/g" | sed "s/SCHEMA_CODE/mhrs.mhrs1/g"`
echo "MESSAGE: ${MESSAGE}"
BASE64_MESSAGE=`echo -n $MESSAGE | base64 | tr -d '\n'`
echo -n $BASE64_MESSAGE | wc -c
MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | ${EVMSIGN} ./.secret`
# ACTION_SIG=`cat action-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g"`
# BASE64_ACTION_SIG=`echo ${ACTION_SIG} | base64 | tr -d '\n'`

cat <<!
{
    "message": "${BASE64_MESSAGE}",
    "signature": "${MESSAGE_SIG}"
}
!

echo -n ${BASE64_MESSAGE} | ${EVMSIGN} ./.secret 1