# odb-spacebears-adapter-release

This example demonstrates a very basic 
[On Demand Broker](https://github.com/pivotal-cf/on-demand-services-sdk) 
BOSH release of the 
[Spacebears](https://github.com/cf-platform-eng/pcf-examples/blob/master/src/README.md) 
service with its service adapter written in Java. 


## Background

The Spacebears service is a simple name/value pair database designed
for demonstration purposes.  You can create/delete "buckets".  Inside those
buckets, you can add/delete name/value pairs. 

There is a [BOSH release of Spacebears](https://github.com/cf-platform-eng/pcf-examples/tree/master/bosh-simple).  
By using it, you can deploy a single instance of Spacebears in your 
BOSH environment. It spins up one VM that runs Spacebears.

This sample uses the [ODB SDK](https://docs.pivotal.io/svc-sdk/odb/0-21/index.html) 
to take this further. With ODB, you deploy a service broker with the 
Spacebears BOSH release. That service broker gets registered in the 
CF Marketplace. Each time a developer issues a 'cf create-service' command 
requesting a Spacebears instance, ODB will spin up a new VM with 
an instance of Spacebears. 

In order to accomplish creating this sample ODB-driven Spacebears BOSH release, 
we have the following dependecies: 
 - The [BOSH release of Spacebears](https://github.com/cf-platform-eng/pcf-examples/tree/master/bosh-simple)
 - The [BOSH release of ODB](https://github.com/pivotal-cf/on-demand-service-broker-release)
 - The Java SDK. We used [OpenJDK](http://openjdk.java.net/)
 - [JSON-java](https://github.com/stleary/JSON-java) Utility used for JSON parsing.
 
The developer of this Spacebears ODB BOSH release needs to write the following code: 
 - BOSH constructs to include the dependencies
 - Java code and bash script to implement a service adapter
 - A packaging script that compiles and deploys the Java-based service adapter
 - The deployment manifest for this bosh release. 
 
## Build your BOSH release

#### Build the BOSH Release 
    ```
    bosh init-release
    bosh generate-job service-adapter
    bosh generate-package odb-service-adapter
    ```

#### Add the Java SDK to the BOSH Release 
    ```
    wget https://download.java.net/java/jdk8u192/archive/b04/binaries/jdk-8u192-ea-bin-b04-linux-x64-01_aug_2018.tar.gz
    bosh add-blob ../downloads/jdk-8u192-ea-bin-b04-linux-x64-01_aug_2018.tar.gz jdk-8u192-ea-bin-b04-linux-x64-01_aug_2018.tar.gz
    ```

#### Add the json jar to the BOSH Release
    ```
    wget http://central.maven.org/maven2/org/json/json/20180130/json-20180130.jar
    bosh add-blob ../downloads/json-20180130.jar json-20180130.jar
    ```

#### Modify the job spec
The only purpose of the service-adapter job in this sample is to create a dependency on the
odb-service-adapter package, so the package gets put on disk where it is needed. 
[Example](https://github.com/cf-platform-eng/pcf-examples/blob/master/odb-spacebears-adapter-release/jobs/service-adapter/spec)

#### Modify the package
Update the spec file to include the file dependencies from the blobs and src. 
[Example](https://github.com/cf-platform-eng/pcf-examples/blob/master/odb-spacebears-adapter-release/packages/odb-service-adapter/spec)

Update the packaging script to compile the adapter and place it in the correct locations. 
[Example](https://github.com/cf-platform-eng/pcf-examples/blob/master/odb-spacebears-adapter-release/packages/odb-service-adapter/packaging)  

#### Create the service adapter
  1) Fetch the deployment manifest for the stand-alone Spacebears bosh release.  
     It can be found [here](https://github.com/cf-platform-eng/pcf-examples/blob/master/bosh-simple/manifests/lite_manifest.yml)
  2) Encode that into the [Java adapter](https://github.com/cf-platform-eng/pcf-examples/blob/master/odb-spacebears-adapter-release/src/example/OdbSpacebearsAdapter.java). 
  3) Write a bash wrapper around the Java main program so it complies to the 
     executable path.  [Example](https://github.com/cf-platform-eng/pcf-examples/blob/master/odb-spacebears-adapter-release/src/code.sh)
  
#### Create the deployment manifest
[Example](https://github.com/cf-platform-eng/pcf-examples/blob/master/odb-spacebears-adapter-release/spacebears_manifest.yml)
If you copy the example, you may need to: 
  - tweek the name of your release
  - use the stemcell that comes with the build of cf dev you are using. 
  - update the credentials for your cf dev instance 
  
  
## Deploy it to cf dev  

These are the steps you'll need to do to get it all running locally. 
  1) Use the [cf dev readme](https://github.com/cloudfoundry-incubator/cfdev) to deploy cf dev 
     to your workstation.  cf dev comes with both a bosh director and cloud foundry making it 
     handy for local bosh development.  
  2) Upload the Spacebears release
  3) Upload the ODB SDK release
  4) Deploy your BOSH Release
  
 #### Upload the Spacebears release
    ```
    git clone https://github.com/cf-platform-eng/pcf-examples.git
    cd pcf-examples/bosh-simple
    ./deploy.sh
    bosh upload-release
    bosh -d bosh_simple delete-deployment
    ```
    The delete-deployment is there simply becuase the deployment is included 
    in the deploy.sh and we don't need it. 
 
 #### Upload the ODB SDK Release
  Get the ODB SDK from Pivotal Network. 
  [url](https://network.pivotal.io/api/v2/products/on-demand-services-sdk/releases/97789/product_files/135987/download)

    ```
    bosh upload-release on-demand-service-broker-release-v0.21.2.tar.gz
    ```

#### Deploy your BOSH Release 
```
bosh create-release --force
bosh upload-release
bosh -d spacebears-on-demand-broker-dev deploy spacebears_manifest.yml
```
 
## Exercise it

Tell cf about the broker:  
```
cf create-service-broker sb broker broker http://10.144.0.144:8080
```


Enable access so it can show up in the marketplace.   Note that that name that enable-service-access wants is the name from the catalog, not the name I gave it on the create. 
```
cf service-brokers
cf enable-service-access spacebears-dev
```

Create a service instance: 
```
cf create-service spacebears-dev dedicated-vm joes-sb
```

To see the bind response, you can use a service key: 
```
cf create-service-key sb-service sb-service-key
cf service-key sb-service sb-service-key
```

Ssh to one of the nodes in the bosh network: 
```
bosh vms
bosh -d spacebears-on-demand-broker-dev ssh
```

Then, issue curl commands against the Spacebears api: 
```
export host=10.144.0.145:9000
curl -u admin:symphony27_Trailers -X POST http://$host/api/admin/bucket/my_bucket
curl -u admin:symphony27_Trailers -X PUT http://$host/api/admin/bucket/my_bucket/credentials -d '{"username": "user","password": "pass"}'
curl -u user:pass -X GET http://$host/api/bucket/my_bucket/
curl -u user:pass -X PUT http://$host/api/bucket/my_bucket/some_val -d 'Test value'
curl -u user:pass -X GET http://$host/api/bucket/my_bucket/
{"some_val":"Test value"}
```