/*
Package specs is a simple wrapper around Go's testing package. It provides some
helper functions which makes it convenient to perform common test operations,
while avoiding typing many of the repetative if-checks.

Example usage:

    import (
      "testing"

      "github.com/knakk/specs"
    )

    func add(a, b int) int {
      sum := a + b
      return sum
    }

    func TestAdd(t *testing.T) {
      s := specs.New(t)

      s.Expect(add(1, 2), 3)
      s.ExpectNot(add(1, 2), 4)
      s.ExpectNotNil(add(0, 0))

      tests := []specs.Spec{
        {add(1, 2), 3},
        {add(2, 2), 4},
        {add(4, -5), -1},
        {add(1000, 1), 1001},
      }

      s.ExpectAll(tests)
    }

All functions expect ExpectAll optionally takes a string as a last argument.
This will be printed instead of the generic message "expected x to be y" if
the test fails:

    s.Expect(3, 4, "3 cannot be 4!")

A common pattern is cheking if a function returned an error before comparing the
actual value against the expected. The function ErrExpect does excactly this:

    i, err := strconv.Atoi("2")
    s.ErrExpect(err, i, 2)

This is the same as:

    i, err := strconv.Atoi("2")
    if err != nil {
      t.Fatal(err)
    }
    if i != 2 {
      t.ErrorF("expected %v to be 2", i)
    }
*/
package specs
