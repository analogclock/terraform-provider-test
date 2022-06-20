---
subcategory: "System Admin"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_admin_tacacs"
description: |-
  TACACS+ server entry configuration.
---

# fortianalyzer_system_admin_tacacs
TACACS+ server entry configuration.

## Example Usage

```hcl
resource "fortianalyzer_system_admin_tacacs" "trname" {
  authen_type   = "auto"
  authorization = "enable"
  key           = ["key"]
  name          = "tacacs_user"
  port          = 49
  server        = "1.1.1.1"
}
```

## Argument Reference


The following arguments are supported:


* `authen_type` - Authentication type. auto - Use PAP, MSCHAP, and CHAP (in that order). ascii - ASCII. pap - PAP. chap - CHAP. mschap - MSCHAP. Valid values: `auto`, `ascii`, `pap`, `chap`, `mschap`.

* `authorization` - Enable/disable TACACS+ authorization. disable - Disable TACACS+ authorization. enable - Enable TACACS+ authorization (service = fortigate). Valid values: `disable`, `enable`.

* `key` - <password_str> key to access server.
* `name` - TACACS+ server entry name.
* `port` - Port number of TACACS+ server.
* `secondary_key` - <password_str> key to access secondary server.
* `secondary_server` - {<name_str|ip_str>} secondary server domain name or IP.
* `server` - {<name_str|ip_str>} server domain name or IP.
* `tertiary_key` - <password_str> key to access tertiary server.
* `tertiary_server` - {<name_str|ip_str>} tertiary server domain name or IP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AdminTacacs can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_admin_tacacs.labelname {{name}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

