---
subcategory: "System SQL"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_sql_customindex"
description: |-
  List of SQL index fields.
---

# fortianalyzer_system_sql_customindex
List of SQL index fields.

## Argument Reference


The following arguments are supported:


* `case_sensitive` - Disable/Enable case sensitive index. disable - Build a case insensitive index. enable - Build a case sensitive index. Valid values: `disable`, `enable`.

* `device_type` - Device type. FortiGate - Device type to FortiGate. FortiMail - Device type to FortiMail. FortiWeb - Device type to FortiWeb. Valid values: `FortiGate`, `FortiMail`, `FortiWeb`.

* `fosid` - Add or Edit log index fields.
* `index_field` - Log field name to be indexed.
* `log_type` - Log type. app-ctrl -  attack -  content -  dlp -  emailfilter -  event -  generic -  history -  traffic -  virus -  voip -  webfilter -  netscan -  fct-event -  fct-traffic -  fct-netscan -  waf -  gtp -  dns -  ssh -  ssl -  file-filter -  asset -  protocol -  siem -  Valid values: `app-ctrl`, `attack`, `content`, `dlp`, `emailfilter`, `event`, `generic`, `history`, `traffic`, `virus`, `voip`, `webfilter`, `netscan`, `fct-event`, `fct-traffic`, `fct-netscan`, `waf`, `gtp`, `dns`, `ssh`, `ssl`, `file-filter`, `asset`, `protocol`, `siem`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SqlCustomIndex can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_sql_customindex.labelname {{fosid}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

