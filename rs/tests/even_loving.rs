use anyhow::{anyhow, Result};
use teal::node::server::{Certifier, Config};

/// A certifier that only accepts even data (as big-endian bytes)
#[derive(Clone)]
pub struct EvenLovingCertifier;

impl Certifier for EvenLovingCertifier {
    fn get_response(&self, _config: &Config, data: &[u8]) -> Result<Vec<u8>> {
        // Check if last byte is even (interpreting data as big-endian integer)
        if let Some(&last) = data.last() {
            if last & 1 == 0 {
                return Ok(data.to_vec());
            }
        }
        Err(anyhow!("invalid data: not even"))
    }
}
