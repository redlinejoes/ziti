/*
	Copyright NetFoundry, Inc.

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

package edge_controller

import (
	"github.com/openziti/ziti/ziti/cmd/ziti/cmd/common"
	cmdutil "github.com/openziti/ziti/ziti/cmd/ziti/cmd/factory"
	"github.com/spf13/cobra"
	"io"
)

// newCreateAuthenticatorCmd creates the 'edge controller create authenticator' command
func newCreateAuthenticatorCmd(f cmdutil.Factory, out io.Writer, errOut io.Writer) *cobra.Command {
	options := edgeOptions{
		CommonOptions:      common.CommonOptions{Factory: f, Out: out, Err: errOut},
		OutputJSONResponse: false,
	}

	cmd := &cobra.Command{
		Use:   "authenticator",
		Short: "creates an authenticator for an identity managed by the Ziti Edge Controller",
		Long:  "creates an authenticator for an identity managed by the Ziti Edge Controller",
	}

	cmd.AddCommand(newCreateAuthenticatorUpdb("updb", options))

	return cmd
}
