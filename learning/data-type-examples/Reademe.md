### Diffrense b/w the printf and println
*** Key difference: Printf() doesn't add a newline automatically—you must add \n if needed. Println() adds a newline automatically. ***

## Airthematic-Operations:
*** int + float64 ❌ not allowed in Go ***

### Why Go gives error?

### Go is strictly typed.

👉 It does NOT do automatic conversion like Python/JavaScript

int ≠ float64

So Go says:

"You must make both types SAME before adding"

Always convert smaller type → bigger type

***
 Common format verbs in Go's fmt package:
%d: Integer (decimal)
%f: Floating-point number
%s: String
%t: Boolean (true/false)
%v: Default format for any value (useful for debugging)
***
