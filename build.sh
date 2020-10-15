rm -rf buildfiles
mkdir buildfiles

go run cmd/ducclang/main.go --f $1 --o buildfiles/output.go
echo 'module build\ngo 1.15' > buildfiles/go.mod

cd buildfiles
go mod edit -require github.com/ducc/lang@master
go mod vendor

go build output.go
mv output ..

