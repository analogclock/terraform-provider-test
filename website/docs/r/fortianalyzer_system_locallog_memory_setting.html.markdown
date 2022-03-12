---
subcategory: "System LocalLog"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_memory_setting"
description: |-
  Settings for memory buffer.
---

# fortianalyzer_system_locallog_memory_setting
Settings for memory buffer.

## Argument Reference


The following arguments are supported:


* `diskfull` - Action upon disk full. overwrite - Overwrite oldest log when disk is full. nolog - Stop logging when disk is full. Valid values: `overwrite`, `nolog`.

* `severity` - Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Enable/disable memory buffer log. disable - Do not log to memory buffer. enable - Log to memory buffer. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogMemorySetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_memory_setting.labelname SystemLocallogMemorySetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

