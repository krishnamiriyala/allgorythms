workspace(name = "allgorhythms")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_skylib",
    url = "https://build-artifactory.eng.vmware.com/nsbu-files-local/git-archives/mirrors_github_bazel-skylib-2b38b2f8bd4b8603d610cfc651fcbb299498147f.tar.gz",
    sha256 = "09aa9200ff4fcd97af2b53bd7e34778755c0137eeb67de527b755a296f4be7c6",
    strip_prefix = "mirrors_github_bazel-skylib-2b38b2f8bd4b8603d610cfc651fcbb299498147f",
)

load("@bazel_skylib//lib:versions.bzl", "versions")

versions.check(minimum_bazel_version="4.0.0", maximum_bazel_version="5.0.0")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")


load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "rules_python",
    sha256 = "9fcf91dbcc31fde6d1edb15f117246d912c33c36f44cf681976bd886538deba6",
    strip_prefix = "rules_python-0.8.0",
    url = "https://github.com/bazelbuild/rules_python/archive/refs/tags/0.8.0.tar.gz",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "f2dcd210c7095febe54b804bb1cd3a58fe8435a909db2ec04e31542631cf715c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.18")
