@echo off

start /B cd ./frontend && go run .
start /B cd ./cart && go run .
start /B cd ./checkout && go run .
start /B cd ./email && go run .
start /B cd ./payment && go run .
start /B cd ./product && go run .
start /B cd ./person_infor && go run .
start /B cd ./user && go run .
start /B cd ./order && go run .

pause
