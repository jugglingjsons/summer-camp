# summer-camp
An application task for the summer-camp by kiwi.com.
Made to be extended wth other microservices in the future.
Assumes you have a kubectl context configured for a project that has kubernetees support activated.
# Pre-requisites
Install:
- google-cloud-sdk
- kubectl
- terraform (v0.11.14) - you can use tfswitch to manage those
## On MacOs:
cat <<-"BREWFILE" > Brewfile
cask 'google-cloud-sdk'
brew 'kubectl'
brew 'terraform'
BREWFILE
brew bundle --verbose
# Usage:
- Specify the GCP app to be used for the 
```
export TF_VAR_project="$(gcloud config get-value project -q)"
export TF_VAR_region="us-east1-b"
export TF_VAR_user="admin"
export TF_VAR_password="" // put some random password in
export DOCKER_HUB_USER="" // provide your username
```
- Initialize terraform & deploy the app to kubernetes onto the context provided in the cli context.
```
terraform init
terraform plan
terraform apply
```
Check service addresses - granted you have the same context configured.
```
kubectl get service
```
# Development:
- recommended: run the app with docker-compose up from the root of the project. Mind you it will take a while to build if there are more microservces present.
# Usage:
## Available endpoints:
### Locally:
- curl -d '{"payload":"bazinga"}' -H "Content-Type: application/json" -X POST http://localhost:8800/create
Should respond with approximately:
```
{
    "response": "Number of invocations for bazinga2, is exactly 1",
    "err": null
}
```
- curl localhost:8800/health,
-  responds with:
```
{
    "ok": true
}
``` and status 200
### Online:
- curl -d '{"payload":"bazinga", "id": "123"}' -H "Content-Type: application/json" -X POST http://localhost:8800/create
- curl localhost:8800/health
## Deployment
It will push to two image hosting repositories for now. Both gcp - requried for the terraform build on gcp and docker hub - to allow for public image sharing.
```
cd summer_app
make build
```
## Attributions:
- Alex Pliutau @packagemain
- go-kit creators
- Joaquin Menchaca
as their work many concepts use here much clearer to me.