#!/bin/bash
# Copyright 2019 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

echo "started build..."

apt-get -y update && apt-get -y upgrade
apt-get install -y git-core
git clone --branch agent-build "https://github.com/GoogleCloudPlatform/osconfig.git"

cd osconfig

<<<<<<< HEAD:packaging/build_packages_deb.sh
apt-get install -y git-core 
git clone "https://github.com/${BASE_REPO}/osconfig.git"
cd osconfig
packaging/setup_deb.sh 
<<<<<<< HEAD:packaging/build_packages_deb.sh
gsutil cp /tmp/debpackage/google-osconfig-agent*.deb "${GCS_PATH}/"
=======
gsutil cp /tmp/debpackage/google-osconfig-agent*.deb "${GCS_PATH}/" 
=======
source ./agent-packaging/packaging-scripts/setup_deb.sh
gsutil cp /tmp/debpackage/google-osconfig-agent*.deb "gs://osconfig-agent-package/"
>>>>>>> Add osconfig agent packaging scripts and docker file:agent-packaging/packaging-scripts/build_packages_deb.sh
>>>>>>> Add osconfig agent packaging scripts and docker file:agent-packaging/packaging-scripts/build_packages_deb.sh

echo 'Package build success'