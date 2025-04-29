use serde::{Deserialize, Serialize};

/// Types ported from Go to Rust

// Operator ID is a unique identifier for an operator
#[derive(Debug, Clone, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub struct OperatorId(pub String);

// Operator information
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct OperatorInfo {
    pub socket: String,
    // Other operator info fields would be added here
}
