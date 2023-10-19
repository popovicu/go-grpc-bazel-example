load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "foo_proto",
    srcs = ["foo.proto"],
)

go_proto_library(
    name = "foo_go_proto",
    proto = ":foo_proto",
    importpath = "popovicu/foo",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    visibility = [
        "//visibility:public",
    ],
)