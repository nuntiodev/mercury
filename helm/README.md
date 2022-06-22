 ## Running the Helm Chart
 In order to install the Helm chart, you first need to install cert-manager:
```
 kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.16.1/cert-manager.yaml
```
Then from the project root folder, run ```make helm-install```. If you want to delete the helm chart, run ```make helm-delete``` in the project root folder. Notice that this helm chart requires a valid Hydra Provider in the same namespace.