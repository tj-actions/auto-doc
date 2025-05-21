{
  description = "auto-doc: Github action that turns your reusable workflows and custom actions into easy to read markdown tables.";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.11";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    system = "x86_64-linux";
  in {
    packages.x86_64-linux.auto-doc = nixpkgs.legacyPackages.${system}.buildGoModule {
      pname = "auto-doc";
      version = "3.6.0";

      src = ./.;
      vendorHash = "sha256-kO5xCO8bjs2vF4CS35odlhz17jPcpvQ6gdrE61p7x/w=";
      doCheck = false;
    };

    packages.x86_64-linux.default = self.packages.x86_64-linux.auto-doc;
  };
}
