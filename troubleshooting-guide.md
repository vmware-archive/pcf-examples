# Troubleshooting Help

Sooner or later you will run into problems that require digging a little bit deeper.

This document will guide you through some of the common issues Tile Developers run into and troubleshooting suggestions.

## Prerequisites

It is important to become familiar with the BOSH CLI and using SSH to connect to the Ops Manager VM and authenticate with the BOSH Director. Instructions on that can be found in [Advanced Troubleshooting with BOSH](http://docs.pivotal.io/pivotalcf/customizing/trouble-advanced.html).

## Commonly Reported Issues and Troubleshooting Suggestions

### Issue: Ops Manager reports install failed after hitting Apply Changes. (AKA "Exit code -1" Error)
- Suggestion 1: There is a more meaningful error hiding in there! Further error logs can be found by ssh-ing into the Ops Manager VM [(Instructions: Advanced Troubleshooting with BOSH)](http://docs.pivotal.io/pivotalcf/customizing/trouble-advanced.html) and view the Ops Manager logs: `less /var/log/opsmanager/production.log`

- Suggestion 2: SSH to Ops Manager VM [(Instructions: Advanced Troubleshooting with BOSH)](http://docs.pivotal.io/pivotalcf/customizing/trouble-advanced.html) and run command `bosh tasks --recent`, look for the relevant (probably failed) task. Run command `bosh tasks <id> --debug` for verbose logs.

### Issue: BOSH Errand in my deployment fails, and then deletes the VM where Errand is run, so I'm having trouble troubleshooting or look into logs.
- Error: `Errand <errand-name> completed with error (exit code 1)`
- Suggestion: Run the errand manually and use the `--keep-alive` flag. From the Ops Manager VM, run `bosh -d <deployment-id> run-errand <errand-name> --keep-alive`. Then you can use the bosh CLI to `bosh ssh` into the Errand VM and view logs in `/var/vcap/sys/log`. View the BOSH documentation on [BOSH SSH](https://bosh.io/docs/sysadmin-commands/#ssh) and [BOSH Errands](https://bosh.io/docs/cli-v2/#errand-mgmt) for more info.

### Issue: BOSH Deployment fails because there is a problem in one of your job start or pre-start scripts.
- Error: `Error: <job name> is not running after update. review logs for failed jobs: <job name>`
- Error: `Action Failed get_task: Task <id> result: 2 of 3 pre-start scripts failed. Failed Jobs: <job-name>, <job-name>. Successful Jobs: <job-name>`
- Suggestion 1: Use the bosh CLI to `bosh ssh` [Instructions: Advanced Troubleshooting with BOSH)](https://docs.pivotal.io/pivotalcf/2-3/customizing/trouble-advanced.html#bosh-ssh) into the VM and view logs in `/var/vcap/sys/log/<job-name>/.log`. View the BOSH documentation on [Job Logs](https://bosh.io/docs/job-logs/) for more info.
*make sure files look like they are meant to (pre-start)*
- Suggestion 2: Another option is to run the failing job start script (often `start.erb`) directly on the VM. SSH into the VM with `bosh ssh`, and run the start script - found in `/var/vcap/jobs/<job-name>/bin/<start-script>`.

### Issue: Operations Manager Web UI has crashed due after staging your Tile (clicking plus button) due to something wrong with Tile Metadata
- Suggestion 1: Use the [OM tool](https://github.com/pivotal-cf/om) to call the Ops Manager API and use the `om unstage-product` command to unstage your Tile. The UI should now be accessible.
*often the reason for this is a bosh add on with no form*

### Issue: Can't SSH into a custom deployment with a Windows stemcell.
- Error: `Error: Action Failed ssh: Getting host public key: OpenSSH is not running: sshd service not running and start type is disabled.  To enable ssh on Windows you must run the enable_ssh job from the windows-utilities-release.`
- Suggestion: You need to augment your deployment manifest with the `windows-utilities` that enable SSH for Windows VMs:

In your `deployment-manifest.yml` under
```
instance_groups:
  jobs:
```
add :

```
      - name: enable_rdp
        release: windows-utilities
        consumes: {}
        provides: {}
        properties:
          enable_rdp:
            enabled: true
      - name: enable_ssh
        release: windows-utilities
        consumes: {}
        provides: {}
        properties:
          enable_ssh:
            enabled: true
```
under `releases:`
add:
```
  - name: windows-utilities
    version: latest
  ```
Redeploy, then you will now be able to ssh into the VM with `bosh ssh`.

## Additional Resources
Here are some more resources on troubleshooting general PCF issues (not specific for Tile Dev):

- [Troubleshooting PCF](http://docs.pivotal.io/pivotalcf/customizing/troubleshooting.html)
- [Troubleshooting Applications](http://docs.pivotal.io/pivotalcf/devguide/deploy-apps/troubleshoot-app-health.html)
- [Advanced Troubleshooting with BOSH](http://docs.pivotal.io/pivotalcf/customizing/trouble-advanced.html)

## Contributing
Please feel free to add to this document with issues you have faced turing Tile Development and suggested troubleshooting steps.
