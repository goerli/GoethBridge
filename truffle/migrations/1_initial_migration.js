var Migrations = artifacts.require("./Migrations.sol");
var Bridge = artifacts.require("./Bridge.sol");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(Bridge);
};
