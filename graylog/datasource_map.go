package graylog

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hen2001/terraform-provider-graylog/graylog/datasource/dashboard"
	"github.com/hen2001/terraform-provider-graylog/graylog/datasource/sidecar"
	"github.com/hen2001/terraform-provider-graylog/graylog/datasource/stream"
	"github.com/hen2001/terraform-provider-graylog/graylog/datasource/system/indices/indexset"
)

var dataSourcesMap = map[string]*schema.Resource{
	"graylog_dashboard": dashboard.DataSource(),
	"graylog_index_set": indexset.DataSource(),
	"graylog_sidecar":   sidecar.DataSource(),
	"graylog_stream":    stream.DataSource(),
}
