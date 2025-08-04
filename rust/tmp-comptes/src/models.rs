use chrono::DateTime;
use chrono::Local;
use serde::Serialize;
use std::fmt;

#[derive(Debug)]
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
    pub date: DateTime<Local>,
    pub budget: String,
}

#[derive(Debug)]
pub struct Income {
    pub label: String,
    pub amount: f32,
    pub date: DateTime<Local>,
}

#[derive(Debug)]
pub(crate) struct Tag {
    pub label: String,
    pub icon: String,
    pub description: String,
}

// Pretty print
impl fmt::Debug for Expense {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", serde_json::to_string_pretty(&self).unwrap())
    }
}
