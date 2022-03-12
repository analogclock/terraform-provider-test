---
subcategory: "System Log"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_log_interfacestats"
description: |-
  Interface statistics settings.
---

# fortianalyzer_system_log_interfacestats
Interface statistics settings.

## Argument Reference


The following arguments are supported:


* `billing_report` - Disable/Enable billing report feature. disable - Disable billing report. enable - Enable billing report. Valid values: `disable`, `enable`.

* `retention_days` - Number of days for interface data storage.
* `sampling_interval` - Interval of receiving interface data from FortiGates in seconds.
* `status` - Disable/Enable interface statistics feature. disable - Disable querying FortiGate interface stats. enable - Enable querying FortiGate interface stats. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LogInterfaceStats can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_log_interfacestats.labelname SystemLogInterfaceStats
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

