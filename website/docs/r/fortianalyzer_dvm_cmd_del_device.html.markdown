---
subcategory: "Device Manager"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_dvm_cmd_del_device"
description: |-
  Delete a device.
---

# fortianalyzer_dvm_cmd_del_device
Delete a device.

## Argument Reference


The following arguments are supported:


* `fazadom` - Name or ID of the ADOM where the command is to be executed on.
* `device` - Name or ID of the target device.
* `flags` - create_task - Create a new task in task manager database. nonblocking - The API will return immediately in for non-blocking call. This flag will be set automatically when the adding, importing, updating, and deleting a list of devices. Valid values: `none`, `create_task`, `nonblocking`.

* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dvm CmdDelDevice can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_dvm_cmd_del_device.labelname DvmCmdDelDevice
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

## Others

~> **Warning:** This resource is an `execution` resource, which means it has no state consistency check function. After each execution, if you want to re-execute it, please use terraform taint or assign a different new value to `force_recreate`, then apply it again.
