---
id: deploy_using_juju
title: Deploy Orchestrator using Juju (Beta)
hide_title: true
---

# Deploy Orchestrator using Juju (Beta)

This how-to guide can be used to deploy Magma's Orchestrator on any cloud environment. It contains
steps to set up a Kubernetes cluster, bootstrap a Juju controller, deploy charmed operators for
Magma Orchestrator and configure DNS A records. For more information on Charmed Magma, please visit
the project's [homepage](https://github.com/canonical/charmed-magma).

> Charmed-Magma is in Beta and is not yet production ready or feature complete.

## Pre-requisites

- Ubuntu 20.04 machine with internet access
- A public domain

## Set up your management environment

From a Ubuntu 20.04 machine, install the following tools:

- [Juju](https://juju.is/docs/olm/installing-juju)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

## Create a Kubernetes cluster and bootstrap a Juju controller

Select a Kubernetes environment and follow the guide to create the cluster and bootstrap
a Juju controller on it.

1. [MicroK8s](https://juju.is/docs/olm/microk8s)
2. [Google Cloud (GKE)](https://juju.is/docs/olm/google-kubernetes-engine-(gke))
3. [Amazon Web Services (EKS)](https://juju.is/docs/olm/amazon-elastic-kubernetes-service-(amazon-eks)#heading--install-the-juju-client)
4. [Microsoft Azure (AKS)](<https://juju.is/docs/olm/azure-kubernetes-service-(azure-aks)>)

## Deploy charmed Magma Orchestrator

From your Ubuntu machine, create an `overlay.yaml` file that contains the following content:

```yaml
applications:
  orc8r-certifier:
    options:
      domain: <your domain name>
  orc8r-nginx:
    options:
      domain: <your domain name>
  tls-certificates-operator:
    options:
      certificate: <your base64 encoded certificate>
      private-key: <your base64 encoded private key>
      ca-certificate: <your base64 encoded ca certificate>
```

Replace the values in angle brackets.

Deploy Orchestrator:

```bash
juju deploy magma-orc8r --overlay overlay.yaml --trust --channel=edge
```

The deployment is completed when all services are in the `Active-Idle` state.

## Import the admin operator HTTPS certificate

Retrieve the PFX package and password that contains the certificates to authenticate against Magma Orchestrator:

```bash
juju scp --container="magma-orc8r-certifier" orc8r-certifier/0:/var/opt/magma/certs/admin_operator.pfx admin_operator.pfx
juju run-action orc8r-certifier/leader get-pfx-package-password --wait
```

> The pfx package was copied to your current working directory and can now be loaded in your browser.

## Setup DNS

Retrieve the services that need to be exposed:

```bash
juju run-action orc8r-orchestrator/leader get-load-balancer-services --wait
```

In your domain registrar, create A records for the following Kubernetes services:

| Hostname                                | Address                                |
|-----------------------------------------|----------------------------------------|
| `bootstrapper-controller.<your domain>` | `<orc8r-bootstrap-nginx External IP>`  |
| `api.<your domain>`                     | `<orc8r-nginx-proxy External IP>`      |
| `controller.<your domain>`              | `<orc8r-clientcert-nginx External IP>` |
| `*.nms.<your domain>`                   | `<nginx-proxy External IP>`            |

## Verify the deployment

Get the master organization's username and password:

```bash
juju run-action nms-magmalte/0 get-master-admin-credentials --wait
```

Confirm successful deployment by visiting `https://master.nms.<your domain>` and logging in
with the `admin-username` and `admin-password` outputted here.
