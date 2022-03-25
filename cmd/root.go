package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

var ContainerName string
var NameSpace string
var UserName string

var rootCmd = &cobra.Command{
	Use:   "go",
	Short: "kubectl exec in pod with username",
	Long: `kubectl exec in pod with username. For example:
kubectl go pod_name`,
	Run: runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	fmt.Printf("execute %s args:%v \n", cmd.Name(), args)

	kubectl, _ := exec.LookPath("kubectl")
	var cmdArgs []string
	if UserName == "" {
		UserName = "dev"
	}
	cmdArgs = []string{"kubectl", "exec", "-it", args[0]}
	if ContainerName != "" {
		cmdArgs = append(cmdArgs, "-c", ContainerName)
	}
	if NameSpace != "" {
		cmdArgs = append(cmdArgs, "-n", NameSpace)
	}
	cmdArgs = append(cmdArgs, "--", "su", "-", UserName)
	env := os.Environ()
	execErr := syscall.Exec(kubectl, cmdArgs, env)
	if execErr != nil {
		panic(execErr)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&ContainerName, "containerName", "c", "", "containerName")
	rootCmd.Flags().StringVarP(&NameSpace, "namespace", "n", "", "namespace")
	rootCmd.Flags().StringVarP(&UserName, "username", "u", "", "username, this user must exist in image, default: dev")
}
