package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// Define the structure to map the Terraform block
type Config struct {
	Terraform struct {
		RequiredVersion string `hcl:"required_version,optional"`
	} `hcl:"terraform,block"`
}

// runCommand executes a shell command and returns its output or an error.
func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return strings.TrimSpace(string(output)), err
}

// getTerraformVersion lists the installed Terraform versions using tfenv.
func getTerraformVersion() {
	output, err := runCommand("tfenv", "list")
	if err != nil {
		fmt.Println("Error fetching Terraform versions:", err)
		return
	}
	fmt.Println("Installed Terraform Versions:\n", output)
}

// installTerraformVersion installs the specified Terraform version using tfenv.
func installTerraformVersion(version string) {
	fmt.Println("Installing Terraform version:", version)
	_, err := runCommand("tfenv", "install", version)
	if err != nil {
		fmt.Println("Error installing Terraform version:", err)
		return
	}
	fmt.Println("Terraform version", version, "installed successfully!")
}

// useTerraformVersion sets the specified Terraform version using tfenv.
func useTerraformVersion(version string) {
	fmt.Println("Switching to Terraform version:", version)
	_, err := runCommand("tfenv", "use", version)
	if err != nil {
		fmt.Println("Error switching Terraform version:", err)
		return
	}
	fmt.Println("Successfully switched to Terraform version", version)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> <path_to_terraform_file>")
		os.Exit(1)
	}

	var version string

	if len(os.Args) == 3 {
		filePath := os.Args[2]
		parser := hclparse.NewParser()

		// Parse the Terraform file
		file, diags := parser.ParseHCLFile(filePath)
		if diags.HasErrors() {
			fmt.Fprintf(os.Stderr, "Error parsing Terraform file: %s\n", diags.Error())
			os.Exit(1)
		}

		// Decode into the Config struct
		var config Config
		decodeDiags := gohcl.DecodeBody(file.Body, nil, &config)
		if decodeDiags.HasErrors() {
			fmt.Fprintf(os.Stderr, "Error decoding Terraform file: %s\n", decodeDiags.Error())
			os.Exit(1)
		}

		// Print the extracted Terraform version
		if config.Terraform.RequiredVersion != "" {
			fmt.Println("Terraform Required Version:", config.Terraform.RequiredVersion)
		} else {
			fmt.Println("No required_version found in the Terraform file.")
		}
		version = config.Terraform.RequiredVersion
	}

	command := os.Args[1]

	switch command {
	case "list":
		getTerraformVersion()
	case "install":
		installTerraformVersion(version)
	case "use":
		useTerraformVersion(version)
	default:
		fmt.Println("Invalid command. Use 'list', 'install', or 'use'.")
	}
}
