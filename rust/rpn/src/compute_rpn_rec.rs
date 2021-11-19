pub fn compute_rpn_rec(tokens: Vec<&str>) -> f32 {
    if tokens.len() == 1 {
        let res: f32 = tokens[0].parse().unwrap();
        return res;
    }

    if tokens.len() == 3 {
        let left: f32 = tokens[0].parse().unwrap();
        let right: f32 = tokens[1].parse().unwrap();
        match tokens[2] {
            "+" => return left + right,
            "-" => return left - right,
            "*" => return left * right,
            "/" => return left / right,
            _ => return 0.0,
        }
    }

    for i in 0..tokens.len() {
        if "+-*/".contains(tokens[i]) {
            let sub_tokens = vec![tokens[i - 2], tokens[i - 1], tokens[i]];
            let sub_res: f32 = compute_rpn_rec(sub_tokens);
            let sub_res_str = sub_res.to_string();
            let new_tokens = [
                tokens[0..(i - 2)].to_vec(),
                vec![&sub_res_str],
                tokens[(i + 1)..].to_vec(),
            ]
            .concat();

            return compute_rpn_rec(new_tokens);
        }
    }

    return 0.0
}

pub fn compute_rpn(input: &str) -> f32 {
    let split = input.split(" ");
    let tokens: Vec<&str> = split.collect();
    return compute_rpn_rec(tokens);
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
