---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_ioc"
description: |-
  IoC settings.
---

# fortianalyzer_system_log_ioc
IoC settings.

## Argument Reference


The following arguments are supported:


* `notification` - Disable/Enable IoC notification. disable - Disable IoC feature. enable - Enable IoC feature. Valid values: `disable`, `enable`.

* `notification_throttle` - Minute value for throttling the rate of IoC notifications.
* `rescan_max_runner` - Max count of cocurrent runner of IoC rescan.
* `rescan_run_at` - When to run IoC rescan.
* `rescan_status` - Disable/Enable IoC rescan. disable - Disable IoC feature. enable - Enable IoC feature. Valid values: `disable`, `enable`.

* `status` - Disable/Enable IoC feature. disable - Disable IoC feature. enable - Enable IoC feature. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogIoc can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_ioc.labelname SystemLogIoc
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

