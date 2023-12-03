# GitOPS Catweb

## Requires
- Flux to monitor manifests/prod folder
- A K8s Cluster with Cert-Manager (cluster issuer) and Ingress installed. You can use the cluster-mgmt project

Before you need to boostrap Flux:
### Export GitLab Token - use a Project access token!

export GITLAB_TOKEN=glpat-XXXX


### In the boostrap Repository:

flux bootstrap gitlab \
  --owner=adess-demos/demo/app-team \
  --repository=flux-config \
  --branch=main \
  --path=./clusters/civo \
  --namespace=flux-system

### App repo
In the new repository, create a deploy token with only the read_repository scope needs to be maintainer

## Add docker secrets to pull

kubectl create secret docker-registry gitlab-registry-credentials \
  --namespace=prod \
  --docker-server=registry.gitlab.com \
  --docker-username=project_52715995_bot_7e551246aa7957ea86f3bcdb7f84b8c1  \
  --docker-password=glpat-XXXX

      imagePullSecrets:
      - name: gitlab-registry-credentials

## Create Secret for the deploy (GitOps Pull) ---- I think this si no longer not needed
flux create secret git flux-deploy-authentication \
         --url=https://gitlab.com/adess-demos/demo/app-team/catweb-gitops/manifests \
         --namespace=default \
         --username=@project_44576698_bot_bbeba7f497d7d33b877aeb6f802bddad \
         --password=glpat-XXXXXX

### To check if the secret was created you can have a look at the K8s resource:
kubectl -n default get secrets flux-deploy-authentication -o yaml



## Add repo
---
apiVersion: source.toolkit.fluxcd.io/v1beta2
kind: GitRepository
metadata:
  name: catweb
  namespace: prod
spec:
  interval: 1m0s
  ref:
    branch: main
  secretRef:
    name: flux-deploy-authentication
  url: https://gitlab.com/adess-demos/demo/app-team/catweb-gitops/manifests


---
apiVersion: source.toolkit.fluxcd.io/v1
kind: GitRepository
metadata:
  name: catweb
  namespace: prod
spec:
  interval: 1m0s
  url: https://gitlab.com/adess-demos/demo/gitops/catweb-gitops/manifests/prod
  ref:
    branch: main