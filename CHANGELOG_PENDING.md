SDK Bug Fixes
---
* `private/model/api`: Fix doc error list #1067. 
  * Fixes: spacing when generating error list in documentation
* `private/protocol/api`: Fixing json to throw an error if a float number is (+/-)Inf and NaN #1068. 
  * Fixes: `BuildJSON` will not return improper float values of NaN or Inf.

SDK Enhancement
---
* `private/model`: Add service response error code generation (#1061)
  * Adds generation of Error codes for errors returned by a service. Only error codes that are modeled are generated.
