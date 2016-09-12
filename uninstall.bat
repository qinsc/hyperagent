@echo off

echo "Delete Notify Service Autostart reg item"
reg delete "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /v HyperAgent /f

cd %~dp0

echo "Stop Notify Service"
taskkill /f /t /im hyperagenttray.exe 

echo "Stop HyperAgent Service"
hyperagent.exe stop

echo "remove HyperAgent Service"
hyperagent.exe remove

echo "Stop Vnc Service"
nssm.exe stop uvnc_service

pause