/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package install

import (
	c "github.com/openziti/ziti/ziti/constants"
)

// UpgradeZitiProxCOptions the options for the upgrade ziti-prox-c command
type UpgradeZitiProxCOptions struct {
	InstallOptions

	Version string
}

// Run implements the command
func (o *UpgradeZitiProxCOptions) Run() error {
	newVersion, err := o.getLatestGitHubReleaseVersion(c.ZITI_SDK_C_GITHUB)
	if err != nil {
		return err
	}

	newVersionStr := newVersion.String()

	if o.Version != "" {
		newVersionStr = o.Version
	}

	o.deleteInstalledBinary(c.ZITI_PROX_C)

	return o.FindVersionAndInstallGitHubRelease(true, c.ZITI_PROX_C, c.ZITI_SDK_C_GITHUB, newVersionStr)
}
