
# Overview

Automation of GCP resources.

## Why not Terraform?

## Features

1. Create a resource
2. Delete a resource
3. Dry run - simply generate a _compliant_ and _hydrated_ YAML file describing the resources
4. Well known defaults to reduce boiler plate code
5. Embedded resources is not supported.  Meaning each resource should be declared separately.  A specific resource can reference an externally created resource via ID.  E.g. project can reference a folder id that contains the project, but that folder should exist (declared separately)

# Supported Resources

- [ ] Folders
- [ ] Projects

# Dependencies

1. Requires `buf` utility.  Further reading on [buf.build][1]

---
[1]: https://buf.build/docs/