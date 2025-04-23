use serde::{Deserialize, Serialize};
use std::fmt;

/// Types ported from Go to Rust

// Operator ID is a unique identifier for an operator
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub struct OperatorId(pub String);

// TaskResponse is the response data from a task
#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct TaskResponse(pub Vec<u8>);

impl TaskResponse {
    pub fn new(data: Vec<u8>) -> Self {
        Self(data)
    }
}

// Operator information
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct OperatorInfo {
    pub socket: String,
    // Other operator info fields would be added here
}