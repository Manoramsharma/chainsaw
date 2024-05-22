# Termination graceful period

Some Kubernetes resources can take time before being stopped. For example, deleting a pod can take time if the underlying container doesn't quit quickly enough.

For this reason, Chainsaw provides the `forceTerminationGracePeriod` configuration option and the corresponding `--force-termination-grace-period` flag. If set, Chainsaw will override the `terminationGracePeriodSeconds` when working with the following resource kinds:

- Pod
- Deployment
- StatefulSet
- DaemonSet
- Job
- CronJob

## Configuration

```yaml
apiVersion: chainsaw.kyverno.io/v1alpha2
kind: Configuration
metadata:
  name: example
spec:
  # ...
  forceTerminationGracePeriod: 5s
  # ...
```

## Flag

```bash
chainsaw test --force-termination-grace-period 5s ...
```