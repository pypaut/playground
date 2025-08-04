use std::io::Error;
use std::time::{SystemTime, UNIX_EPOCH};
use chrono::{DateTime, Local, TimeZone};
use postgres::{Client, NoTls};
use crate::models::{Budget, Expense, Income, Tag};

pub struct Datastore {
    pub client: Client,
}

pub fn new_datastore() -> Datastore {
    let client = Client::connect(
        "postgresql://user:pass@localhost:9001/db",
        NoTls).unwrap();
    return Datastore { client };
}

impl Datastore {
    pub fn list_budgets(&mut self) -> Vec<Budget> {
        let mut budgets: Vec<Budget> = Vec::new();
        for row in self.client.query("SELECT * FROM budgets", &[]).unwrap() {
            let date_system: SystemTime = row.get(2);
            let date: DateTime<Local> = system_time_to_date_time(date_system);
            budgets.push(Budget{
                label: row.get(0),
                amount: row.get(1),
                date,
                tag: row.get(3)
            })
        }
        
        budgets
    }

    pub fn add_budget(&mut self, budget: Budget) {
        self.client.query(
            "INSERT INTO budgets (label, amount, date, tag_id) VALUES ($1, $2, $3, $4)",
            &[&budget.label, &budget.amount, &budget.date.to_string(), &budget.tag],
        ).unwrap();
    }

    pub fn list_expenses(&mut self) -> Vec<Expense> {
        let mut expenses: Vec<Expense> = Vec::new();
        for row in self.client.query("SELECT * FROM expenses", &[]).unwrap() {
            let date_system: SystemTime = row.get(2);
            let date: DateTime<Local> = system_time_to_date_time(date_system);
            expenses.push(Expense{
                label: row.get(0),
                amount: row.get(1),
                date,
                budget: row.get(3)
            })
        }

        expenses
    }
    
    pub fn list_incomes(&mut self) -> Vec<Income> {
        let mut incomes: Vec<Income> = Vec::new();
        for row in self.client.query("SELECT * FROM incomes", &[]).unwrap() {
            let date_system: SystemTime = row.get(2);
            let date: DateTime<Local> = system_time_to_date_time(date_system);
            incomes.push(Income{
                label: row.get(0),
                amount: row.get(1),
                date,
            })
        }

        incomes
    }

    pub fn list_tags(&mut self) -> Vec<Tag> {
        let mut tags: Vec<Tag> = Vec::new();
        for row in self.client.query("SELECT * FROM tags", &[]).unwrap() {
            tags.push(Tag{
                label: row.get(0),
                description: row.get(1),
                icon: row.get(2),
            })
        }

        tags
    }
}

fn system_time_to_date_time(t: SystemTime) -> DateTime<Local> {
    // Source : https://users.rust-lang.org/t/convert-std-time-systemtime-to-chrono-datetime-datetime/7684/4
    let (sec, nsec) = match t.duration_since(UNIX_EPOCH) {
        Ok(dur) => (dur.as_secs() as i64, dur.subsec_nanos()),
        Err(e) => { // unlikely but should be handled
            let dur = e.duration();
            let (sec, nsec) = (dur.as_secs() as i64, dur.subsec_nanos());
            if nsec == 0 {
                (-sec, 0)
            } else {
                (-sec - 1, 1_000_000_000 - nsec)
            }
        },
    };
    
    Local.timestamp_opt(sec, nsec).unwrap()
}