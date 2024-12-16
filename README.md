
# Entro securety test


## How to run:

### Define aws credentials
```
$ export AWS_REGION=region
$ export AWS_ACCESS_KEY_ID=key-id
$ export AWS_SECRET_ACCESS_KEY=key
```


### Setup EKS cluster using terraform

```
$ cd tf
$ terraform init
$ terraform plan
$ terraform apply
```

### Set kubectl to use new cluster:
```
$ aws eks list-clusters 
$ aws eks update-kubeconfig --name gm-eks
```

### Create k8s resources

```
$ kubectl apply -f k8s/deployment.yaml
$ kubectl apply -f k8s/deployment.yaml
$ kubectl describe svc entro-service
```

