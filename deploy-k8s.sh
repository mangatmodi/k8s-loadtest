#!/bin/bash

# The file depends on gpg, kubectl binaries to be available on machine
function checkvar {
    VAR_VALUE=$(echo "${!1}")
    if [[ -z "$VAR_VALUE" ]]; then
      echo "$1 env variable is not set"
      exit 1
    fi
}

if [[ -z "$1" ]]; then
  echo "Usage: $(basename ${0}) project [apply/delete]"
  exit 1
fi

checkvar "DOCKER_IMAGE"

project_dir=$1
command=$2

script_dir="$( cd "$(dirname "$0")" ; pwd -P )"
curr_dir=$(pwd -P)
project_dir="$( cd ${project_dir} ; pwd -P )"

echo "Deploying docker image ${DOCKER_IMAGE} to k8s"

manifest=${project_dir}/k8.yml
manifest_name=$(echo "${manifest}" | sed "s/.*\///")
echo "${manifest_name}"

# replace environment variables
envsubst < ${manifest} > ${manifest_name}-derived.yaml
echo "Applying manifest definition"
cat ${manifest_name}-derived.yaml

#apply  the config
echo "Using kubeconfig $(kubectl config view -o template --template='{{ index . "current-context" }}')"
kubectl ${command} -f ${manifest_name}-derived.yaml -n alt
rm -f *-derived.yaml
