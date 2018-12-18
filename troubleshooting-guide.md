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
- Suggestion: Run the errand manually and use the `--keep-alive` flag. SSH to Ops Manager VM, run `bosh -d <deployment-id> run-errand <errand-name> --keep-alive`. Then you can use the bosh CLI to `bosh ssh` into the Errand VM and view logs in `/var/vcap/sys/log`. View the BOSH documentation on [BOSH Errands](https://bosh.io/docs/cli-v2/#errand-mgmt) for more info.

### Issue: BOSH Deployment fails with error similar to: "Error: <job name> is not running after update. review logs for failed jobs: <job name>"
- Suggestion 1: Use the bosh CLI to `bosh ssh` [Instructions: Advanced Troubleshooting with BOSH)](https://docs.pivotal.io/pivotalcf/2-3/customizing/trouble-advanced.html#bosh-ssh) into the VM and view logs in `/var/vcap/sys/log/<job-name>/.log`. View the BOSH documentation on [Job Logs](https://bosh.io/docs/job-logs/) for more info.
- Suggestion 2: Another option is to run the failing job start script (often `start.erb`) directly on the VM. SSH into the VM with `bosh ssh`, and run the start script - found in `/var/vcap/jobs/<job-name>/bin/<start-script>`.

### Issue: Operations Manager Web UI has crashed due after staging your Tile (clicking plus button) due to something wrong with Tile Metadata
- Suggestion 1: Use the [OM tool](https://github.com/pivotal-cf/om) to call the Ops Manager API and use the `om unstage-product` command to unstage your Tile. The UI should now be accessible.

## Additional Resources
Here are some more resources on troubleshooting general PCF issues (not specific for Tile Dev):

- [Troubleshooting PCF](http://docs.pivotal.io/pivotalcf/customizing/troubleshooting.html)
- [Troubleshooting Applications](http://docs.pivotal.io/pivotalcf/devguide/deploy-apps/troubleshoot-app-health.html)
- [Advanced Troubleshooting with BOSH](http://docs.pivotal.io/pivotalcf/customizing/trouble-advanced.html)

## Contributing
Please feel free to add to this document with issues you have faced turing Tile Development and suggested troubleshooting steps.
