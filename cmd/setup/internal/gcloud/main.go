package gcloud

import (
	"fmt"
	"os/exec"
	"strings"
)

// Exists checks whether `gcloud` can be found in the path.
func Exists() bool {
	_, err := exec.LookPath("gcloud")
	if err != nil {
		return false
	}
	return true
}

// CurrentProject returns the default gcloud project, which is set as the
// `core/project` value in the gcloud config.
//
// The underlying command is:
//   gcloud config get-value core/project
func CurrentProject() (string, error) {
	output, err := simpleStringExec("gcloud", "config", "get-value", "core/project")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(output), nil
}

// ListEnabledServices returns a list of services that have been enabled
// for the specified project. It's equivalent to calling:
//
//    gcloud --project=[PROJECT] services list --enabled --format=value(config.name)
//
// TODO(fred): The underlying endpoint is paginated and I'm not sure how robust
//     this is for projects with long lists of enabled services. It would depend on
//     whether the library automatically paginates under the hood. It's worth taking
//     a look at default limit behavior.
//
// TODO(fred): Can possibly be done safer without gcloud! Although I didn't find
//     it at the time I implemented this, there's a Terraform provider for managing
//     services so there must be an API.
func ListEnabledServices(project string) (map[string]bool, error) {
	services := make(map[string]bool, 20)

	// gcloud services list --enabled --format="value(config.name)"
	output, err := simpleStringExec("gcloud", fmt.Sprintf("--project=%s", project), "services", "list", "--enabled", "--format=value(config.name)")
	if err != nil {
		return services, err
	}

	for _, service := range strings.Split(strings.TrimSpace(output), "\n") {
		services[service] = true
	}

	return services, nil
}

// EnableServices enables a list of provided services. It's equivalent to calling:
//     gcloud --project=[PROJECT] services enable [SERVICES...]
//
// TODO(fred): Can possibly be done safer without gcloud! Although I didn't find
//     it at the time I implemented this, there's a Terraform provider for managing
//     services so there must be an API.
func EnableServices(project string, services []string) error {
	arg := append([]string{fmt.Sprintf("--project=%s", project), "services", "enable"}, services...)
	_, err := simpleStringExec("gcloud", arg...)
	if err != nil {
		return err
	}
	return nil
}

// AcquireDefaultCredentials btains user access credentials via a web flow and puts them in the
// well-known location for Application Default Credentials (ADC). Libraries like Terraform
// will go there to look for credentials.
//
// This is equivalent to:
//     gcloud --project=[PROJECT] auth application-default login
//
// The command name is a bit of a handful, but there's a helpful explanation of
// what it does in here:
//
//     gcloud auth application-default login --help
//
func AcquireDefaultCredentials(project string) error {
	_, err := simpleStringExec("gcloud", fmt.Sprintf("--project=%s", project), "auth", "application-default", "login")
	if err != nil {
		return err
	}
	return nil
}
