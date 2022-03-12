---
subcategory: "System SNMP"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_snmp_community"
description: |-
  SNMP community configuration.
---

# fortianalyzer_system_snmp_community
SNMP community configuration.

## Argument Reference


The following arguments are supported:


* `events` - SNMP trap events. disk_low - Disk usage too high. intf_ip_chg - Interface IP address changed. sys_reboot - System reboot. cpu_high - CPU usage too high. mem_low - Available memory is low. log-alert - Log base alert message. log-rate - High incoming log rate detected. log-data-rate - High incoming log data rate detected. lic-gbday - High licensed log GB/day detected. lic-dev-quota - High licensed device quota detected. cpu-high-exclude-nice - CPU usage exclude NICE threshold. Valid values: `disk_low`, `intf_ip_chg`, `sys_reboot`, `cpu_high`, `mem_low`, `log-alert`, `log-rate`, `log-data-rate`, `lic-gbday`, `lic-dev-quota`, `cpu-high-exclude-nice`.

* `hosts` - Hosts. The structure of `hosts` block is documented below.
* `hosts6` - Hosts6. The structure of `hosts6` block is documented below.
* `fosid` - Community ID.
* `name` - Community name.
* `query_v1_port` - SNMP v1 query port.
* `query_v1_status` - Enable/disable SNMP v1 query. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `query_v2c_port` - SNMP v2c query port.
* `query_v2c_status` - Enable/disable SNMP v2c query. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `status` - Enable/disable community. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `trap_v1_rport` - SNMP v1 trap remote port.
* `trap_v1_status` - Enable/disable SNMP v1 trap. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `trap_v2c_rport` - SNMP v2c trap remote port.
* `trap_v2c_status` - Enable/disable SNMP v2c trap. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `hosts` block supports:

* `id` - Host entry ID.
* `interface` - Allow interface name.
* `ip` - Allow host IP address.

The `hosts6` block supports:

* `id` - Host entry ID.
* `interface` - Allow interface name.
* `ip` - Allow host IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SnmpCommunity can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_snmp_community.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

