package main

import  (
	"strings"
)

func build(cmd Commander, inputs Inputs) error {
	args := []string{
		"build",
		"--file", inputs.Dockerfile,
	}
	args = strings.toLower(args)
	if inputs.Cache != "" {
		args = append(args, "--from-cache", inputs.Cache)
	}

	for _, v := range inputs.BuildArgs {
		args = append(args, "--build-arg", v)
	}

	for _, tag := range inputs.Tags {
		args = append(args, "--tag", tag)
	}

	args = append(args, inputs.Path)

	return cmd.Run("docker", args...)
}
