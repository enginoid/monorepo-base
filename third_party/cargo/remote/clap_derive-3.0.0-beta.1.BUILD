"""
cargo-raze crate build file.

DO NOT EDIT! Replaced on runs of cargo-raze
"""
package(default_visibility = [
  # Public for visibility by "@raze__crate__version//" targets.
  #
  # Prefer access through "//third_party/cargo", which limits external
  # visibility to explicit Cargo.toml dependencies.
  "//visibility:public",
])

licenses([
  "restricted", # "MIT OR Apache-2.0"
])

load(
    "@io_bazel_rules_rust//rust:rust.bzl",
    "rust_library",
    "rust_binary",
    "rust_test",
)


# Unsupported target "after_help" with type "example" omitted
# Unsupported target "arg_enum_basic" with type "example" omitted
# Unsupported target "arg_enum_basic" with type "test" omitted
# Unsupported target "arg_enum_case_sensitive" with type "example" omitted
# Unsupported target "arg_enum_case_sensitive" with type "test" omitted
# Unsupported target "argument_naming" with type "test" omitted
# Unsupported target "arguments" with type "test" omitted
# Unsupported target "at_least_two" with type "example" omitted
# Unsupported target "author_version_about" with type "test" omitted
# Unsupported target "basic" with type "example" omitted
# Unsupported target "basic" with type "test" omitted
# Unsupported target "boxed" with type "test" omitted

rust_library(
    name = "clap_derive",
    crate_root = "clap_derive/src/lib.rs",
    crate_type = "proc-macro",
    edition = "2018",
    srcs = glob(["**/*.rs"]),
    deps = [
        "@raze__heck__0_3_1//:heck",
        "@raze__proc_macro_error__0_4_12//:proc_macro_error",
        "@raze__proc_macro2__1_0_10//:proc_macro2",
        "@raze__quote__1_0_3//:quote",
        "@raze__syn__1_0_17//:syn",
    ],
    rustc_flags = [
        "--cap-lints=allow",
    ],
    version = "3.0.0-beta.1",
    crate_features = [
        "default",
    ],
)

# Unsupported target "custom-string-parsers" with type "test" omitted
# Unsupported target "default_value" with type "test" omitted
# Unsupported target "deny-warnings" with type "test" omitted
# Unsupported target "deny_missing_docs" with type "example" omitted
# Unsupported target "doc-comments-help" with type "test" omitted
# Unsupported target "doc_comments" with type "example" omitted
# Unsupported target "enum_in_args" with type "example" omitted
# Unsupported target "enum_tuple" with type "example" omitted
# Unsupported target "env" with type "example" omitted
# Unsupported target "example" with type "example" omitted
# Unsupported target "explicit_name_no_renaming" with type "test" omitted
# Unsupported target "flags" with type "test" omitted
# Unsupported target "flatten" with type "example" omitted
# Unsupported target "flatten" with type "test" omitted
# Unsupported target "from_crate" with type "example" omitted
# Unsupported target "git" with type "example" omitted
# Unsupported target "group" with type "example" omitted
# Unsupported target "issues" with type "test" omitted
# Unsupported target "keyvalue" with type "example" omitted
# Unsupported target "macro-errors" with type "test" omitted
# Unsupported target "negative_flag" with type "example" omitted
# Unsupported target "nested-subcommands" with type "test" omitted
# Unsupported target "non_literal_attributes" with type "test" omitted
# Unsupported target "options" with type "test" omitted
# Unsupported target "privacy" with type "test" omitted
# Unsupported target "raw_bool_literal" with type "test" omitted
# Unsupported target "raw_idents" with type "test" omitted
# Unsupported target "rename_all" with type "example" omitted
# Unsupported target "rename_all_env" with type "test" omitted
# Unsupported target "skip" with type "example" omitted
# Unsupported target "skip" with type "test" omitted
# Unsupported target "special_types" with type "test" omitted
# Unsupported target "subcommand_aliases" with type "example" omitted
# Unsupported target "subcommands" with type "test" omitted
# Unsupported target "true_or_false" with type "example" omitted
# Unsupported target "utils" with type "test" omitted
