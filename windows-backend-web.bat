TASKKILL /IM air.exe /F
TASKKILL /IM spire.exe /F
go install github.com/cosmtrek/air@v1.26.0

"C:\Program Files\Git\git-bash.exe" -c 'air -c .air.windows.toml'
