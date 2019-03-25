# `gotestlisttags`

To reproduce:

```bash
$ go test --tags="not_ignored" --list="Test" .
```

Expected output:

```
➜  gotestlisttags git:(master) ✗ go test --tags="not_ignored" --list="Test" .
TestShouldNotBeIgnored
ok      gotestlisttags  0.006s
```

Actual ouput:

```
➜  gotestlisttags git:(master) ✗ go test --tags="not_ignored" --list="Test" .
TestShouldBeIgnored
TestShouldNotBeIgnored
ok      gotestlisttags  0.006s
```