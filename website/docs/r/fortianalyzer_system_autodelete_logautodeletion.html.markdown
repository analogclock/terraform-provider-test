---
subcategory: "System AutoDelete"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_autodelete_logautodeletion"
description: |-
  Automatic deletion policy for device logs.
---

# fortianalyzer_system_autodelete_logautodeletion
Automatic deletion policy for device logs.

## Argument Reference


The following arguments are supported:


* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoDeleteLogAutoDeletion can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_autodelete_logautodeletion.labelname SystemAutoDeleteLogAutoDeletion
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

