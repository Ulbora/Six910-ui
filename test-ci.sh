cd bkupsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd carouselsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd contentsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd countrysrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd csssrv
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
cd menusrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd statesrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd templatesrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd usersrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd findfflsrv
go test -coverprofile=coverage.out
sleep 15
cd ..
cd managers
go test -coverprofile=coverage.out
sleep 15
cd ..
# cd handlers
# go test -coverprofile=coverage.out
# sleep 15
# cd ..
# cd compresstoken
# go test -coverprofile=coverage.out
# sleep 15
