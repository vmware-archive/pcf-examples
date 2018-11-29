# Suggested development plan for BOSH Add-Ons

Developing a BOSH add-on from scratch can be challenging, especially if not trying to do this iteratively.  This document has a suggested plan to developing an add-on, which may make it simpler to get started.

Note that you cannot use BOSH Lite for developing add-ons, since the way VMs are provisioned in VirtualBox don't match well enough to a full BOSH environment.  It's recommended to use our Platform Engineering PIE environments.

Here are the steps we'll do:

1. Prove the software works as expected on a BOSH-managed VM
2. Create a BOSH release and deploy with a manifest
3. Add configuration properties
4. Make a runtime configuration to test as an add-on
5. Make a tile.yml that wraps your deployment
6. Make a tile.yml that wraps your runtime configuration

## Prove your software

The first step would be to check that your software will work on its own, before being deployed by BOSH.

1. First, you'll need a VM for your software to be installed on.  Just deploy a sample BOSH deployment to get a VM instance.  Our [BOSH Sample](https://github.com/cf-platform-eng/pcf-examples/tree/master/bosh-simple) project will work just fine.
2. SSH into the VM:
    1. `bosh vms --deployment=bosh_simple`
    2. Find a VM instance in your deployment
    3. `bosh ssh --deployment=bosh_simple bosh_simple/11111111-2222-3333-4444-555555555555`
3. Download your software package to the VM.  Some options for doing so:
    1. Use curl or wget to fetch a .tgz
    2. Use apt-get to install an Ubuntu package
    3. Others?
4. Configure as necessary
    1. Taking note of the steps needed as you'll need to make this a script later
    2. What parts can be put into the script?
    3. What parts require user configuration?
        1. Software options
        2. Server locations
        3. License keys
        4. Credentials (Usernames and passwords)
5. Test some basic functionality

After this point, if everything is working as expected, you should be ready to start building a BOSH release.

## Create BOSH packages and jobs and deploy with a manifest

Your first task should be trying to create a BOSH release to deploy your software on its own.

[This page](https://bosh.io/docs/create-release) has good information about creating a BOSH release, including the packages and jobs that are included.

### Build a BOSH package

A package tells BOSH how to put your software onto the VM.  You can build and install software from source, or unpack and install binary distributions.

#### Source packages ([BOSH docs](https://bosh.io/docs/create-release/#source))

If using a source-based package, put your source inside of the `src` directory and use the `packaging` script to compile and install it.

#### Binary packages ([BOSH docs](https://bosh.io/docs/compiled-releases/))

If using a binary-based pacakge, add your binaries into the local blob store and use the `packaging` script to unpack and install it.  See the [BOSH Sample](https://github.com/cf-platform-eng/pcf-examples/tree/master/bosh-simple)'s deploy.sh script to see how binary blobs are added and referenced in the packages.

### Build BOSH jobs ([BOSH docs](https://bosh.io/docs/create-release/#job-skel))

You should have a BOSH job for each service or application you have.  It includes a `spec` file for defining which package(s) it uses, template files to use, and properties to customize behavior.

### Create a BOSH release

When you have all of the above, create a versioned release with `bosh create-release --force`

### Deploy with a manifest file

Finally, create a simple deployment manifest file and use BOSH to deploy your software into its own VM.  By not using a runtime configuration (and making this an add-on) you make it much faster to test updates to your packages, jobs and releases.  Start with something like this as a stub (replacing anything inside of `<>`'s):

```yaml
---
name: <deployment-name>

releases:
- name: <release-name>
  version: latest

stemcells:
- alias: "default"
  os: "ubuntu-xenial"
  version: latest

update:
  canaries: 1
  max_in_flight: 2
  canary_watch_time: 1000-30000
  update_watch_time: 1000-30000

instance_groups:
- name: <instance-group-name>
  vm_type: default
  stemcell: default
  instances: 1
  azs: [z1]
  networks:
  - name: default
  jobs:
  - name: <job-name>
    release: <release-name>
```

Now, test the deployment by using `bosh deploy --deployment <deployment-name>`

## Add configuration properties ([BOSH docs](https://bosh.io/docs/props-common/))

To pass user-configurable values into your software (license keys, credentials, parameters, etc...) you will set properties.  It's best to do this a few at a time.  If one causes a problem, it will be easier to track down.

Properties are defined in the `properties` section of the job's `spec` file ([BOSH docs](https://bosh.io/docs/jobs/#spec)).  When a property is used in a template file, it will use the property's value in that template during deployment.  For example, in this JSON config file:

```json
$ cat jobs/<job-name>/templates/config.json.erb
{
    "license_key": "<%= p(license_key) %>",
    "feature_a": {
        "enabled": <%= p("feature_a.enabled"), false %>
    },
    ...
}
```

Then in the job's `spec` file you can have:

```yaml
templates:
  config.json.erb: config/config.json
properties:
  feature_a.enabled:
    description: Should feature A be enabled?
    default: false
  license_key:
    description: License key for this service
```

Finally, you can set the values for these properties in the deployment manifest ([BOSH docs](https://bosh.io/docs/manifest-v2/#instance-groups)):

```yaml
...
instance_groups:
- name: <instance-group-name>
  ...
  jobs:
  - name: <job-name>
    release: <release-name>
    properties:
      license_key: Super-secret-license-key
      feature_a:
        enabled: true
```

## Build BOSH runtime config

When your software is deploying and running successfully as a BOSH deployment, it's time to change it into an add-on.  You define an add-on inside of a runtime config.  This tells BOSH when to include your release inside of other deployments.

The runtime configuration is very similar to a deployment manifest with a few exceptions:

1. There is no specifictations for requirements of the VM (stemcell definition, network, persistent disk, etc...).  This is because it's not your add-on which determines that, but the other deployment that specifies the VM that will be provisioned.
2. You get to specify inclusion and exclusion criteria ([BOSH docs](https://bosh.io/docs/runtime-config/#placement-rules)).  This is how BOSH will determine which deployments will include your add-on.
3. You must specify a version for your release.  No using `latest` anymore.

Here's an example runtime config for deploying an add-on:

```yaml
releases:
- name: <release-name>
  version: 0+dev.1

addons:
- name: <addon-name>
  jobs:
  - name: <job-name>
    release: <release-name>
    properties:
      license_key: Super-secret-license-key
      feature_a:
        enabled: true
  include:
    deployments: my-canary-deployment
  exclude:
    stemcell:
    - os: [windows1803, windows2012R2, windows2016]
```

When testing your add-on, it's best to specify an inclusion rule that targets a specific deployment.  One that you control.  This way, you're sure that your add-on will not be added to other deployments, which is nice when on a shared an environment such as our PIE environments.

Set your runtime config using `bosh update-runtime-config`, and then deploy (or re-deploy) another bosh deployment that matches your inclusion rules.

## Build tile.yml based on runtime config

When everything is looking good with a direct BOSH deployment, it's time to wrap it in a tile.

In a new directory, use tile-generator to create a stub tile.yml file: `tile init` ([Docs](https://docs.pivotal.io/tiledev/2-3/tile-generator.html)).  The contents of your runtime configuration manifest will nearly all be copied into the tile.yml.  This is also where you can specify the forms that the PCF Operators will use to set your properties.  Here's an example tile.yml for an add-on:

```yaml
name: <tile-name>
icon_file: resources/icon.png
label: Brief Text for the Tile Icon
description: Longer description of the tile's purpose

forms:
- name: service_properties
  label: Service Properties
  description: Properties for configuring the service
  properties:
  - name: license_key
    type: string
    label: License Key
    description: The license key for the service
- name: feature_a_properties
  label: Feature A
  properties:
  - name: feature_a.enabled
    type: boolean
    label: Enable
    description: Should Feature A be enabled?

packages:
- name: <package-name>
  type: bosh-release
  NEED MORE ACCURATE DETAILS HERE
runtime_configs:
- name: <runtime-config-name>
  runtime_config:
    releases:
    - name: <release-name>
      version: 0+dev.1
    addons:
    - name: <addon-name>
      jobs:
      - name: <job-name>
        release: <release-name>
        properties:
          license_key: (( .properties.license_key ))
          feature_a:
            enabled: (( .properties.feature_a.enabled ))
      include:
        deployments: my-canary-deployment
      exclude:
        stemcell:
        - os: [windows1803, windows2012R2, windows2016]
```

Run `tile build`.  If all succeeds, you will find a `.pivotal` file that can be uploaded to Ops Manager, staged, configured and deployed.