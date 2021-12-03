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

:: Install Choco
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

:: Install Make
choco install -y make

:: Install NodeJS
:: choco install -y nodejs-lts
choco install -y nodejs --version=12.13.0

:: Install Go (If want to use outside of docker)
choco install -y golang

choco install -y unzip

:: Copy windows bashrc profiles and WSL2 settings
xcopy "%~dp0scripts\windows\profile\*.*" "%UserProfile%\" /K /D /H /Y

:: Copy .env vars
echo f | xcopy "%~dp0frontend\.env.example.windows" "%~dp0frontend\.env"  /F /Y
echo f | xcopy "%~dp0.env.dev" "%~dp0.env"  /F /Y

:: Launch Git Bash (MinGW)
"C:\Program Files\Git\git-bash.exe" --cd-to-home
