---
subcategory: "System AutoDelete"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_autodelete"
description: |-
  Automatic deletion policy for logs, reports, archived, and quarantined files.
---

# fortianalyzer_system_autodelete
Automatic deletion policy for logs, reports, archived, and quarantined files.

## Argument Reference


The following arguments are supported:


* `dlp_files_auto_deletion` - Dlp-Files-Auto-Deletion. The structure of `dlp_files_auto_deletion` block is documented below.
* `log_auto_deletion` - Log-Auto-Deletion. The structure of `log_auto_deletion` block is documented below.
* `quarantine_files_auto_deletion` - Quarantine-Files-Auto-Deletion. The structure of `quarantine_files_auto_deletion` block is documented below.
* `report_auto_deletion` - Report-Auto-Deletion. The structure of `report_auto_deletion` block is documented below.
* `status_fake` - Fake value for the menu to work.

The `dlp_files_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.

The `log_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.

The `quarantine_files_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.

The `report_auto_deletion` block supports:

* `retention` - Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: `days`, `weeks`, `months`.

* `runat` - Automatic deletion run at (0 - 23) o'clock.
* `status` - Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: `disable`, `enable`.

* `value` - Automatic deletion in x days, weeks, or months.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoDelete can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_autodelete.labelname SystemAutoDelete
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

