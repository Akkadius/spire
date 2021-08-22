@echo off
goto check_Permissions

:check_Permissions
    echo Administrative permissions required. Detecting permissions...

    net session >nul 2>&1
    if %errorLevel% == 0 (
        echo Success: Administrative permissions confirmed.
    ) else (
        echo Failure: Current permissions inadequate.
        echo.
        echo Please run as administrator
        echo.
        echo Press any key to continue...
        pause >nul
        exit
    )

@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

choco install -y make

choco install -y nodejs-lts

echo %~dp0

xcopy "%~dp0scripts\windows\profile\*.*" "%UserProfile%\" /K /D /H /Y
echo f | xcopy "%~dp0frontend\.env.example" "%~dp0frontend\.env"  /F /Y

"C:\Program Files\Git\git-bash.exe" --cd-to-home

