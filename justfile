#
# NB: the assumption is this is for a microservice or single use binary
#

# just default
default:
  @just --list

# git justfile


# git save
save JIRA="missing JIRA ticket #" message="yet another commit":
  @git add .
  @git commit -m "{{JIRA}}:{{message}}"

# git push to remote repo
push destination="origin":
  @git push -u {{destination}} 

# git pull
pull:
  @git pull

# golang justfile

# run go code
run:
  @go run .

# build go code
build:
  @go build -o bin  

# test go code
test:
  @go test -v

