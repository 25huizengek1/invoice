{
  buildGoModule,
  fetchFromGitHub,
  ...
}:

buildGoModule (finalAttrs: {
  pname = "go-tools";
  version = "0.35.0";

  src = fetchFromGitHub {
    owner = "golang";
    repo = "tools";
    tag = finalAttrs.version;
    hash = "sha256-ctKWzDupoJNSesiecLxiux6fxJBjehb8md15p2kQ0mg=";
  };

  vendorHash = "sha256-L2VYebgRTdiJyIBW437hvt8RyF4D4P8rjFvjNiDtu6Q=";

  subPackages = [
    "cmd/auth"
    "cmd/benchcmp"
    "cmd/bisect"
    # Tests fail
    # "cmd/bundle"
    "cmd/callgraph"
    "cmd/compilebench"
    "cmd/deadcode"
    "cmd/digraph"
    "cmd/eg"
    "cmd/file2fuzz"
    "cmd/fiximports"
    "cmd/go-contrib-init"
    "cmd/godex"
    # Tests fail
    # "cmd/godoc"
    "cmd/goimports"
    "cmd/gomvpkg"
    "cmd/gonew"
    "cmd/gotype"
    "cmd/goyacc"
    "cmd/html2article"
    "cmd/present"
    "cmd/present2md"
    "cmd/signature-fuzzer"
    "cmd/splitdwarf"
    "cmd/ssadump"
    "cmd/stress"
    "cmd/stringer"
    "cmd/toolstash"
  ];
})
