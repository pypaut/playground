mod models;
mod datastore;

use crate::models::{Budget, Expense, Tag};
use chrono::Local;
use std::io;
use std::io::Write;
use crate::datastore::Datastore;

/// Manage budget and expenses.
///
/// TODO
/// - "interactive" mode vs cli mode (current way would be interactive)
/// - CRUD on each struct
/// - display current month data
/// - add income type
/// - add tests
/// - better selection method (menu selectable with arrows)
/// - use postgres instead of hardcoded slices

fn main() {
    let ds: Datastore = datastore::new_datastore();
    // Hardcoded values (for now)
    // let tags: [Tag; 4] = [
    //     Tag {
    //         label: "Factures".to_string(),
    //         icon: "🧾".to_string(),
    //         description: "Paiements récurrents, charges fixes, abonnements".to_string(),
    //         color: "".to_string(),
    //     },
    //     Tag {
    //         label: "Épargnes".to_string(),
    //         icon: "💰".to_string(),
    //         description: "On met de côté".to_string(),
    //         color: "".to_string(),
    //     },
    //     Tag {
    //         label: "Dépenses courantes".to_string(),
    //         icon: "💳".to_string(),
    //         description: "".to_string(),
    //         color: "".to_string(),
    //     },
    //     Tag {
    //         label: "Dépenses variables".to_string(),
    //         icon: "💶".to_string(),
    //         description: "Dépenses autres".to_string(),
    //         color: "".to_string(),
    //     },
    // ];

    // let budgets: [Budget; 3] = [
    //     Budget {
    //         label: "Courses".to_string(),
    //         amount: 450.,
    //         date: Local::now(),
    //         tag: "Dépenses courantes".to_string(),
    //     },
    //     Budget {
    //         label: "Épargne chats".to_string(),
    //         amount: 0.0,
    //         date: Local::now(),
    //         tag: "Épargnes".to_string(),
    //     },
    //     Budget {
    //         label: "Cadeau".to_string(),
    //         amount: 0.0,
    //         date: Local::now(),
    //         tag: "Dépenses variables".to_string(),
    //     },
    // ];

    // let actions: [String; 1] = ["New expense".to_string()];

    // let mut expenses: Vec<Expense> = Vec::new();

    // // Get action to do
    // let action_index = get_action(&actions);
    // println!();
    // println!("--> {}", actions[action_index]);

    // // Run action
    // match action_index {
    //     0 => {
    //         let expense = new_expense(&budgets);
    //         expenses.push(expense);
    //         println!("New expense: {:?}", expenses.last().unwrap());
    //     }
    //     1_usize.. => {
    //         println!("uh what?");
    //         return;
    //     }
    // };
}

// fn g// et_action(actions: &[String; 1]) -> usize {
    // loop {
    //     println!();
    //     println!("What do you want to do?");
    //     for (i, a) in actions.iter().enumerate() {
    //         println!("{}: {}", i, a);
    //     }
    //     print!("> ");
    //     io::stdout().flush().unwrap();
    //     let mut action_index_buf = String::new();
    //     io::stdin()
    //         .read_line(&mut action_index_buf)
    //         .expect("Failed to read action");

    //     // Convert category number to u8
    //     let action_index: usize = match action_index_buf.trim().parse() {
    //         Ok(num) => num,
    //         Err(_) => {
    //             println!("Invalid action index!");
    //             continue;
    //         }
    //     };

    //     if action_index >= actions.len() {
    //         println!("Invalid action index!");
    //         continue;
    //     }

    //     return action_index;
    // }
// }

// fn n// ew_expense(budgets: &[Budget; 3]) -> Expense {
    // loop {
    //     // Select budget to which attach this expense
    //     println!();
    //     println!("Select the budget");
    //     for (i, a) in budgets.iter().enumerate() {
    //         println!("{}: {}", i, a.label);
    //     }
    //     println!();
    //     print!("> ");
    //     io::stdout().flush().unwrap();
    //     let mut budget_buf = String::new();
    //     io::stdin()
    //         .read_line(&mut budget_buf)
    //         .expect("Failed to read action");
    //     let budget: String = budget_buf.trim().parse().unwrap();

    //     // Get expense label
    //     print!("Expense label (default: {}): ", budget);
    //     io::stdout().flush().unwrap();
    //     let mut expense_label = String::new();
    //     io::stdin()
    //         .read_line(&mut expense_label)
    //         .expect("Failed to read expense label");
    //     expense_label = expense_label.trim().parse().unwrap();
    //     if expense_label.is_empty() {
    //         println!("Label cannot be empty!");
    //         continue;
    //     }

    //     if expense_label == "" {
    //         // For single-expense budgets
    //         expense_label = budget.clone();
    //     }

    //     // Get expense amount
    //     print!("Expense amount: ");
    //     io::stdout().flush().unwrap();
    //     let mut expense_amount_buf = String::new();
    //     io::stdin()
    //         .read_line(&mut expense_amount_buf)
    //         .expect("Failed to read expense amount");

    //     // Convert amount to f32
    //     let expense_amount: f32 = match expense_amount_buf.trim().parse() {
    //         Ok(num) => num,
    //         Err(_) => {
    //             println!("Invalid expense amount!");
    //             continue;
    //         }
    //     };

    //     return Expense {
    //         label: expense_label,
    //         amount: expense_amount,
    //         budget,
    //         date: Local::now(),
    //     };
    // }

// }
