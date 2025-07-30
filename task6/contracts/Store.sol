// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 计数器合约
contract Store {
    uint256 private counter;
    
    event CounterIncreased(uint256 newValue);
    
    // 增加计数器（写操作）
    function increment() public returns (uint256) {
        counter++;
        emit CounterIncreased(counter);
        return counter;
    }
    
    // 查询计数器（读操作）
    function getCounter() public view returns (uint256) {
        return counter;
    }
}