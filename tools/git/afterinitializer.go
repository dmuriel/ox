package git

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/wawandco/oxpecker/lifecycle/new"
)

type AfterInitializer struct{}

func (ri AfterInitializer) Name() string {
	return "git/repoinitializer"
}

func (ri AfterInitializer) AfterInitialize(ctx context.Context, options new.Options) error {
	_, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("[warning] Git repo was not initialized given git was not present")
		return nil
	}

	err = os.Chdir(options.Folder)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
