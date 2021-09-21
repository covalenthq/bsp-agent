const { deployProxy } = require('@openzeppelin/truffle-upgrades');

const ProofChain = artifacts.require("ProofChain");

module.exports = async function(deployer, network, accounts) {
  const instance = await deployProxy(ProofChain, [accounts[0]], { deployer });
  console.log("Deployed", instance.address);
};
