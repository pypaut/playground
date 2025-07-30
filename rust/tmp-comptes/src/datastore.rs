use std::time::{SystemTime, UNIX_EPOCH};
use chrono::{DateTime, Local, TimeZone};
use postgres::{Client, NoTls};
use crate::models::Budget;

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
            let date_system: SystemTime = row.get(3);
            let date: DateTime<Local> = system_time_to_date_time(date_system);
            budgets.push(Budget{
                id: row.get(0),
                label: row.get(1),
                amount: row.get(2),
                date,
                tag_id: row.get(4)
            })
        }
        
        return budgets;
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