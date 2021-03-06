//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package ospackage

import (
	"github.com/GoogleCloudPlatform/compute-image-tools/go/packages"

	osconfigpb "github.com/GoogleCloudPlatform/osconfig/_internal/gapi-cloud-osconfig-go/google.golang.org/genproto/googleapis/cloud/osconfig/v1alpha1"
)

// changes represents the delta between the actual and the desired package installation state.
type changes struct {
	packagesToInstall []string
	packagesToUpgrade []string
	packagesToRemove  []string
}

// getNecessaryChanges compares the current state and the desired state to determine which packages
// need to be installed, upgraded, or removed.
func getNecessaryChanges(installedPkgs []packages.PkgInfo, upgradablePkgs []packages.PkgInfo, packageInstalls []*osconfigpb.Package, packageRemovals []*osconfigpb.Package) changes {
	installedPkgMap := make(map[string]bool)
	for _, pkg := range installedPkgs {
		installedPkgMap[pkg.Name] = true
	}

	upgradeablePkgMap := make(map[string]bool)
	for _, pkg := range upgradablePkgs {
		upgradeablePkgMap[pkg.Name] = true
	}

	var pkgsToInstall []string
	var pkgsToUpgrade []string

	for _, pkg := range packageInstalls {
		_, isInstalled := installedPkgMap[pkg.Name]
		_, isUpgradable := upgradeablePkgMap[pkg.Name]

		if !isInstalled {
			pkgsToInstall = append(pkgsToInstall, pkg.Name)
		} else if isInstalled && isUpgradable {
			pkgsToUpgrade = append(pkgsToUpgrade, pkg.Name)
		}
	}

	var pkgsToRemove []string
	for _, pkg := range packageRemovals {
		_, isInstalled := installedPkgMap[pkg.Name]

		if isInstalled {
			pkgsToRemove = append(pkgsToRemove, pkg.Name)
		}
	}

	return changes{
		packagesToInstall: pkgsToInstall,
		packagesToUpgrade: pkgsToUpgrade,
		packagesToRemove:  pkgsToRemove,
	}
}
