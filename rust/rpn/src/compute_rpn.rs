pub fn compute_rpn(input: &str) -> f32 {
    let split = input.split(" ");
    let tokens: Vec<&str> = split.collect();

    if tokens.len() == 1 {
        let res: f32 = tokens[0].parse().unwrap();
        return res;
    }

    let mut stack: Vec<f32> = Vec::new();
    for token in tokens {
        if "+-*/".contains(token) {
            let right: f32 = stack.pop().unwrap();
            let left: f32 = stack.pop().unwrap();

            match token {
                "+" => stack.push(left + right),
                "-" => stack.push(left - right),
                "*" => stack.push(left * right),
                "/" => stack.push(left / right),
                _ => return 0.0,
            }
        } else {
            let number: f32 = token.parse().unwrap();
            stack.push(number);
        }
    }

    return stack[0];
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_simple_eval() {
        assert_eq!(compute_rpn("1"), 1.0);
        assert_eq!(compute_rpn("0"), 0.0);
        assert_eq!(compute_rpn("42"), 42.0);
    }

    #[test]
    fn test_add() {
        assert_eq!(compute_rpn("1 1 +"), 2.0);
        assert_eq!(compute_rpn("0 0 +"), 0.0);
        assert_eq!(compute_rpn("42.0 523 +"), 565.0);
        assert_eq!(compute_rpn("42 -523.0 +"), -481.0);
    }

    #[test]
    fn test_sub() {
        assert_eq!(compute_rpn("1 1 -"), 0.0);
        assert_eq!(compute_rpn("0.0 0 -"), 0.0);
        assert_eq!(compute_rpn("42 523 -"), -481.0);
        assert_eq!(compute_rpn("16 4 -"), 12.0);
        assert_eq!(compute_rpn("522 523 -"), -1.0);
    }

    #[test]
    fn test_mul() {
        assert_eq!(compute_rpn("1 1 *"), 1.0);
        assert_eq!(compute_rpn("0.0 0 *"), 0.0);
        assert_eq!(compute_rpn("42 0 *"), 0.0);
        assert_eq!(compute_rpn("16 4 *"), 64.0);
        assert_eq!(compute_rpn("522 2 *"), 1044.0);
    }

    #[test]
    fn test_div() {
        assert_eq!(compute_rpn("1 1 /"), 1.0);
        assert_eq!(compute_rpn("0.0 231 /"), 0.0);
        assert_eq!(compute_rpn("42 2 /"), 21.0);
        assert_eq!(compute_rpn("16 4 /"), 4.0);
        assert_eq!(compute_rpn("522 522 /"), 1.0);
    }

    #[test]
    fn test_advanced_1() {
        assert_eq!(compute_rpn("3 4 + 2 -"), 5.0);
    }

    #[test]
    fn test_advanced_2() {
        assert_eq!(compute_rpn("3 5 8 * 7 + *"), 141.0)
    }
}
