{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "AllowsTraffic if true, the flow will be accepted. Otherwise other actions like \"logs\" can still be done, but the traffic will be rejected.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "allowsTraffic",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "boolean",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "DestinationPorts contains the list of allowed ports and ranges.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "destinationPorts",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": "ports_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "EncryptionEnabled defines if the flow has to be encrypted.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "encryptionEnabled",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "boolean",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "LogsEnabled defines if the flow has to be logged.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "logsEnabled",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "boolean",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Object of the policy.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "object",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": "policies_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Subject of the policy.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "subject",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": "policies_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "netpol",
            "netpols"
        ],
        "create": null,
        "delete": true,
        "description": "Allows to define networking policies to allow or prevent processing units identitied by their tags to talk to other processing units or external services (also identified by their tags).",
        "entity_name": "NetworkAccessPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named",
            "@propagated",
            "@schedulable"
        ],
        "get": true,
        "package": "Squall API",
        "resource_name": "networkaccesspolicies",
        "rest_name": "networkaccesspolicy",
        "root": null,
        "update": true
    }
}