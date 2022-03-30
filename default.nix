let
  pkgs = import <nixpkgs> {};
  unstable = import <nixpkgs-unstable> {};

  shell = pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages; [
      podman podman-compose fuse-overlayfs python39 python39Packages.tika unstable.go_1_18 sqlite
    ];
    shellHook = ''
        export TMPDIR=~/.local/tmp
        mkdir -p $TMPDIR
    '';
  };
in shell
