TASKKILL /IM air.exe /F
TASKKILL /IM spire.exe /F
go install github.com/cosmtrek/air@latest

"C:\Program Files\Git\git-bash.exe" -c 'air -c .air.windows.toml'
