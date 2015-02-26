echo off
cd "%~dp0"

set GOPATH=%~dp0_vendor
REM echo %GOPATH%

dev_appserver.py ^
	app.yaml ^
	--host 0.0.0.0 ^
	--enable_sendmail

pause > nul