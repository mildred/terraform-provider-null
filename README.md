Mildred's terraform `null` Provider
===================================

Maintainers
-----------

This provider plugin is not maintained by the Terraform team but is a private effort.

Resources
---------

### null_resource

Inputs:

- `triggers`: optional `map` (forces new resource)

Outputs:

- `id`: generated `string`

  Random id, regenerated when the resource is created
  Items that can be specified in the resource and that force resource creation

Data Sources
------------

### null_data_source

Inputs:

- `inputs`: optional `map`

  A list of items that are input to the data source

- `has_computed_default`: optional `string`

Outputs:

- `id`: computed `string` with value `"static"`

- `outputs`: computed `map`

  The list of inputs that is set when the data source is read

- `random`: computed `string`

  A random string that is regenerated on each new read

- `has_computed_default`: computed `string`

