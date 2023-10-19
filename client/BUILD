load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "client",
    srcs = ["client.go"],
    deps = [
        "//proto:foo_go_proto",
        "@org_golang_google_grpc//:go_default_library",
	"@org_golang_google_grpc//credentials/insecure",
    ],
)