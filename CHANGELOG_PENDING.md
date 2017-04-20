### SDK Features

### SDK Enhancements
* `aws/endpoints`: Add utilities improving endpoints lookup (#1218)
  * Adds several utilities to the endpoints packages to make looking up partitions, regions, and services easier.
  * Fixes #994

### SDK Bugs
* `private/protocol/xml/xmlutil`: Fix unmarshaling dropping errors (#1219)
  * The XML unmarshaler would drop any serialization or body read error that occurred on the floor effectively hiding any errors that would occur.
  * Fixes #1205
