@echo off

echo "Add Notify Service Autostart reg item"
reg add "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /v HyperAgent /t REG_SZ /f /d "%~dp0startTray.vbs"

cd %~dp0

echo "Copy XCGUI.dll to C:\Windows\SysWOW64"
cp -f  XCGUI.dll C:\Windows\SysWOW64\XCGUI.dll

echo "Copy nssm.exe to C:\Windows\SysWOW64"
cp -f nssm.exe C:\Windows\SysWOW64\nssm.exe

echo "Start Notify Service: %~dp0startTray.vbs"
cd 
start /b startTray.vbs

echo "Install HyperAgent Service"
hyperagent.exe install

echo "Start HyperAgent Service"
hyperagent.exe ^start

echo "Start Vnc Service"
nssm.exe ^start uvnc_service

pause

