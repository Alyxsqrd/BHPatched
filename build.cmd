@echo off

set BUILDS=%~dp0builds
if exist "%BUILDS%" rmdir /s /q "%BUILDS%"

set GOOS=windows
set GOARCH=amd64
go build -trimpath -ldflags "-s -w -H windowsgui" -o "%BUILDS%\patched.exe"

xcopy "%~dp0includes" "%BUILDS%" /e /i
7z x -o"%BUILDS%\workshop" -x!bin "%BUILDS%\workshop.7z" && del "%BUILDS%\workshop.7z"

7z a -mx9 "%BUILDS%\Brick Hill Patched.7z" "%BUILDS%\"

echo. && pause
