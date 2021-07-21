/*
 * Copyright (c) 2021 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package main

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-lambda-api-lib/golang/lib/execution"
	"github.com/kurtosis-tech/kurtosis-lambda-starter-pack/golang/lambda"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	successExitCode = 0
	failureExitCode = 1
)

func main() {

	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<
	configurator := lambda.NewExampleLambdaConfigurator()
	// >>>>>>>>>>>>>>>>>>> REPLACE WITH YOUR OWN CONFIGURATOR <<<<<<<<<<<<<<<<<<<<<<<<


	lambdaExecutor := execution.NewLambdaExecutor(configurator)
	if err := lambdaExecutor.Run(); err != nil {
		logrus.Errorf("An error occurred running the lambda executor:")
		fmt.Fprintln(logrus.StandardLogger().Out, err)
		os.Exit(failureExitCode)
	}
	os.Exit(successExitCode)
}
