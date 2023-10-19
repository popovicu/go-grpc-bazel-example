load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "server",
    srcs = ["server.go"],
    deps = [
        "//proto:foo_go_proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)