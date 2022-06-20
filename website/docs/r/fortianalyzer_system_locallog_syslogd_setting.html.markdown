---
subcategory: "System LocalLog"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_locallog_syslogd_setting"
description: |-
  Settings for remote syslog server.
---

# fortianalyzer_system_locallog_syslogd_setting
Settings for remote syslog server.

## Example Usage

```hcl
resource "fortianalyzer_system_locallog_syslogd_setting" "trname" {
  csv      = "disable"
  facility = "local7"
  severity = "notification"
  status   = "disable"
}
```

## Argument Reference


The following arguments are supported:


* `cert` - Select local certificate used for secure connection.
* `csv` - CSV format. disable - Disable CSV format. enable - Enable CSV format. Valid values: `disable`, `enable`.

* `facility` - Remote syslog facility. kernel - Kernel messages. user - Random user-level messages. ntp - NTP daemon. audit - Log audit. alert - Log alert. clock - Clock daemon. mail - Mail system. daemon - System daemons. auth - Security/authorization messages. syslog - Messages generated internally by syslog daemon. lpr - Line printer subsystem. news - Network news subsystem. uucp - Network news subsystem. cron - Clock daemon. authpriv - Security/authorization messages (private). ftp - FTP daemon. local0 - Reserved for local use. local1 - Reserved for local use. local2 - Reserved for local use. local3 - Reserved for local use. local4 - Reserved for local use. local5 - Reserved for local use. local6 - Reserved for local use. local7 - Reserved for local use. Valid values: `kernel`, `user`, `ntp`, `audit`, `alert`, `clock`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.

* `reliable` - Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: `disable`, `enable`.

* `secure_connection` - Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: `disable`, `enable`.

* `severity` - Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `status` - Remote syslog log. disable - Do not log to remote syslog server. enable - Log to remote syslog server. Valid values: `disable`, `enable`.

* `syslog_name` - Remote syslog server name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LocallogSyslogdSetting can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_locallog_syslogd_setting.labelname SystemLocallogSyslogdSetting
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

