var BridgeableToken = artifacts.require("./BridgeableToken.sol");

module.exports = function(deployer) {
  deployer.deploy(BridgeableToken);
};

