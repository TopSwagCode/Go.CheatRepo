# Tests

TLDR:

* File names has to be *_test.go eg. Api_Test.go
* func names has to be Test* eg: TestThatGoIsAwesome
* func has to have (t *testing.T) as parameter
* To run tests type: ```$ go test``` in the terminal

You can use T for failing your test. Skip tests and a lot more. Try playing with it. 2 Quick examples can be seen below. One failing test and a skipped test.

```Go
func TestThat1IsBiggerThan2(t *testing.T) {
	if 1 > 2{
		t.Errorf("Failed test. Testing that 1 is bigger than 2 has failed")
	}
}
```

```Go
func TestThat1IsLassThan2(t *testing.T) {
    t.Skip()
	if 1 < 2{
		t.Errorf("Testing that 1 is less than 2 has failed")
	}
}
```

Furthermore if you want some Setup / Teardown for your tests you can implement a TestMain func for that. A small example can be seen below here.

```Go
func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")

	exitVal := m.Run()

	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
```

When you feel like you have a good graps on doing basic tests. Take a look on the Mocking part, for mocking out dependencies.

Also you can spend some time looking on how to Benchmark your code here: TODO.