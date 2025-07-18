{
  buildGoModule,
  lib,
  ...
}:

buildGoModule (finalAttrs: {
  pname = "invoice";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-5JRX+t5BuImGUUcixvG/M4RFHBu85PedKTj987GHzws=";

  meta = {
    description = "Generate invoices from the command line";
    homepage = "https://github.com/maaslalani/invoice";
    license = lib.licenses.mit;
  };
})
