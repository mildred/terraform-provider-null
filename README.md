Mildred's terraform `null` Provider
===================================

Maintainers
-----------

This provider plugin is not maintained by the Terraform team but is a private effort.

Usage
-----

The `null` provider is a rather-unusual provider that has constructs that
intentionally do nothing. This may sound strange, and indeed these constructs
do not need to be used in most cases, but they can be useful in various
situations to help orchestrate tricky behavior or work around limitations.

The documentation of each feature of this provider, accessible via the
navigation, gives examples of situations where these constructs may prove
useful.

Usage of the `null` provider can make a Terraform configuration harder to
understand. While it can be useful in certain cases, it should be applied with
care and other solutions preferred when available.

See documentation for resources:

- [null_resource](docs/resource.html.markdown)

and for data sources:

- [null_data_source](docs/data_source.html.markdown)

