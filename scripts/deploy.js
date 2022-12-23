// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const { ethers } = require("hardhat");
const { toWei, fromWei } = require("./lib.js");
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  console.log(hre.network.name);

  let initialSupply = toWei(1000000);
 

  // Polygon addresses
  const USDT = "0xc2132D05D31c914a87C6611C10748AEb04B58e8F";
  const USDC = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174";
  const DAI = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174";
  const WBTC = "0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6";
  const WETH = "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619";


  const accounts = await hre.ethers.getSigners();
  for (const account of accounts) {
    console.log(account.address);
  }
  owner_account = accounts[0];
  const balance = await ethers.provider.getBalance(owner_account.address);
  console.log(ethers.utils.formatEther(balance), "ETH");
  
  console.log("owner address:", owner_account.address);




  const ERC20Sample_F = await hre.ethers.getContractFactory("ERC20Sample");
  const our_currency = await ERC20Sample_F.deploy("VoxPopuly", "VXPPL",initialSupply);
  await our_currency.deployed();
  console.log("Our currency deployed to: ", our_currency.address);

  // We get the contract to deploy
  const CurrenciesERC20_F = await hre.ethers.getContractFactory("CurrenciesERC20");
  const Currencies = await CurrenciesERC20_F.deploy(USDT,USDC,DAI,WETH,WBTC,our_currency.address);

  await Currencies.deployed();
  console.log("Currencies util contract deployed to:", Currencies.address);


  const MetaMarketplace_F = await hre.ethers.getContractFactory("MetaMarketplace");
  const MetaMarketplace = await MetaMarketplace_F.deploy(Currencies.address,"0xEbE648689E98abA446e38621E5a3491db03a7621",owner_account.address);
  await MetaMarketplace.deployed();
  console.log("MetaMarketplace deployed to:", MetaMarketplace.address);


}


// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
