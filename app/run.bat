@echo off

start cmd /K "cd ./frontend && go run ."
start cmd /K "cd ./cart && go run ."
start cmd /K "cd ./checkout && go run ."
start cmd /K "cd ./email && go run ."
start cmd /K "cd ./payment && go run ."
start cmd /K "cd ./product && go run ."
start cmd /K "cd ./person_infor && go run ."
start cmd /K "cd ./user && go run ."
start cmd /K "cd ./order && go run ."

pause
