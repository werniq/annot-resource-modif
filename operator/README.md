# operator
The **Resource Modifier Controller** is a Kubernetes operator built with Kubebuilder. 
It allows users to modify existing Kubernetes resources based on custom-defined rules 
and annotations provided through a Custom Resource Definition.

### Key Features
- Modify various Kubernetes resource types (e.g., Pods, Deployments, Nodes).
- Use annotations to define specific actions for resource modification.
- Examples of supported annotations:
    - `removeAnyFinalizers`: Removes all finalizers from a resource.
    - `addLabel:<key>:<value>`: Adds a label to a resource.
    - `removeLabel:<key>`: Removes a label from a resource.
    - `scale:<replicas>`: Scales a scalable resource like a Deployment.
    - `updateImage:<containerName>:<newImage>`: Updates the container image.

### Supported Annotations

1. `removeAnyFinalizer` - deletes all finalizers from a resource, to allow it's deletion.
2. `addLabel:<key>:<value>` - adds a specific label to the resource
3. `removeLabel:<key>` - removes specific label from resource
4. `addAnnotation:<key>:<value>` - adds specific annotation 
5. `removeAnnotation:<key>` - removes annotation by it's key
6. `scale:<replicas>` - Scale the resource (e.g. Deployment), if possible
7. `restart` - Restart the resource (if applicable, e.g. Pod or Deployment)
8. `taint:<key>:<value>:<effect>` - Applies a taint to a Node resource
9. `toleration:<key>:<value>:<effect>` - Adds a toleration to the resource
10. `setResourceLimit:<cpu>:<memory>` - Sets CPU and memory limit for the resource (e.g. Pod or Container). Use `default` keyword if You don't wish to modify resource limit. Example: `200m:default`, or `default:500Mi`
11. `setResourceRequest:<cpu>:<memory>` - Set CPU and memory Request, if applicable. Same rules as in `setResourceLimit`
12. `addEnvironmentVariable:<name>:<value>` - Adds an environment variable to a container in a Pod.
13. `deleteResource` - Entirely deletes the resource.
14. `updateImage:<containerName>:<image>` - Update the image of specific container.
15. `addVolume:<volumeName>` - Adds a volume to a Pod or Deployment.
16. `removeVolume:<volumeName>` - Removes a volume
17. `patch:<jsonPath>:<value>` - Apply a json patch to the resource.
18. `addOwnerReference:<kind>:<name>:<uid>` - Add new owner reference
19. `cordonNode` - Mark node as unschedulable.
20. `uncordonNode` - Mark node as schedulable.
21. `evictPods` - Evict all pods running on the Node.
22. `addAffinity:<type>:<key>:<operator>:<value>` - Adds affinity rules to Pod or Deployment.
23. `setServiceType:<type>` - Updates a type of service
24. `setIngressHost:<host>` - Updates the host field in an Ingress.
25. `addConfigMapRef:<name>:<path>` - Mounts a ConfigMap as a volume of a Pod.



## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```
