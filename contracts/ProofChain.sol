// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

contract ProofChain is Initializable {
	using CountersUpgradeable for CountersUpgradeable.Counter;
	using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;
	using EnumerableSetUpgradeable for EnumerableSetUpgradeable.Bytes32Set;

  struct RoleData {
    mapping (address => bool) preapprovals;
    mapping (address => bool) members;
    uint256 requiredStake;
    bytes32 adminRole;
  }

  bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;
	bytes32 public constant GOVERNANCE_ROLE = keccak256("GOVERNANCE_ROLE");
	bytes32 public constant BLOCK_SPECIMEN_PRODUCER_ROLE = keccak256("BLOCK_SPECIMEN_PRODUCER_ROLE");
	bytes32 public constant STAKING_ORACLE_ROLE = keccak256("STAKING_ORACLE_ROLE");

  event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole);
  event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender);
  event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender);

	event BlockSpecimenPublicationProofAppended(
		uint64 seq,           	// sequential ID of this *Appended log-event,
														// among all *Appended events emitted by this contract
														// -- equivalent to a block height

		address extractWorker,	// submitter of the proof
		uint64 chainID,       	// chainID the specimen pertains to
		uint64 chainHeightPos,	// height of first block contained in specimen
		uint64 chainHeightLen,	// number of contiguous blocks contained in specimen
														// (block specimen may only contain contiguous blocks)
		uint64 specimenSize,  	// specimen object file size, measured in bytes
		uint256 specimenHash		// SHA-256 content-hash of specimen object file;
														// used to retrieve specimen from IPFS
	);

  mapping (bytes32 => RoleData) private roles;
	EnumerableSetUpgradeable.Bytes32Set rolesWithRequiredStake;

	EnumerableSetUpgradeable.AddressSet stakers;
	mapping (address => uint256) stakedBalances;

	CountersUpgradeable.Counter proofSeq;

	function initialize(address initialOwner) public initializer {
		grantRole(DEFAULT_ADMIN_ROLE, initialOwner);
		grantRole(GOVERNANCE_ROLE, initialOwner);
		setRoleAdmin(STAKING_ORACLE_ROLE, GOVERNANCE_ROLE);
		setRoleAdmin(BLOCK_SPECIMEN_PRODUCER_ROLE, GOVERNANCE_ROLE);
	}

  function hasRole(bytes32 role, address account) public view returns (bool) {
    return roles[role].members[account];
  }

  function getRoleAdmin(bytes32 role) public view returns (bytes32) {
    return roles[role].adminRole;
  }

  function setRoleAdmin(bytes32 role, bytes32 adminRole) internal {
    emit RoleAdminChanged(role, getRoleAdmin(role), adminRole);
    roles[role].adminRole = adminRole;
  }

  function grantRole(bytes32 role, address account) internal {
    if (!hasRole(role, account)) {
      roles[role].members[account] = true;
      emit RoleGranted(role, account, msg.sender);
    }
  }

  function revokeRole(bytes32 role, address account) internal {
    if (hasRole(role, account)) {
      delete roles[role].members[account];
      emit RoleRevoked(role, account, msg.sender);
    }
  }

	function getRequiredStakeForRole(bytes32 role) public view returns (uint256) {
		return roles[role].requiredStake;
	}

	function setRequiredStakeForRole(bytes32 role, uint256 amount) public {
		require(hasRole(GOVERNANCE_ROLE, msg.sender));

		uint256 prevRequiredStakeForRole = roles[role].requiredStake;

		roles[role].requiredStake = amount;

		if(prevRequiredStakeForRole == 0 && amount > 0) {
			rolesWithRequiredStake.add(role);
		} else if(prevRequiredStakeForRole > 0 && amount == 0) {
			rolesWithRequiredStake.remove(role);
		}

		if(prevRequiredStakeForRole < amount) {
			uint256 numStakers = stakers.length();
			for(uint256 i = 0; i < numStakers; i++) {
				revokeRoleIfNoLongerSufficientlyStaked(role, stakers.at(i));
			}
		} else if(prevRequiredStakeForRole > amount) {
			uint256 numStakers = stakers.length();
			for(uint256 i = 0; i < numStakers; i++) {
				grantRoleIfPreapprovedAndSufficientlyStaked(role, stakers.at(i));
			}
		}
	}

	function getStakedBalance(address addr) public view returns (uint256) {
		return stakedBalances[addr];
	}

	function setStakedBalance(address addr, uint256 amount) public {
		require(hasRole(STAKING_ORACLE_ROLE, msg.sender));

		uint256 prevStakedBalance = stakedBalances[addr];

		stakedBalances[addr] = amount;

		if(prevStakedBalance == 0 && amount > 0) {
			stakers.add(addr);
		} else if(prevStakedBalance > 0 && amount == 0) {
			stakers.remove(addr);
		}

		if(prevStakedBalance < amount) {
			uint256 numRolesWithRequiredStake = rolesWithRequiredStake.length();
			for(uint256 i = 0; i < numRolesWithRequiredStake; i++) {
				grantRoleIfPreapprovedAndSufficientlyStaked(rolesWithRequiredStake.at(i), addr);
			}
		} else if(prevStakedBalance > amount) {
			uint256 numRolesWithRequiredStake = rolesWithRequiredStake.length();
			for(uint256 i = 0; i < numRolesWithRequiredStake; i++) {
				revokeRoleIfNoLongerSufficientlyStaked(rolesWithRequiredStake.at(i), addr);
			}
		}
	}

	function isPreapprovedForRole(bytes32 role) public view returns (bool) {
		return roles[role].preapprovals[msg.sender];
	}

	function isPreapprovedForRole(bytes32 role, address account) public view returns (bool) {
		return roles[role].preapprovals[account];
	}

	function isSufficientlyStakedForRole(bytes32 role) public view returns (bool) {
		return stakedBalances[msg.sender] >= roles[role].requiredStake;
	}

	function isSufficientlyStakedForRole(bytes32 role, address account) public view returns (bool) {
		return stakedBalances[account] >= roles[role].requiredStake;
	}

	function grantRolePreapproval(bytes32 role, address account) public {
		require(hasRole(getRoleAdmin(role), msg.sender), "AccessControl: sender must be an admin to grant");
		roles[role].preapprovals[account] = true;
		grantRoleIfPreapprovedAndSufficientlyStaked(role, account);
	}

	function revokeRolePreapproval(bytes32 role, address account) public {
		require(hasRole(getRoleAdmin(role), msg.sender), "AccessControl: sender must be an admin to grant");
		delete roles[role].preapprovals[account];
		revokeRole(role, account);
	}

	function grantRoleIfPreapprovedAndSufficientlyStaked(bytes32 role, address account) private {
		if(isPreapprovedForRole(role, account) && isSufficientlyStakedForRole(role, account)) {
      roles[role].members[account] = true;
      emit RoleGranted(role, account, msg.sender);
		}
	}

	function revokeRoleIfNoLongerSufficientlyStaked(bytes32 role, address account) private {
		if(!isSufficientlyStakedForRole(role, account)) {
      delete roles[role].members[account];
      emit RoleRevoked(role, account, msg.sender);
		}
	}


	function nextSeq() private returns (uint256) {
		uint256 curProofSeq = proofSeq.current();
		proofSeq.increment();
		return curProofSeq;
	}

	function proveBlockSpecimenProduced(uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash) public {
		require(hasRole(BLOCK_SPECIMEN_PRODUCER_ROLE, msg.sender));

		emit BlockSpecimenPublicationProofAppended(
			uint64(nextSeq()),
			msg.sender,
			chainID,
			chainHeightPos,
			chainHeightLen,
			specimenSize,
			specimenHash
		);
	}
}
