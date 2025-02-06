// SPDX-License-Identifier: UNLICENSED 
pragma solidity ^0.8.12;

import {ServiceManagerBase, IAVSDirectory, IRewardsCoordinator, IServiceManager} from "eigenlayer-middleware/ServiceManagerBase.sol";
import {ISlashingRegistryCoordinator} from "eigenlayer-middleware/interfaces/ISlashingRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {IPermissionController} from "eigenlayer-core/contracts/interfaces/IPermissionController.sol";
import {IAllocationManager} from "eigenlayer-core/contracts/interfaces/IAllocationManager.sol";

contract MinimalServiceManager is ServiceManagerBase {
    constructor(
        IAVSDirectory __avsDirectory,
        IRewardsCoordinator __rewardsCoordinator,
        ISlashingRegistryCoordinator __slashingRegistryCoordinator,
        IStakeRegistry __stakeRegistry,
        IPermissionController __permissionController,
        IAllocationManager __allocationManager
    )
        ServiceManagerBase(
            __avsDirectory,
            __rewardsCoordinator,
            __slashingRegistryCoordinator,
            __stakeRegistry,
            __permissionController,
            __allocationManager
        )
    { }

    function initialize(
        address initialOwner,
        address _rewardsInitiator
    ) external initializer {
        __ServiceManagerBase_init(initialOwner, _rewardsInitiator);
    }
}
