#!/bin/bash
set -e

export JAVA_HOME=/var/vcap/packages/odb-service-adapter/jdk

CLASS_FILES_PATH=/var/vcap/packages/odb-service-adapter/class_files
SRC_FILES=/var/vcap/packages/odb-service-adapter/src
$JAVA_HOME/bin/java -cp $CLASS_FILES_PATH/json.jar:$SRC_FILES example.OdbSpacebearsAdapter "$@"
