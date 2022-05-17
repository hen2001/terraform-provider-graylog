package configuration

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hen2001/terraform-provider-graylog/graylog/client"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	if _, _, err := cl.SidecarConfiguration.Update(ctx, d.Id(), data); err != nil {
		return fmt.Errorf("failed to update a configuration %s: %w", d.Id(), err)
	}
	return nil
}
