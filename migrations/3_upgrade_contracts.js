const { upgradeProxy } = require('@openzeppelin/truffle-upgrades');

const ProofChain = artifacts.require("ProofChain");

module.exports = async function(deployer) {
  const existing = await ProofChain.deployed();
  const instance = await upgradeProxy(existing.address, ProofChain, { deployer });
  console.log("Upgraded", instance.address);
};
