# p6s 

p6s (projects) is an Operator built to recreate the functionality of OpenShift's Project resource in vanilla Kubernetes.

Simply create a [Project](config/samples/p6s_v1alpha1_project.yaml) and p6s will create a matching namespace and create
the objects in the [ProjectTemplate](config/samples/p6s_v1alpha1_projecttemplate.yaml).

## Install

At this stage only the dev deploy is supported. This may require installing other tools.

This will deploy the CRDs and control-plane deployment in the p6s-system namespace.

```sh
make deploy
```

Create a `ProjectTemplate` called `default` in the `p6s-system` namespace to declare resources to be created upon
project request.

At this time only one `ProjectTemplate` is supported.

You should refine the permissions granted in the `p6s-manager-role` to allow it to create the resources it needs to
and remove **only** the wildcard permissions.
