package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/argoproj/argo-cd/v2/cmd/util"

	appcontroller "github.com/argoproj/argo-cd/v2/cmd/argocd-application-controller/commands"
	applicationset "github.com/argoproj/argo-cd/v2/cmd/argocd-applicationset-controller/commands"
	cmpserver "github.com/argoproj/argo-cd/v2/cmd/argocd-cmp-server/commands"
	dex "github.com/argoproj/argo-cd/v2/cmd/argocd-dex/commands"
	gitaskpass "github.com/argoproj/argo-cd/v2/cmd/argocd-git-ask-pass/commands"
	k8sauth "github.com/argoproj/argo-cd/v2/cmd/argocd-k8s-auth/commands"
	notification "github.com/argoproj/argo-cd/v2/cmd/argocd-notification/commands"
	reposerver "github.com/argoproj/argo-cd/v2/cmd/argocd-repo-server/commands"
	apiserver "github.com/argoproj/argo-cd/v2/cmd/argocd-server/commands"
	cli "github.com/argoproj/argo-cd/v2/cmd/argocd/commands"
	"github.com/spf13/cobra"
)

const (
	binaryNameEnv = "ARGOCD_BINARY_NAME"
)

func main() {

	// Declare a variable of type *cobra.Command
	var command *cobra.Command

	// Returns the first argument passed with comand
	//repo-server: [ "$BIN_MODE" = 'true' ] && COMMAND=./dist/argocd || COMMAND='go run ./cmd/main.go' && sh -c "GOCOVERDIR=${ARGOCD_COVERAGE_DIR:-/tmp/coverage/repo-server} FORCE_LOG_COLORS=1 ARGOCD_FAKE_IN_CLUSTER=true ARGOCD_GNUPGHOME=${ARGOCD_GNUPGHOME:-/tmp/argocd-local/gpg/keys} ARGOCD_PLUGINSOCKFILEPATH=${ARGOCD_PLUGINSOCKFILEPATH:-./test/cmp}  ARGOCD_GPG_DATA_PATH=${ARGOCD_GPG_DATA_PATH:-/tmp/argocd-local/gpg/source} ARGOCD_TLS_DATA_PATH=${ARGOCD_TLS_DATA_PATH:-/tmp/argocd-local/tls} ARGOCD_SSH_DATA_PATH=${ARGOCD_SSH_DATA_PATH:-/tmp/argocd-local/ssh} ARGOCD_BINARY_NAME=argocd-repo-server ARGOCD_GPG_ENABLED=${ARGOCD_GPG_ENABLED:-false}
	//  $COMMAND --loglevel debug --port ${ARGOCD_E2E_REPOSERVER_PORT:-8081} --redis localhost:${ARGOCD_E2E_REDIS_PORT:-6379} --otlp-address=${ARGOCD_OTLP_ADDRESS}"
	// $COMMAND value is passed to the variable binaryName which is argocd
	binaryName := filepath.Base(os.Args[0])
	fmt.Println(binaryName)
	fmt.Println("======= in cmd/main.go =========")
	// binaryNameEnv = ARGOCD_BINARY_NAME = for Repo server the value comes from its Procfile in ARGOCD_BINARY_NAME=argocd-repo-server
	if val := os.Getenv(binaryNameEnv); val != "" {
		binaryName = val
	}
	isCLI := false
	switch binaryName {
	case "argocd", "argocd-linux-amd64", "argocd-darwin-amd64", "argocd-windows-amd64.exe":
		command = cli.NewCommand()
		isCLI = true
	case "argocd-server":
		command = apiserver.NewCommand()
	case "argocd-application-controller":
		command = appcontroller.NewCommand()
	case "argocd-repo-server":
		command = reposerver.NewCommand()
	case "argocd-cmp-server":
		command = cmpserver.NewCommand()
		isCLI = true
	case "argocd-dex":
		command = dex.NewCommand()
	case "argocd-notifications":
		command = notification.NewCommand()
	case "argocd-git-ask-pass":
		command = gitaskpass.NewCommand()
		isCLI = true
	case "argocd-applicationset-controller":
		command = applicationset.NewCommand()
	case "argocd-k8s-auth":
		command = k8sauth.NewCommand()
		isCLI = true
	default:
		command = cli.NewCommand()
		isCLI = true
	}
	util.SetAutoMaxProcs(isCLI)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
