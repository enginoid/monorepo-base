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


# Unsupported target "01_default" with type "bench" omitted
# Unsupported target "01a_quick_example" with type "example" omitted
# Unsupported target "01b_quick_example" with type "example" omitted
# Unsupported target "01c_quick_example" with type "example" omitted
# Unsupported target "02_apps" with type "example" omitted
# Unsupported target "02_simple" with type "bench" omitted
# Unsupported target "03_args" with type "example" omitted
# Unsupported target "03_complex" with type "bench" omitted
# Unsupported target "04_new_help" with type "bench" omitted
# Unsupported target "04_using_matches" with type "example" omitted
# Unsupported target "05_flag_args" with type "example" omitted
# Unsupported target "05_ripgrep" with type "bench" omitted
# Unsupported target "06_positional_args" with type "example" omitted
# Unsupported target "06_rustup" with type "bench" omitted
# Unsupported target "07_option_args" with type "example" omitted
# Unsupported target "08_subcommands" with type "example" omitted
# Unsupported target "09_auto_version" with type "example" omitted
# Unsupported target "10_default_values" with type "example" omitted
# Unsupported target "11_only_specific_values" with type "example" omitted
# Unsupported target "12_typed_values" with type "example" omitted
# Unsupported target "13a_enum_values_automatic" with type "example" omitted
# Unsupported target "13b_enum_values_manual" with type "example" omitted
# Unsupported target "14_groups" with type "example" omitted
# Unsupported target "15_custom_validator" with type "example" omitted
# Unsupported target "16_app_settings" with type "example" omitted
# Unsupported target "17_yaml" with type "example" omitted
# Unsupported target "18_builder_macro" with type "example" omitted
# Unsupported target "19_auto_authors" with type "example" omitted
# Unsupported target "20_subcommands" with type "example" omitted
# Unsupported target "21_aliases" with type "example" omitted
# Unsupported target "22_stop_parsing_with_--" with type "example" omitted
# Unsupported target "app_from_crate" with type "test" omitted
# Unsupported target "app_settings" with type "test" omitted
# Unsupported target "arg_aliases" with type "test" omitted
# Unsupported target "borrowed" with type "test" omitted
# Unsupported target "cargo" with type "test" omitted

rust_library(
    name = "clap",
    crate_root = "src/lib.rs",
    crate_type = "lib",
    edition = "2018",
    srcs = glob(["**/*.rs"]),
    deps = [
        "@raze__atty__0_2_14//:atty",
        "@raze__bitflags__1_2_1//:bitflags",
        "@raze__clap_derive__3_0_0_beta_1//:clap_derive",
        "@raze__indexmap__1_3_2//:indexmap",
        "@raze__lazy_static__1_4_0//:lazy_static",
        "@raze__strsim__0_9_3//:strsim",
        "@raze__termcolor__1_1_0//:termcolor",
        "@raze__textwrap__0_11_0//:textwrap",
        "@raze__unicode_width__0_1_7//:unicode_width",
        "@raze__vec_map__0_8_1//:vec_map",
    ],
    rustc_flags = [
        "--cap-lints=allow",
    ],
    version = "3.0.0-beta.1",
    crate_features = [
        "atty",
        "cargo",
        "clap_derive",
        "color",
        "default",
        "derive",
        "lazy_static",
        "std",
        "strsim",
        "suggestions",
        "termcolor",
        "vec_map",
    ],
)

# Unsupported target "conflicts" with type "test" omitted
# Unsupported target "default_vals" with type "test" omitted
# Unsupported target "delimiters" with type "test" omitted
# Unsupported target "derive_order" with type "test" omitted
# Unsupported target "env" with type "test" omitted
# Unsupported target "flags" with type "test" omitted
# Unsupported target "global_args" with type "test" omitted
# Unsupported target "groups" with type "test" omitted
# Unsupported target "help" with type "test" omitted
# Unsupported target "hidden_args" with type "test" omitted
# Unsupported target "indices" with type "test" omitted
# Unsupported target "macros" with type "test" omitted
# Unsupported target "multiple_occurrences" with type "test" omitted
# Unsupported target "multiple_values" with type "test" omitted
# Unsupported target "opts" with type "test" omitted
# Unsupported target "positionals" with type "test" omitted
# Unsupported target "posix_compatible" with type "test" omitted
# Unsupported target "possible_values" with type "test" omitted
# Unsupported target "propagate_globals" with type "test" omitted
# Unsupported target "require" with type "test" omitted
# Unsupported target "subcommands" with type "test" omitted
# Unsupported target "template_help" with type "test" omitted
# Unsupported target "tests" with type "test" omitted
# Unsupported target "unique_args" with type "test" omitted
# Unsupported target "utf8" with type "test" omitted
# Unsupported target "utils" with type "test" omitted
# Unsupported target "version" with type "test" omitted
# Unsupported target "version-numbers" with type "test" omitted
# Unsupported target "yaml" with type "test" omitted
