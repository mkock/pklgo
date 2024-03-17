# Apple Pkl Go

This is an example application using Apple's new [Pickle (Pkl) framework](https://pkl-lang.org).

A `config` package is automatically generated from the `config/config.pkl` file
in the same directory. The command is:

```bash
pkl-gen-go config/config.pkl --base-path github.com/mkock/pklgo
```

In order to load, validate and parse the configuration, the package `github.com/apple/pkl-go`
is required.

The configuration file itself is fairly straightforward, but it does require a couple of annotations
and a funky import statement:

```golang
import "package://pkg.pkl-lang.org/pkl-go/pkl.golang@0.6.0#/go.pkl"
```

The class definitions are important to achieve the correct types in the generated code, otherwise they
will all have type `any`.
