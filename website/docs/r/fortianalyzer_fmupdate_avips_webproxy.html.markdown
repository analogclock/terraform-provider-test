---
subcategory: "Fmupdate"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_fmupdate_avips_webproxy"
description: |-
  Configure the web proxy for use with FortiGuard antivirus and IPS updates.
---

# fortianalyzer_fmupdate_avips_webproxy
Configure the web proxy for use with FortiGuard antivirus and IPS updates.

## Example Usage

```hcl
resource "fortianalyzer_fmupdate_avips_webproxy" "trname" {
  mode     = "tunnel"
  password = ["password"]
  port     = 80
  status   = "disable"
  username = "webproxy"
}
```

## Argument Reference


The following arguments are supported:


* `address` - web proxy address.
* `mode` - Web proxy mode proxy - HTTP proxy mode tunnel - HTTP tunnel mode (default) Valid values: `proxy`, `tunnel`.

* `password` - The password for the user name used for authentication.
* `port` - The port number of the web proxy (1 - 65535, default = 80).
* `status` - Enable/disable connections through the web proxy (default = disable). disable - Disable setting. enable - Enable setting. Valid values: `disable`, `enable`.

* `username` - The user name used for authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Fmupdate AvIpsWebProxy can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_fmupdate_avips_webproxy.labelname FmupdateAvIpsWebProxy
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

