cd bkupsrv
go test -race > race.out
sleep 15
cd ..
cd carouselsrv
go test -race > race.out
sleep 15
cd ..
cd contentsrv
go test -race > race.out
sleep 15
cd ..
cd countrysrv
go test -race > race.out
sleep 15
cd ..
cd csssrv
go test -race > race.out
sleep 15
cd ..
cd imgsrv
go test -race > race.out
sleep 15
cd ..
cd mailsrv
go test -race > race.out
sleep 15
cd ..
cd menusrv
go test -race > race.out
sleep 15
cd ..
cd statesrv
go test -race > race.out
sleep 15
cd ..
cd templatesrv
go test -race > race.out
sleep 15
cd ..
cd usersrv
go test -race > race.out
sleep 15
cd ..
cd findfflsrv
go test -race > race.out
sleep 15
cd ..
cd managers
go test -race > race.out
sleep 15
cd ..
cd handlers
go test -race > race.out
sleep 15
cd ..
# cd compresstoken
# go test -coverprofile=coverage.out
# sleep 15
