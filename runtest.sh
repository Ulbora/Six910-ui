cd bkupsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd contentsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd imgsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd mailsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd templatesrv
go test -coverprofile=coverage.out
sleep 15
cd ..
# cd compresstoken
# go test -coverprofile=coverage.out
# sleep 15
