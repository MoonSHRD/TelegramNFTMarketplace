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
  //const erc20sample = await ERC20Sample.deploy("Token", "TKN", initialSupply);



  //let murs_account = ethers.utils.getAddress("0x383A9e83E36796106EaC11E8c2Fbe8b92Ff46D3a");
  //let account_owner = await hre.ethers.getSigner();
  //const owner = await hre.ethers.utils.getAddress("0x16d97A46030C5D3D705bca45439e48529997D8b2");
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
  const Currencies = await CurrenciesERC20_F.deploy("0xe583769738b6dd4E7CAF8451050d1948BE717679","0x07865c6E87B9F70255377e024ace6630C1Eaa37F","0x73967c6a0904aA032C103b4104747E88c566B1A2","0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6","0xC04B0d3107736C32e19F1c62b2aF67BE61d63a05",our_currency.address);

  await Currencies.deployed();
  console.log("Currencies util contract deployed to:", Currencies.address);


  const MetaMarketplace_F = await hre.ethers.getContractFactory("MetaMarketplace");
  const MetaMarketplace = await MetaMarketplace_F.deploy(Currencies.address,"0xEf087AaF882c22350ABc6A287563cE037AC46b06",owner_account.address);
  await MetaMarketplace.deployed();
  console.log("MetaMarketplace deployed to:", MetaMarketplace.address);


}


// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
