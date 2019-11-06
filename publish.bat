@echo off

set p=publish

if exist ./%p% (
    rd /s/q %p%
)

md %p%

go build -o go-admin-panel.exe

XCOPY go-admin-panel.exe %p% /s /e
XCOPY config\*.* %p%\config\  /s /e
XCOPY resources\*.* %p%\resources\  /s /e
XCOPY misc\launcher\*.bat %p%\ /s /e
del go-admin-panel.exe

echo "build success!"

@pause