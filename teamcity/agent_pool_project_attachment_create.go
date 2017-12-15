package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) CreateAgentPoolProjectAttachment(pool int, apa *types.AgentPoolAttachment) error {
	path := fmt.Sprintf("/httpAuth/app/rest/agentPools/id:%d/projects", pool)
	var poolReturn *types.Project

	err := c.doRetryRequest("POST", path, apa, &poolReturn)
	if err != nil {
		return err
	}

	if poolReturn == nil {
		return errors.New("Project Pool Attachement not created")
	}

	return nil
}