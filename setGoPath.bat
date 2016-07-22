@echo off
for /f "delims=" %%i in ('pwd') do (set currPath=%%i)
echo currPath = %currPath%
set GOPATH=E:\work\git\hyperagent
echo "GOPATH=%GOPATH%"
cmd /k