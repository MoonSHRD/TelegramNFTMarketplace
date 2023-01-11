pragma solidity ^0.8.0;

// SPDX-License-Identifier: UNLICENSED


/**
 *  @author Jack Bekket
 *  ALL RIGHTS RESERVED
 */
library FeesCalculator {

    /**
    *   Calculate fee (UnSafeMath) -- use it only if it ^0.8.0
    *   @dev calculate fee
    *   @param amount number from whom we take fee
    *   @param scale scale for rounding. 100 is 1/100 (percent). we can encreace scale if we want better division (like we need to take 0.5% instead of 5%, then scale = 1000)
    */
    function calculateAbstractFee(uint256 amount, uint256 scale, uint256 promille_fee_) public pure returns(uint256) {
        uint a = amount / scale;
        uint b = amount % scale;
        uint c = promille_fee_ / scale;
        uint d = promille_fee_ % scale;
        return a * c * scale + a * d + b * c + (b * d + scale - 1) / scale;
    }
}