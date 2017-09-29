#cd utils
#./get_schema.sh
#cp schema.json /tmp
#cd ..

#go build -a -o main *.go
#./main
cp /tmp/vtgraphql_schema/*.go vtgraphql_schema/
cp /tmp/vtgraphql_api/*.go vtgraphql_api/
# Make sure that it works
go build -a -o main *.go
