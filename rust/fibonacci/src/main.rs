fn fibonacci(n: u64) -> u64 {
    if n == 0 {
        return 0;
    } else if n == 1 {
        return 1;
    }

    fibonacci(n-2) + fibonacci(n-1)
}

fn main() {
    println!("fibonacci(0): {}", fibonacci(0));
    println!("fibonacci(1): {}", fibonacci(1));
    println!("fibonacci(10): {}", fibonacci(10));
    println!("fibonacci(20): {}", fibonacci(20));
}
