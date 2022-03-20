# aisd-tester

To build the app you have to run `go build` command. Then in your c++ project create directory called tests and inside create as many files as you want. However the test file strucutre should be following:

```
input
<your input>
output
<expected output>
```

Now while running the tester you just have to specify the path to the c++ project passed as an argument. Please note that the c++ file should be named main.cpp for know cuz im lazy.


