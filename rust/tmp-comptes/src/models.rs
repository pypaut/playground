use chrono::DateTime;
use chrono::Local;
use serde::Serialize;
use std::fmt;

#[derive(Debug)]
pub struct Budget {
    pub id: i32 ,
    pub label: String,
    pub amount: f32,
    pub date: DateTime<Local>,
    pub tag_id: i32,
}

#[derive(Serialize)]
pub struct Expense {
    pub id: i32,
    pub label: String,
    pub amount: f32,
    pub date: DateTime<Local>,
    pub budget_id: i32,
}

pub struct Income {
    pub id: i32,
    pub label: String,
    pub amount: f32,
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
