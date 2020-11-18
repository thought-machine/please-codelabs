go_library(
    name = "greetings",
    srcs = ["greetings.go"],
    visibility = ["//src/..."],
)

go_test(
    name = "greetings_test",
    srcs = ["greetings_test.go"],
    deps = [
        ":greetings",
        "//third_party/go:assert",
    ],
    external = True,
)