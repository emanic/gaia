{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "EnforcerID is the ID of the enforcer associated with the processing unit",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "enforcerID",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "LastSyncTime is the time when the policy was last resolved",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "lastSyncTime",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "NativeContextID is the Docker UUID or service PID",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "nativeContextID",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Initialized",
                "Paused",
                "Running",
                "Stopped",
                "Terminated"
            ],
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": "Initialized",
            "deprecated": false,
            "description": "OperationalStatus of the processing unit",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "operationalStatus",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Docker",
                "LinuxService",
                "RKT"
            ],
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Type of the container ecosystem",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "type",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "fileaccess",
            "update": false
        },
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "renderedpolicy",
            "update": false
        },
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "vulnerability",
            "update": false
        }
    ],
    "model": {
        "aliases": [
            "pu",
            "pus"
        ],
        "create": false,
        "delete": true,
        "description": "A Processing Unit reprents anything that can compute. It can be a Docker container, or a simple Unix process. They are created, updated and deleted by the system as they come and go. You can only modify its tags.  Processing Units use Network Access Policies to define which other Processing Units or External Services they can communicate with and File Access Policies to define what File Paths they can use.",
        "entity_name": "ProcessingUnit",
        "extends": [
            "@archivable",
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "get": true,
        "package": "Squall API",
        "resource_name": "processingunits",
        "rest_name": "processingunit",
        "root": null,
        "update": true
    }
}