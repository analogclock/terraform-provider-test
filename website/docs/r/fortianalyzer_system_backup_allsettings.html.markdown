---
subcategory: "System Others"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_backup_allsettings"
description: |-
  Scheduled backup settings.
---

# fortianalyzer_system_backup_allsettings
Scheduled backup settings.

## Argument Reference


The following arguments are supported:


* `cert` - SSH certificate for authentication.
* `crptpasswd` - Optional password to protect backup content.
* `directory` - Directory in which file will be stored on backup server.
* `passwd` - Backup server login user password.
* `protocol` - Protocol used to backup. sftp - SFTP. ftp - FTP. scp - SCP. Valid values: `sftp`, `ftp`, `scp`.

* `server` - Backup server name/IP and port.
* `status` - Enable/disable schedule backup. disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `time` - Time to backup.
* `user` - Backup server login user.
* `week_days` - Week days to backup. monday - Monday. tuesday - Tuesday. wednesday - Wednesday. thursday - Thursday. friday - Friday. saturday - Saturday. sunday - Sunday. Valid values: `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System BackupAllSettings can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_backup_allsettings.labelname SystemBackupAllSettings
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

