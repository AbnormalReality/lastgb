package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	starter "github.com/AbnormalReality/lastgb/lesson3/tests_tdd"
	"github.com/stretchr/testify/assert"
)

// basic unit tests
func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Sergey")
	assert.Equal(t, "Hello Sergey. Welcome!", greeting)
	another_greeting := starter.SayHello("asdf ghjkl")
	assert.Equal(t, "Hello asdf ghjkl. Welcome!", another_greeting)
}

//TDD:
/*
func TestOddOrEven(t *testing.T) {
  assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
}

--- FAIL: TestOddOrEven (0.00s)
    starter_test.go:18:
        	Error Trace:	starter_test.go:18
        	Error:      	Not equal:
        	            	expected: "45 is an odd number"
        	            	actual  : ""

func OddOrEven(num int) string {
  return fmt.Sprintf("%v is an odd number", num)
}

func TestOddOrEven(t *testing.T) {
  assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
  assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
}

--- FAIL: TestOddOrEven (0.00s)
    starter_test.go:19:
        	Error Trace:	starter_test.go:19
        	Error:      	Not equal:
        	            	expected: "42 is an even number"
        	            	actual  : "42 is an odd number"

func OddOrEven(num int) string {
  criteria := math.Mod(float64(num), 2)
  if criteria == 1 {
    return fmt.Sprintf("%v is an odd number", num)
  }
  return fmt.Sprintf("%v is an even number", num)
}


func TestOddOrEven(t *testing.T) {
  assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
  assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
  assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
}

func TestOddOrEven(t *testing.T) {
  assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
  assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
  assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
  assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
}

--- FAIL: TestOddOrEven (0.00s)
    starter_test.go:21:
        	Error Trace:	starter_test.go:21
        	Error:      	Not equal:
        	            	expected: "-45 is an odd number"
        	            	actual  : "-45 is an even number"

func OddOrEven(num int) string {
  criteria := math.Mod(float64(num), 2)
  if criteria == 1 || criteria == -1 {
    return fmt.Sprintf("%v is an odd number", num)
  }
  return fmt.Sprintf("%v is an even number", num)
}


and to final version

=== RUN   TestSayHello
--- PASS: TestSayHello (0.00s)
=== RUN   TestOddOrEven
=== RUN   TestOddOrEven/Check_Non_Negative_Numbers
=== RUN   TestOddOrEven/Check_Negative_Numbers
--- PASS: TestOddOrEven (0.00s)
    --- PASS: TestOddOrEven/Check_Non_Negative_Numbers (0.00s)
    --- PASS: TestOddOrEven/Check_Negative_Numbers (0.00s) */
func TestOddOrEven(t *testing.T) {
	t.Run("Check Non Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("Check Negative Numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
		assert.Equal(t, "-42 is an even number", starter.OddOrEven(-42))
	})
}

func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.Checkhealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))

		assert.Equal(t, nil, err)
	})
}
