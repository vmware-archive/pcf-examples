## bosh-simple-addon

This contains the bare minimum for deploying addons using a BOSH runtime
config. To simplify the process, this uses a lightweight release that is
downloadable by BOSH.

The `login_banner` job inside of the `os-conf` release will show text when
ssh'ing into any VM where this addon is installed.

This addon is installed onto any VM where the Spacebears job exists (as
installed from the other sample deployments and tiles in this project)

### Setting up the runtime config

```bash
bosh update-runtime-config --name=bosh_simple_runtime_config runtime-config.yml
```

NOTE: After updating the runtime config, any deployments that match the
inclusion criteria will not automatically include the addon.  The addon
will be included only when those deployments are redeployed.

### Deploy

Deploy (or re-deploy) a Spacebears deployment or tile.

### Test

SSH into a Spacebears vm:

```bash
ubuntu@ops-manager:~$ bosh ssh -d bosh_simple
...
Welcome to a tile-based Spacebears VM!
...
spacebears_db_node/e76315b7-dea8-49ac-a0b6-8c9398533c6e:~$
```

### Cleanup

```bash
bosh delete-config --name=bosh_simple_runtime_config
```
