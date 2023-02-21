<div align="center">
  <img src="resources/OvalancheLogo.png?raw=true">
</div>

---

Node implementation for the [Ovalanche](https://avax.network) network -
a blockchains platform with high throughput, blazing fast transactions and free transfers!

### Building OvalancheGo

Build OvalancheGo by running the build script:

```sh
./scripts/build.sh
```

### Running local testnet

https://docs.avax.network/quickstart/create-a-local-test-network

### P-chain endpoints

#### `platform.transfer` (zero fee transfers!)

arguments:
* username
* password
* amount - amount of OVAX to transfer
* to - address to transfer OVAX to

note: to use this endpoint you have to create user and import key via `keystore.createUser` and `platform.importKey` endpoints

#### `platform.previewSpend` (preview a transaction)

arguments:
* amount - amount of OVAX to transfer
* from - address to transfer OVAX from
* to - address to transfer OVAX to

#### `platform.getFlag` (get the winning flag if you have positive balance!)

arguments:
* username
* password

as well as standard `platform` endpoints defined in https://docs.avax.network/apis/avalanchego/apis/p-chain

### Prefunded addresses
* `P-custom1v9wzm63cr5ta2u0n5m7q4repdzfnpnare9uxv8` (300 000 000 OVAX)
* `P-custom17yexyhlpvevk33ydux8upm7fu0mlaw750tsdzp` (300 000 00 OVAX)