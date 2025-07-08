use chrono::DateTime;
use chrono::Local;
use serde::Serialize;
use std::fmt;

pub struct Budget {
    pub label: String,
    pub amount: f32,
    pub date: DateTime<Local>,
    pub tag: String,
}

#[derive(Serialize)]
pub struct Expense {
    pub label: String,
    pub amount: f32,
    pub budget: String,
    pub date: DateTime<Local>,
}

pub(crate) struct Tag {
    pub label: String,
    pub icon: String,
    pub description: String,
    pub color: String,
}

// Pretty print
impl fmt::Debug for Expense {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", serde_json::to_string_pretty(&self).unwrap())
    }
}
