package test

import (
	"context"

	"go.keploy.io/server/pkg/hooks"
	"go.keploy.io/server/pkg/models"
	"go.keploy.io/server/pkg/platform"
	"go.keploy.io/server/pkg/platform/yaml"
)

type Tester interface {
	// Test(path, testReportPath string, appCmd string, testsets []string, appContainer, networkName string, Delay uint64, passThorughPorts []uint, apiTimeout uint64) bool
	Test(path, testReportPath string, appCmd string, options TestOptions) bool
	RunTestSet(testSet, path, testReportPath, appCmd, appContainer, appNetwork string, delay uint64, pid uint32, ys platform.TestCaseDB, loadedHook *hooks.Hook, testReportfs yaml.TestReportFS, testRunChan chan string, apiTimeout uint64, ctx context.Context, noiseConfig map[string]interface{}, serveTest bool) models.TestRunStatus
}
