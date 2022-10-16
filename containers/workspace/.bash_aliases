alias cmd='rm internal/http/spa/*packr.go && go run main.go'
alias c='cmd'
alias tests='go test -count=1 ./...'
alias test='tests'
alias coverage='go test -count=1 ./... | grep -v "no test files"'
alias awslocal='aws --endpoint-url=http://localstack:4566'
alias dynamo='aws --endpoint-url http://dynamodb:8000 dynamodb'
alias create-migration='./scripts/migration-create.sh'
alias up-migration='./scripts/migration-up.sh'
alias force-migration='./scripts/migration-force.sh'
alias down-migration='./scripts/migration-down.sh'
alias mc='~/src/scripts/mc.sh'
alias rc='redis-cli -h redis'
alias ?='/opt/terminal-help.sh'

alias s='./scripts/generate-swagger.sh'

/opt/terminal-help.sh

#############################################
# run wire when inject files change
#############################################
echo "Clearing excess [wire] watcher(s) and starting a new one"
pkill -f "./boot"
while inotifywait -qqre modify "./boot"; do echo "Running wire..." && cd boot && wire && cd ../; done &
