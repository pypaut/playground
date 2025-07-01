use chrono::DateTime;
use chrono::Local;
use serde::Serialize;
use std::fmt;

#[derive(Serialize)]
pub struct Expense {
    pub label: String,
    pub amount: f32,
    pub comment: String,
    pub category_name: String,
    pub date: DateTime<Local>,
}

pub(crate) struct Category {
    pub name: String,
    pub icon: String,
    pub description: String,
}

// Pretty print
impl fmt::Debug for Expense {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", serde_json::to_string_pretty(&self).unwrap())
    }
}
