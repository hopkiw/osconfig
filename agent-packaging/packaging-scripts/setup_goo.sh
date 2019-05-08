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

echo "running common script..."
source ./agent-packaging/packaging-scripts/common.sh

<<<<<<< HEAD:packaging/setup_goo.sh
echo "Building package"
GOPATH=${GOPATH} ${GO} get github.com/google/googet/goopack
${GOPATH}/bin/goopack packaging/googet/google-osconfig-agent.goospec
=======
rm -rf ${GOPATH}/src/github.com/GoogleCloudPlatform/osconfig
sudo cp -r ../../../osconfig/ /usr/share/gocode/src/github.com/GoogleCloudPlatform/

echo "Building package"
sudo su -c "GOPATH=${GOPATH} ${GO} get -d github.com/google/googet/goopack"
$GO run github.com/google/googet/goopack agent-packaging/packaging-scripts/googet/google-osconfig-agent.goospec
>>>>>>> Add osconfig agent packaging scripts and docker file:agent-packaging/packaging-scripts/setup_goo.sh