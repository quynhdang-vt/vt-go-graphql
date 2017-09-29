#!/bin/bash
set -x

#TODO -- SET your own TOKEN and GraphQL Endpoint
token=$DEV_TOKEN
api_url=$DEV_GRAPHQL_ENDPOINT

#api_url=http://localhost:3000/graphql

infile=$1
outfile=$2

if [ $infile=="" ];then
  infile=query_schema.json
fi

if [ $outfile=="" ];then
  outfile=schema.json
fi

echo "Retrieving schema from $api_url with queryin $infile, output to $outfile..."
outfile1=s.json
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $token" -d @$infile -o $outfile1 $api_url

cat $outfile1 | jq '' > $outfile
rm -f $outfile1


