# Resource: graylog_pipeline_rule

* [Example](https://github.com/hen2001/terraform-provider-graylog/blob/master/examples/v0.12/pipeline.tf)
* [Source Code](https://github.com/hen2001/terraform-provider-graylog/blob/master/graylog/resource/system/pipeline/rule/resource.go)

## Argument Reference

* `source` - (Required) The source of the Pipeline Rule. The data type is `string`.
* `description` - (Optional) description of the Pipeline Rule. The data type is `string`.

## Attributes Reference

Nothing.

## Import

`graylog_pipeline_rule` can be imported using the Pipeline Rule id, e.g.

```console
$ terraform import graylog_pipeline_rule.test 5c4acaefc9e77bbbbbbbbbbb
```
