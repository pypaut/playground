mod models;

use crate::models::{Category, Expense};
use chrono::{DateTime, Local};
use std::io;
use std::io::Write;

/// Manage budget and expenses.
struct Cli {
    category_name: String,
    expense_name: String,
    expense_amount: f32,
    expense_comment: String,
    date: DateTime<Local>,
}

fn main() {
    // Hardcoded categories (for now)
    let categories: [Category; 2] = [
        Category {
            name: "Courses".to_string(),
            icon: "🛒".to_string(),
            description: "Nourriture et autres consommables".to_string(),
        },
        Category {
            name: "Factures".to_string(),
            icon: "🧾".to_string(),
            description: "Paiements récurrents".to_string(),
        },
    ];

    let mut expenses: Vec<Expense> = Vec::new();

    // Get category
    println!("Category:");
    for (i, c) in categories.iter().enumerate() {
        println!("{}: {} {} ({})", i, c.icon, c.name, c.description);
    }

    print!("> ");
    io::stdout().flush().unwrap();
    let mut category_index_buf = String::new();
    io::stdin()
        .read_line(&mut category_index_buf)
        .expect("Failed to read category index");
    category_index_buf = category_index_buf.trim().parse().unwrap();

    // Convert category number to u8
    let category_index: usize = match category_index_buf.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("Invalid category index!");
            return;
        }
    };

    // Retrieve category name
    let category_name = &(categories[category_index]).name;

    // Get expense label
    print!("Expense label: ");
    io::stdout().flush().unwrap();
    let mut expense_label = String::new();
    io::stdin()
        .read_line(&mut expense_label)
        .expect("Failed to read expense label");
    expense_label = expense_label.trim().parse().unwrap();
    if expense_label.is_empty() {
        println!("Label cannot be empty!");
        return;
    }

    // Get expense amount
    print!("Expense amount: ");
    io::stdout().flush().unwrap();
    let mut expense_amount_buf = String::new();
    io::stdin()
        .read_line(&mut expense_amount_buf)
        .expect("Failed to read expense amount");

    // Convert amount to f32
    let expense_amount = match expense_amount_buf.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("Invalid expense amount!");
            return;
        }
    };

    // Get expense comment (default: empty)
    print!("Expense comment (default: empty): ");
    io::stdout().flush().unwrap();
    let mut expense_comment = String::new();
    io::stdin()
        .read_line(&mut expense_comment)
        .expect("Failed to read expense comment");
    expense_comment = expense_comment.trim().parse().unwrap();

    // Get expense date (for now, defaults to "now")
    let expense_date = Local::now();

    let expense = Expense {
        label: expense_label,
        amount: expense_amount,
        comment: expense_comment,
        category_name: category_name.to_string(),
        date: expense_date,
    };
    expenses.push(expense);

    println!("New expense: {:?}", expenses.last().unwrap());
}
