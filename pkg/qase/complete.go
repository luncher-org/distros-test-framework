package qase

import (
	"fmt"

	"github.com/rancher/distros-test-framework/shared"
)

func (c Client) completeRun(runID int32) error {
	baseRes, res, err := c.QaseAPI.RunsAPI.CompleteRun(c.Ctx, projectID, runID).Execute()
	if err != nil {
		return fmt.Errorf("failed to complete run: %w, response: %v", err, res)
	}

	shared.LogLevel("debug", "Run completed: %v\n", &baseRes)

	return nil
}
