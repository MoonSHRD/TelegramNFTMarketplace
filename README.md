

## Generating ABI
```
solc --abi --bin ./contracts/MetaMarketplace.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,

```


## Generating GO
```
abigen --abi="build/MetaMarketplace.abi" --pkg=MetaMarketplace --out="./go/MetaMarketplace/MetaMarketplace.go"
```



# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```
