---
subcategory: "System LocalLog"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_setting"
description: |-
  Settings for locallog logging.
---

# fortianalyzer_system_locallog_setting
Settings for locallog logging.

## Argument Reference


The following arguments are supported:


* `log_interval_dev_no_logging` - Interval in minute for logging the event of no logs received from a device.
* `log_interval_disk_full` - Interval in minute for logging the event of disk full.
* `log_interval_gbday_exceeded` - Interval in minute for logging the event of the GB/Day license exceeded.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_setting.labelname SystemLocallogSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

