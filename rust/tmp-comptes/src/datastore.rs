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
            budgets.push(Budget{
                id: row.get(0),
                label: row.get(1),
                amount: row.get(2),
                date: row.get(3),
                tag_id: row.get(5)
            })
        }
        
        return budgets;
    }
}