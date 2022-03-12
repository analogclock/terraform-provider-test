---
subcategory: "System Saml"
layout: "fortianalyzer"
page_title: "FortiAnalyzer: fortianalyzer_system_saml_fabricidp"
description: |-
  Authorized identity providers.
---

# fortianalyzer_system_saml_fabricidp
Authorized identity providers.

## Argument Reference


The following arguments are supported:


* `dev_id` - IDP Device ID.
* `idp_cert` - IDP Certificate name.
* `idp_entity_id` - IDP entity ID.
* `idp_single_logout_url` - IDP single logout url.
* `idp_single_sign_on_url` - IDP single sign-on URL.
* `idp_status` - Enable/disable SAML authentication (default = disable). disable - Disable SAML authentication. enable - Enabld SAML authentication. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{dev_id}}.

## Import

System SamlFabricIdp can be imported using any of these accepted formats:
```

$ export "FORTIANALYZER_IMPORT_TABLE"="true"
$ terraform import fortianalyzer_system_saml_fabricidp.labelname {{dev_id}}
$ unset "FORTIANALYZER_IMPORT_TABLE"
```

