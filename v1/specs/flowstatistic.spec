# Model
model:
  rest_name: flowstatistic
  resource_name: flowstatistics
  entity_name: FlowStatistic
  package: jenova
  description: |-
    Returns network access statistics on a particular processing unit or group of
    processing units based on their tags.
  aliases:
  - flowstat
  - flowstats
  extends:
  - '@identifiable-nopk-nostored'

# Attributes
attributes:
  v1:
  - name: dataPoints
    description: |-
      DataPoints is a list of time/value pairs that represent the flow events over
      time.
    type: external
    exposed: true
    subtype: datapoints_list
    read_only: true
    autogenerated: true

  - name: destinationIDs
    description: DestinationIDs is the IDs of the destination.
    type: external
    exposed: true
    subtype: flowstatistic_origin_list
    read_only: true
    autogenerated: true
    format: free

  - name: destinationTags
    description: DestinationTags contains the tags used to identify destination.
    type: external
    exposed: true
    subtype: selectors_list
    read_only: true
    autogenerated: true

  - name: metric
    description: Metric is the kind of metric the statistic represents.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Flows
    - Ports
    autogenerated: true
    default_value: Flows

  - name: mode
    description: Mode defines if the metric is for accepted or rejected flows.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Accepted
    - Any
    - Rejected
    autogenerated: true
    default_value: Accepted

  - name: sourceIDs
    description: SourceIDs is the sources of the stats.
    type: external
    exposed: true
    subtype: flowstatistic_origin_list
    read_only: true
    autogenerated: true
    format: free

  - name: sourceTags
    description: SourceTags contains the tags used to identify the source.
    type: external
    exposed: true
    subtype: selectors_list
    read_only: true
    autogenerated: true

  - name: type
    description: Type is the type of representation.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Repartition
    - Serie
    autogenerated: true
    default_value: Serie

  - name: userIdentifier
    description: |-
      UserIdentifier can be set by the user as a query parameter. It will be returned
      in the FlowStatistic object.
    type: string
    exposed: true
    autogenerated: true
    format: free
    orderable: true
