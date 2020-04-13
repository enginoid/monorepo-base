
 - [ ] Talk about `bazel run //cmd/setup`

```
# Set up a project and select it.
gcloud init
gcloud auth application-default login

gcloud auth configure-docker

(cd ./terraform; terraform init -input=false)
(cd ./terraform; terraform apply -input=false -auto-approve)

# gcloud container clusters describe --format json $(cd terraform && terraform output cluster_name)
gcloud container clusters get-credentials \
    monorepo-base-staging \
    --zone europe-west2-c \
    --project monorepo-base
```