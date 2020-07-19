cd bkupsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd contsrv
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
cd tmptsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
# cd compresstoken
# go test -coverprofile=coverage.out
# sleep 15
