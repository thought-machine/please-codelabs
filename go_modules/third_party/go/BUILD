# Dependencies in Please vs. go build
go_module(
    name = "protobuf_go",
    # By default, we only install the top level package i.e. golang.org/x/sys. To
    # compile everything, use this wildcard.
    install = ["..."],
    module = "google.golang.org/protobuf",
    version = "v1.25.0",
    # We must tell Please that :protobuf_go depends on :cmp so we can link to it.
    deps = [":cmp"],
)

go_module(
    name = "cmp",
    install = ["cmp/..."],
    module = "github.com/google/go-cmp",
    version = "v0.5.5",
)

# go_mod_download()
go_mod_download(
    name = "protobuf_download",
    module = "github.com/golang/protobuf",
    version = "v1.4.3",
)

go_module(
    name = "protobuf",
    # Depend on our download rule instead of providing a version
    download = ":protobuf_download",
    install = ["..."],
    module = "github.com/golang/protobuf",
    # Let's skip compiling this package which as we're compiling this separately.
    strip = ["protoc-gen-go"],
    deps = [":protobuf_go"],
)

go_module(
    name = "protoc-gen-go",
    # Mark this as binary so Please knows it can be executed
    binary = True,
    # Depend on our download rule instead of providing a version
    download = ":protobuf_download",
    install = ["protoc-gen-go"],
    module = "github.com/golang/protobuf",
    deps = [":protobuf_go"],
)

# Resolving cyclic dependencies
go_module(
    name = "xsys",
    module = "golang.org/x/sys",
    install = ["..."],
    version = "v0.0.0-20210415045647-66c3f260301c",
)

go_module(
    name = "net",
    install = ["..."],
    module = "golang.org/x/net",
    version = "136a25c244d3019482a795d728110278d6ba09a4",
    deps = [
        ":crypto",
        ":text",
    ],
)

go_module(
    name = "text",
    install = [
        "secure/...",
        "unicode/...",
        "transform",
        "encoding/...",
    ],
    module = "golang.org/x/text",
    version = "v0.3.5",
)

go_module(
    name = "crypto",
    install = [
        "ssh/terminal",
        "cast5",
    ],
    module = "golang.org/x/crypto",
    version = "7b85b097bf7527677d54d3220065e966a0e3b613",
)

go_module(
    name = "grpc",
    module = "google.golang.org/grpc",
    version = "v1.34.0",
    # Installing just a subset of stuff to reduce the complexity of this example. You may want to just install "...",
    # and add the rest of the dependencies.
    install = [
        ".",
        "codes",
        "status",
    ],
    deps = [
        ":genproto_rpc",
        ":cmp",
        ":protobuf",
        ":xsys",
        ":net",
        ":protobuf_go",
    ],
)

go_mod_download(
    name = "genproto_download",
    module = "google.golang.org/genproto",
    version = "v0.0.0-20210315173758-2651cd453018",
)

go_module(
    name = "genproto_rpc",
    download = ":genproto_download",
    install = [
        "googleapis/rpc/...",
    ],
    module = "google.golang.org/genproto",
    deps = [
        ":protobuf",
    ],
)

go_module(
    name = "genproto_api",
    download = ":genproto_download",
    install = [
        "googleapis/api/...",
    ],
    module = "google.golang.org/genproto",
    deps = [
        ":grpc",
        ":protobuf",
    ],
)