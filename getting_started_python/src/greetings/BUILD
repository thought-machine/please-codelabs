python_library(
    name = "greetings",
    srcs = ["greetings.py"],
    visibility = ["//src/..."],
    deps = ["//third_party/python:numpy"],
)

python_test(
    name = "greetings_test",
    srcs = ["greetings_test.py"],
    # Here we have used the shorthand `:greetings` label format. This format can be used to refer to a rule in the same
    # package and is shorthand for `//src/greetings:greetings`.
    deps = [":greetings"],
)