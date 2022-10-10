"""This module defines rules for building Skia Infrastructure web applications."""

load("@build_bazel_rules_nodejs//:index.bzl", "npm_package_bin", _nodejs_test = "nodejs_test")
load("@io_bazel_rules_docker//container:flatten.bzl", "container_flatten")
load("@io_bazel_rules_sass//:defs.bzl", "sass_binary")
load("//bazel/test_on_env:test_on_env.bzl", "test_on_env")
load("//infra-sk/html_insert_assets:index.bzl", "html_insert_assets")
load("//infra-sk/karma_test:index.bzl", _karma_test = "karma_test")
load("//infra-sk/sk_demo_page_server:index.bzl", _sk_demo_page_server = "sk_demo_page_server")
load("//infra-sk/esbuild:esbuild.bzl", "esbuild_dev_bundle", "esbuild_prod_bundle")
load(":ts_library.bzl", _ts_library = "ts_library")
load(":sass_library.bzl", _sass_library = "sass_library")

# https://github.com/bazelbuild/bazel-skylib/blob/main/rules/common_settings.bzl
load("@bazel_skylib//rules:common_settings.bzl", skylib_bool_flag = "bool_flag")

# Re-export these common rules so we only have to load this .bzl file from our BUILD.bazel files.
karma_test = _karma_test
sass_library = _sass_library
sk_demo_page_server = _sk_demo_page_server
ts_library = _ts_library

def sk_element(
        name,
        ts_srcs,
        sass_srcs = [],
        ts_deps = [],
        sass_deps = [],
        sk_element_deps = [],
        visibility = None):
    """Defines a custom element for Skia Infrastructure web applications.

    This is just a convenience macro that generates the ts_library and sass_library targets
    required to build a custom element.

    This macro generates a "ghost" Sass stylesheet which includes `@import` statements for each
    sk_element dependency (sk_element_deps argument), and for any elements-sk modules imported from
    the TypeScript sources (ts_srcs argument). This ghost stylesheet will be included in the CSS
    bundles produced by any sk_page targets that depend on this element (directly or transitively).
    It is therefore *not* necessary to explicitly import the stylesheets of any depended on
    sk_elements or elements-sk modules.

    Args:
      name: The name of the target.
      ts_srcs: TypeScript source files.
      sass_srcs: Sass source files.
      ts_deps: Any ts_library dependencies.
      sass_deps: Any sass_library dependencies. This can include .css or .scss files from NPM
        modules, e.g. "npm//:node_modules/some-module/hello.scss".
      sk_element_deps: Any sk_element dependencies. Equivalent to adding the ts_library and
        sass_library of each sk_element to ts_deps and sass_deps, respectively.
      visibility: Visibility of the generated ts_library and sass_library targets.
    """

    # Generate a Sass stylesheet with import statements for each elements-sk module required by the
    # TypeScript sources (ts_srcs argument).
    generate_sass_stylesheet_with_elements_sk_imports_from_typescript_sources(
        name = name + "_elements_sk_deps_scss",
        ts_srcs = ts_srcs,
        scss_output_file = name + "__generated_elements_sk_deps.scss",
    )

    # Generate a "ghost" Sass entry-point stylesheet with import statements for the following
    # files:
    #
    #  - Each file in the sass_srcs argument.
    #  - The generated Sass stylesheet with elements-sk imports.
    #  - The "ghost" entry-point stylesheets of each sk_element in the sk_element_deps argument.
    #
    # This stylesheet will be included in the CSS bundles generated by any sk_page targets that
    # depend on this sk_element directly or transitively.
    generate_sass_stylesheet_with_imports(
        name = name + "_ghost_entrypoint_scss",
        scss_files_to_import = sass_srcs + [name + "_elements_sk_deps_scss"] +
                               [
                                   make_label_target_explicit(dep) + "_ghost_entrypoint_scss"
                                   for dep in sk_element_deps
                               ],
        scss_output_file = name + "__generated_ghost_entrypoint.scss",
        visibility = visibility,
    )

    # Extend ts_deps and sass_deps with the ts_library and sass_library targets produced by each
    # sk_element dependency in the sk_element_deps argument.
    all_ts_deps = [dep for dep in ts_deps]
    all_sass_deps = [dep for dep in sass_deps]
    for sk_element_dep in sk_element_deps:
        all_ts_deps.append(sk_element_dep)
        all_sass_deps.append(make_label_target_explicit(sk_element_dep) + "_styles")

    ts_library(
        name = name,
        srcs = ts_srcs,
        deps = all_ts_deps,
        visibility = visibility,
    )

    sass_library(
        name = name + "_styles",
        srcs = sass_srcs + [
            name + "_elements_sk_deps_scss",
            name + "_ghost_entrypoint_scss",
        ],
        deps = all_sass_deps,
        visibility = visibility,
    )

def generate_sass_stylesheet_with_elements_sk_imports_from_typescript_sources(
        name,
        ts_srcs,
        scss_output_file):
    """Generates a .scss file with `@import` statements for the elements-sk imports in the ts_srcs.

    Args:
      name: The name of the target.
      ts_srcs: A list of .ts files.
      scss_output_file: Name of the .scss file to generate.
    """
    native.genrule(
        name = name,
        srcs = ts_srcs + ["//infra-sk/generate_elements_sk_scss_imports"],
        outs = [scss_output_file],
        cmd = "$(location //infra-sk/generate_elements_sk_scss_imports) $(SRCS) > $@",
    )

def generate_sass_stylesheet_with_imports(name, scss_files_to_import, scss_output_file, visibility = None):
    """Generates a .scss file with one `@import` statement for each file in scss_files_to_import.

    Args:
      name: The name of the target.
      scss_files_to_import: A list of .scss files.
      scss_output_file: Name of the .scss file to generate.
      visibility: Visibility of the target.
    """

    # Build a list of shell commands to generate the output stylesheet.
    cmds = ["touch $@"]
    for scss_file in scss_files_to_import:
        import_statement = "@import '$(rootpath %s)';" % scss_file
        cmds.append("echo \"%s\" >> $@" % import_statement)

    native.genrule(
        name = name,
        srcs = scss_files_to_import,
        outs = [scss_output_file],
        cmd = " && ".join(cmds),
        visibility = visibility,
    )

def make_label_target_explicit(label):
    """Takes a label with a potentially implicit target name, and makes the target name explicit.

    For example, if the label is "//path/to/pkg", this macro will return "//path/to/pkg:pkg". If the
    label is already in the latter form, the label will be returned unchanged.

    Reference: https://docs.bazel.build/versions/master/build-ref.html#labels

    Args:
      label: A Bazel label.

    Returns:
      The given label, expanded to make the target name explicit.
     """
    if label.find(":") != -1:
        return label
    pkg_name = label.split("/").pop()
    return label + ":" + pkg_name

def nodejs_test(
        name,
        src,
        data = [],
        deps = [],
        tags = [],
        visibility = None,
        env = {},
        wait_for_debugger = False,
        _internal_skip_naming_convention_enforcement = False):
    """Runs a Node.js unit test using the Mocha test runner.

    For tests that should run in the browser, please use karma_test instead.

    To debug the Mocha test, set wait_for_debugger to True, then attach your debugger using the URL
    printed to stdout.

    Example debugging session with Chrome DevTools (assumes wait_for_debugger = True):

    1. Add a `debugger` statement in your test code (e.g. //path/to/foo_nodejs_test.ts):
    ```
    describe('foo', () => {
      it('should do something', () => {
        debugger;
        ...
      })
    })
    ```
    2. Run `bazel run //path/to:foo_nodejs_test`.
    3. Launch Chrome **in the machine where the test is running**, otherwise Chrome won't see the
       Node.js process.
    4. Enter chrome://inspect in the URL bar, then press return.
    5. You should see an "inspect" link under the "Remote Target" heading.
    6. Click that link to launch a Chrome DevTools window attached to your Node.js process.
    7. Click the "Resume script execution" button (looks like a play/pause icon).
    8. Test execution should start, and eventually pause at your `debugger` statement.

    Args:
      name: Name of the target.
      src: A single TypeScript source file.
      data: Any data dependencies.
      deps: Any ts_library dependencies.
      tags: Tags for the generated nodejs_test rule.
      visibility: Visibility of the generated nodejs_test rule.
      env: A dictionary of additional environment variables to set when the target is executed.
      wait_for_debugger: Whether to invoke the Mocha test runner with --inspect-brk. If true,
        the Node.js interpreter will wait for a debugger (such as Chrome DevTools) to be attached
        before continuing.
      _internal_skip_naming_convention_enforcement: Not part of the public API - do not use.
    """

    # This macro is called by sk_element_puppeteer_test, which uses a different naming convention.
    if not _internal_skip_naming_convention_enforcement and not src.endswith("_nodejs_test.ts"):
        fail("Node.js tests must end with \"_nodejs_test.ts\".")

    mocha_deps = [
        "@npm//mocha",
        "@npm//ts-node",
        "//:tsconfig.json",
    ]

    _nodejs_test(
        name = name,
        entry_point = "@npm//:node_modules/mocha/bin/mocha",
        data = data + [src] + deps + [dep for dep in mocha_deps if dep not in deps],
        templated_args = [
            "--require ts-node/register/transpile-only",
            "--timeout 60000",
            "--colors",
            # See https://github.com/bazelbuild/rules_nodejs/commit/fdde32fa5653999b15459c4deebfeaa86a099135.
            "--bazel_patch_module_resolver",
            "$(rootpath %s)" % src,
        ] + (["--inspect-brk"] if wait_for_debugger else []),
        env = env,
        tags = tags,
        visibility = visibility,
    )

def sk_element_puppeteer_test(name, src, sk_demo_page_server, deps = []):
    """Defines a Puppeteer test for the demo page served by an sk_demo_page_server.

    Puppeteer tests should save any screenshots inside the $TEST_UNDECLARED_OUTPUTS_DIR directory.
    To reduce the chances of name collisions, tests must save their screenshots under the
    $TEST_UNDECLARED_OUTPUTS_DIR/puppeteer-test-screenshots subdirectory. This convention will
    allow us to recover screenshots from multiple tests in a consistent way.

    Screenshots, and any other undeclared outputs of a test, can be found under //_bazel_testlogs
    bundled as a single .zip file per test target. For example, if we run a Puppeteer test with e.g.
    "bazel test //path/to/my:puppeteer_test", any screenshots taken by this test will be found
    inside //_bazel_testlogs/path/to/my/puppeteer_test/test.outputs/outputs.zip.

    To read more about undeclared test outputs, please see the following link:
    https://docs.bazel.build/versions/master/test-encyclopedia.html#test-interaction-with-the-filesystem.

    This rule also generates a "<name>_debug" test target that waits for a debugger (such as
    Chrome DevTools, or the VS Code Node.js debugger) to be attached to the Node.js process running
    your test before continuing execution. For example, if your test target is
    //path/to:my_puppeteer_test, invoke `bazel test //path/to:my_puppeteer_test`, to run your test
    the usual way, or invoke `bazel run //path/to:my_puppeteer_test_debug` if you'd like to attach
    a debugger. Tip: add one or more `debugger` statement in your test code in order to set
    breakpoints. See the nodejs_test rule's docstring for an example debug session.

    Additionally, this rule generates a "<name>_debug_headful" target that is identical to the
    aforementioned "<name>_debug" target, except that the Chromium instance started by Puppeteer
    runs in headful mode. Use this target to visually inspect how your test interacts with the demo
    page under test as you step through your test code with the attached debugger. Example
    invocation: `bazel run //path/to:my_puppeteer_test_debug_headful`.

    Args:
      name: Name of the rule.
      src: A single TypeScript source file.
      sk_demo_page_server: Label for the sk_demo_page_server target.
      deps: Any ts_library dependencies.
    """

    if not src.endswith("_puppeteer_test.ts"):
        fail("Puppeteer tests must end with \"_puppeteer_test.ts\".")

    for debug, headful in [(False, False), (True, False), (True, True)]:
        suffix = ""
        if debug:
            suffix += "_debug"
        if headful:
            suffix += "_headful"

        nodejs_test(
            name = name + "_test_only" + suffix,
            src = src,
            tags = ["manual"],  # Exclude it from wildcards, e.g. "bazel test //...".
            deps = deps,
            wait_for_debugger = debug,
            env = {"PUPPETEER_TEST_SHOW_BROWSER": "true"} if headful else {},
            _internal_skip_naming_convention_enforcement = True,
        )

        test_on_env(
            name = name + suffix,
            env = sk_demo_page_server,
            test = name + "_test_only" + suffix,
            timeout_secs = 60 * 60 * 24 * 365 if debug else 10,
            tags = ([
                "manual",  # Exclude it from wildcards, e.g. "bazel test //...".
                "no-remote",  # Do not run on RBE.
            ] if debug or headful else []),
        )

def copy_file(name, src, dst, visibility = None):
    """Copies a single file to a destination path, making parent directories as needed."""
    native.genrule(
        name = name,
        srcs = [src],
        outs = [dst],
        cmd = "mkdir -p $$(dirname $@) && cp $< $@",
        visibility = visibility,
    )

def sk_page(
        name,
        html_file,
        ts_entry_point,
        scss_entry_point = None,
        ts_deps = [],
        sass_deps = [],
        sk_element_deps = [],
        assets_serving_path = "/",
        copy_files = None,
        nonce = None):
    """Builds a static HTML page, and its CSS and JavaScript development and production bundles.

    This macro generates the following files, where <name> is the given target name:

        development/<name>.html
        development/<name>.js
        development/<name>.css
        production/<name>.html
        production/<name>.js
        production/<name>.css

    The <name> target defined by this macro generates all of the above files.

    Tags <script> and <link> will be inserted into the output HTML pointing to the generated
    bundles. The serving path for said bundles defaults to "/" and can be overridden via the
    assets_serving_path argument.

    A timestamp will be appended to the URLs for any referenced assets for cache busting purposes,
    e.g. <script src="/index.js?v=27396986"></script>.

    If the nonce argument is provided, a nonce attribute will be inserted to all <link> and <script>
    tags. For example, if the nonce argument is set to "{% .Nonce %}", then the generated HTML will
    contain tags such as <script nonce="{% .Nonce %}" src="/index.js?v=27396986"></script>.

    Args:
      name: The prefix used for the names of all the targets generated by this macro.
      html_file: The page's HTML file.
      ts_entry_point: TypeScript file used as the entry point for the JavaScript bundles.
      scss_entry_point: Sass file used as the entry point for the CSS bundles.
      ts_deps: Any ts_library dependencies.
      sass_deps: Any sass_library dependencies. This can include .css or .scss files from NPM
        modules, e.g. "npm//:node_modules/some-module/hello.scss".
      sk_element_deps: Any sk_element dependencies. Equivalent to adding the ts_library and
        sass_library of each sk_element to deps and sass_deps, respectively.
      assets_serving_path: Path prefix for the inserted <script> and <link> tags.
      copy_files: Any files that should just be copied into the final build directory. These are
        assets needed by the page that are not loaded in via imports (e.g. images, WASM).
      nonce: If set, its contents will be added as a "nonce" attribute to any <script> and <link>
        tags inserted into the page's HTML file.
    """

    # Extend ts_deps and sass_deps with the ts_library and sass_library targets produced by each
    # sk_element dependency in the sk_element_deps argument.
    all_ts_deps = [dep for dep in ts_deps]
    all_sass_deps = [dep for dep in sass_deps]
    for sk_element_dep in sk_element_deps:
        all_ts_deps.append(sk_element_dep)
        all_sass_deps.append(make_label_target_explicit(sk_element_dep) + "_styles")

    # Output directories.
    DEV_OUT_DIR = "development"
    PROD_OUT_DIR = "production"

    #######################
    # JavaScript bundles. #
    #######################

    ts_library(
        name = "%s_ts_lib" % name,
        srcs = [ts_entry_point],
        deps = all_ts_deps,
    )

    # Generates file development/<name>.js.
    esbuild_dev_bundle(
        name = "%s_js_dev" % name,
        entry_point = ts_entry_point,
        deps = [":%s_ts_lib" % name],
        output = "%s/%s.js" % (DEV_OUT_DIR, name),
    )

    # Generates file production/<name>.js.
    esbuild_prod_bundle(
        name = "%s_js_prod" % name,
        entry_point = ts_entry_point,
        deps = [":%s_ts_lib" % name],
        output = "%s/%s.js" % (PROD_OUT_DIR, name),
    )

    ################
    # CSS Bundles. #
    ################

    # Generate a blank Sass entry-point file to appease the sass_library rule, if one is not given.
    if not scss_entry_point:
        scss_entry_point = name + "__generated_empty_scss_entry_point"
        native.genrule(
            name = scss_entry_point,
            outs = [scss_entry_point + ".scss"],
            cmd = "touch $@",
        )

    # Generate a Sass stylesheet with any elements-sk imports required by the TypeScript
    # entry-point file.
    generate_sass_stylesheet_with_elements_sk_imports_from_typescript_sources(
        name = name + "_elements_sk_deps_scss",
        ts_srcs = [ts_entry_point],
        scss_output_file = name + "__generated_elements_sk_deps.scss",
    )

    # Create a sass_library including the scss_entry_point file, and all the Sass dependencies.
    sass_library(
        name = name + "_styles",
        srcs = [
            scss_entry_point,
            name + "_elements_sk_deps_scss",
        ],
        deps = all_sass_deps,
    )

    # Generate a "ghost" Sass entry-point stylesheet with import statements for the following
    # files:
    #
    #  - This page's Sass entry-point file (scss_entry_point argument).
    #  - The generated Sass stylesheet with elements-sk imports.
    #  - The "ghost" entry-point stylesheets of each sk_element in the sk_element_deps argument.
    #
    # We will use this generated stylesheet as the entry-points for the sass_binaries below.
    generate_sass_stylesheet_with_imports(
        name = name + "_ghost_entrypoint_scss",
        scss_files_to_import = ([scss_entry_point] if scss_entry_point else []) +
                               [name + "_elements_sk_deps_scss"] +
                               [
                                   make_label_target_explicit(dep) + "_ghost_entrypoint_scss"
                                   for dep in sk_element_deps
                               ],
        scss_output_file = name + "__generated_ghost_entrypoint.scss",
    )

    # Notes:
    #  - Sass compilation errors are not visible unless "bazel build" is invoked with flag
    #    "--strategy=SassCompiler=sandboxed" (now set by default in //.bazelrc). This is due to a
    #    known issue with sass_binary. For more details please see
    #    https://github.com/bazelbuild/rules_sass/issues/96.

    # Generates file development/<name>.css.
    sass_binary(
        name = "%s_css_dev" % name,
        src = name + "_ghost_entrypoint_scss",
        output_name = "%s/%s.css" % (DEV_OUT_DIR, name),
        deps = [name + "_styles"],
        include_paths = [
            "//external/npm",  # Allows @use "node_modules/some_package/some_file.css" to work
        ],
        output_style = "expanded",
        sourcemap = True,
        sourcemap_embed_sources = True,
        visibility = ["//visibility:public"],
    )

    # Generates file <name>_unoptimized.css.
    sass_binary(
        name = "%s_css_unoptimized_prod" % name,
        src = name + "_ghost_entrypoint_scss",
        output_name = "%s_unoptimized.css" % name,
        deps = [name + "_styles"],
        include_paths = [
            "//external/npm",  # Allows @use "node_modules/some_package/some_file.css" to work
        ],
        output_style = "compressed",
        sourcemap = False,
        visibility = ["//visibility:public"],
    )

    # Generates file production/<name>.css.
    #
    # That sass tool used by `sass_binary` doesn't remove duplicate CSS rules,
    # so the output can contain many copies of the CSS rules like 'colors.scss'.
    # We pass the CSS through csso to remove duplicate CSS rules, which can
    # reduce a file to 1/10 its unoptimized size.
    npm_package_bin(
        name = "%s_css_prod" % name,
        tool = "@npm//csso-cli/bin:csso",
        chdir = "$(RULEDIR)",
        data = [":%s_css_unoptimized_prod" % name],
        stdout = "%s/%s.css" % (PROD_OUT_DIR, name),
        args = ["%s_unoptimized.css" % name],
        visibility = ["//visibility:public"],
    )

    #####################
    # Static resources. #
    #####################

    if copy_files:
        for pair in copy_files:
            copy_file(
                name = "%s_copy_prod" % pair["dst"],
                src = pair["src"],
                dst = PROD_OUT_DIR + "/" + pair["dst"],
                visibility = ["//visibility:public"],
            )
            copy_file(
                name = "%s_copy_dev" % pair["dst"],
                src = pair["src"],
                dst = DEV_OUT_DIR + "/" + pair["dst"],
                visibility = ["//visibility:public"],
            )

    ###############
    # HTML files. #
    ###############

    if assets_serving_path.endswith("/"):
        assets_serving_path = assets_serving_path[:-1]

    # Generates file development/<name>.html.
    html_insert_assets(
        name = "%s_html_dev" % name,
        html_src = html_file,
        html_out = "%s/%s.html" % (DEV_OUT_DIR, name),
        js_src = "%s/%s.js" % (DEV_OUT_DIR, name),
        js_serving_path = "%s/%s.js" % (assets_serving_path, name),
        css_src = "%s/%s.css" % (DEV_OUT_DIR, name),
        css_serving_path = "%s/%s.css" % (assets_serving_path, name),
        nonce = nonce,
    )

    # Generates file production/<name>.html.
    html_insert_assets(
        name = "%s_html_prod" % name,
        html_src = html_file,
        html_out = "%s/%s.html" % (PROD_OUT_DIR, name),
        js_src = "%s/%s.js" % (PROD_OUT_DIR, name),
        js_serving_path = "%s/%s.js" % (assets_serving_path, name),
        css_src = "%s/%s.css" % (PROD_OUT_DIR, name),
        css_serving_path = "%s/%s.css" % (assets_serving_path, name),
        nonce = nonce,
    )

    ###########################
    # Convenience filegroups. #
    ###########################

    # Generates all output files (that is, the development and production bundles).
    native.filegroup(
        name = name,
        srcs = [
            ":%s_dev" % name,
            ":%s_prod" % name,
        ],
        visibility = ["//visibility:public"],
    )

    # Generates the development bundle.
    native.filegroup(
        name = "%s_dev" % name,
        srcs = [
            "development/%s.html" % name,
            "development/%s.js" % name,
            "development/%s.css" % name,
            "development/%s.css.map" % name,
        ],
        visibility = ["//visibility:public"],
    )

    # Generates the production bundle.
    native.filegroup(
        name = "%s_prod" % name,
        srcs = [
            "production/%s.html" % name,
            "production/%s.js" % name,
            "production/%s.css" % name,
        ],
        visibility = ["//visibility:public"],
    )

def extract_files_from_skia_wasm_container(name, container_files, outs, enabled_flag, **kwargs):
    """Extracts files from the Skia WASM container image (gcr.io/skia-public/skia-wasm-release).

    This macro takes as inputs a list of paths inside the Docker container (container_files
    argument), and a list of the same length with the destination paths for each of the files to
    extract (outs argument), relative to the directory where the macro is instantiated.

    This image will be pulled if the enabled_flag is true, so users who need to set that flag to
    true (e.g. infra folks pushing a clean image) should be granted permissions and have Docker
    authentication set up. Users who are working with a local build should set the flag to false
    and not need Docker authentication.

    Args:
      name: Name of the target.
      container_files: List of absolute paths inside the Docker container to extract.
      outs: Destination paths for each file to extract, relative to the target's directory.
      enabled_flag: Label. If set, should be the name of a bool_flag. If the bool_flag
          is True, the real skia_wasm_container image will be pulled and the images loaded as per
          usual. If the bool_flag is false, the container will not be loaded, but any outs will be
          created as empty files to make dependent rules happy. It is up to the caller to use that
          same flag to properly ignore those empty files if the flag is false (e.g. via a select).
      **kwargs: Any flags that should be forwarded to the generated rule
    """

    if len(container_files) != len(outs):
        fail("Arguments container_files and outs must have the same length.")

    # Generates a .tar file with the contents of the image's filesystem (and a .json metadata file
    # which we ignore).
    #
    # See the rule implementation here:
    # https://github.com/bazelbuild/rules_docker/blob/02ad0a48fac9afb644908a634e8b2139c5e84670/container/flatten.bzl#L48
    #
    # Notes:
    #  - It is unclear whether container_flatten is part of the public API because it is not
    #    documented. But the fact that container_flatten is re-exported along with other, well
    #    documented rules in rules_docker suggests that it might indeed be part of the public API
    #    (see [1] and [2]).
    #  - If they ever make breaking changes to container_flatten, then it is probably best to fork
    #    it. The rule itself is relatively small; it is just a wrapper around a Go program that does
    #    all the heavy lifting.
    #  - This rule was chosen because most other rules in the rules_docker repository produce .tar
    #    files with layered outputs, which means we would have to do the flattening ourselves.
    #
    # [1] https://github.com/bazelbuild/rules_docker/blob/6c29619903b6bc533ad91967f41f2a3448758e6f/container/container.bzl#L28
    # [2] https://github.com/bazelbuild/rules_docker/blob/e290d0975ab19f9811933d2c93be275b6daf7427/container/BUILD#L158

    # We expect enabled_flag to be a bool_flag, which means we should have two labels defined
    # based on the passed in flag name, one with a _true suffix and one with a _false suffix. If
    # the flag is set to false, we pull from a public image that has no files in it.
    container_flatten(
        name = name + "_skia_wasm_container_filesystem",
        image = select({
            enabled_flag + "_true": "@container_pull_skia_wasm//image",
            enabled_flag + "_false": "@empty_container//image",
        }),
    )

    # Name of the .tar file produced by the container_flatten target.
    skia_wasm_filesystem_tar = name + "_skia_wasm_container_filesystem.tar"

    # Shell command that returns the directory of the BUILD file where this macro was instantiated.
    #
    # This works because:
    #  - The $< variable[1] is expanded by the below genrule to the path of its only input file.
    #  - The only input file to the genrule is the .tar file produced by the above container_flatten
    #    target (see the genrule's srcs attribute).
    #  - Said .tar file is produced in the directory of the BUILD file where this macro was
    #    instantiated.
    #
    # [1] https://bazel.build/reference/be/make-variables
    output_dir = "$$(dirname $<)"

    # Directory where we will untar the .tar file produced by the container_flatten target.
    skia_wasm_filesystem_dir = output_dir + "/" + skia_wasm_filesystem_tar + "_untarred"

    # Untar the .tar file produced by the container_flatten rule.
    cmd = ("mkdir -p " + skia_wasm_filesystem_dir +
           " && tar xf $< -C " + skia_wasm_filesystem_dir)

    # Copy each requested file from the container filesystem to its desired destination.
    for src, dst in zip(container_files, outs):
        copy = " && (cp %s/%s %s/%s" % (skia_wasm_filesystem_dir, src, output_dir, dst)

        # If the enabled_flag is False, we will be loading from an empty image and thus the
        # files will not exist. As such, we need to create them or Bazel will fail because the
        # expected generated files will not be created. It is up to the dependent rules to properly
        # ignore these files. We cannot easily have the genrule change its behavior depending on
        # how the flag is set, so we always touch the file as a backup. We write stderr to
        # /dev/null to squash any errors about files not existing.
        copy += " 2>/dev/null || touch %s/%s) " % (output_dir, dst)
        cmd += copy

    native.genrule(
        name = name,
        srcs = [skia_wasm_filesystem_tar],
        outs = outs,
        cmd = cmd,
        **kwargs
    )

def bool_flag(flag_name, default = True, name = ""):  # buildifier: disable=unused-variable
    """Create a boolean flag and corresponding config_settings.

    bool_flag is a Bazel Macro that defines a boolean flag with the given name two config_settings,
    one for True, one for False. Reminder that Bazel has special syntax for unsetting boolean flags,
    but this does not work well with aliases.
    https://docs.bazel.build/versions/main/skylark/config.html#using-build-settings-on-the-command-line
    Thus it is best to define both an "enabled" alias and a "disabled" alias.

    Args:
        flag_name: string, the name of the flag to create and use for the config_settings
        default: boolean, if the flag should default to on or off.
        name: string unused, https://github.com/bazelbuild/buildtools/blob/master/WARNINGS.md#unnamed-macro
    """
    skylib_bool_flag(name = flag_name, build_setting_default = default)

    native.config_setting(
        name = flag_name + "_true",
        flag_values = {
            # The value must be a string, but it will be parsed to a boolean
            # https://docs.bazel.build/versions/main/skylark/config.html#build-settings-and-select
            ":" + flag_name: "True",
        },
    )

    native.config_setting(
        name = flag_name + "_false",
        flag_values = {
            ":" + flag_name: "False",
        },
    )
