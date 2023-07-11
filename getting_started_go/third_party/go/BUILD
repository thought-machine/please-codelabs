
go_toolchain(
    name = "toolchain",
    version = "1.20",
)

# We give direct modules a name and install list so we can reference them nicely
go_repo(
    name = "testify",
    module = "github.com/stretchr/testify",
    version="v1.8.2",
    # We add the subset of packages we actually depend on here
    install = [
        "assert",
        "require",
    ]
)

# Indirect modules are referenced internally, so we don't have to name them if we don't want to. They can still be
# referenced by the following build label naming convention: ///third_party/go/github.com_owner_repo//package.
#
# NB: Any slashes in the module name will be replaced by _
go_repo(module="github.com/davecgh/go-spew", version="v1.1.1")
go_repo(module="github.com/pmezard/go-difflib", version="v1.0.0")
go_repo(module="github.com/stretchr/objx", version="v0.5.0")
go_repo(module="gopkg.in/yaml.v3", version="v3.0.1")
go_repo(module="gopkg.in/check.v1", version="v0.0.0-20161208181325-20d25e280405")
