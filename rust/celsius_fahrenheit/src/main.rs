use std::io;

fn main() {
    let mut tmp_type: u8;

    loop {
        // Input temperature type
        println!("Please select the temperature unit :");
        println!("1 for Celsius->Fahreinheit");
        println!("2 for Fahreinheit->Celsius");
        let mut tmp_type_str = String::new();
        io::stdin()
            .read_line(&mut tmp_type_str)
            .expect("Failed to read line");

        tmp_type = match tmp_type_str.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Should be 1 or 2.");
                continue;
            }
        };

        if ![1, 2].contains(&tmp_type) {
            println!("Should be 1 or 2.");
            continue;
        }

        break;
    }

    loop {
        // Input temperature value
        let mut temp = String::new();
        println!("Please input your temperature.");
        io::stdin()
            .read_line(&mut temp)
            .expect("Failed to read line");

        let temp: f32 = match temp.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Should be a number!");
                continue;
            }
        };

        let new_temp = match tmp_type {
            1 => celsius_to_fahrenheit(temp),
            2 => fahrenheit_to_celsius(temp),
            0_u8 | 3_u8..=u8::MAX => continue,
        };

        println!("Result : {new_temp}");
        break;
    }
}

fn celsius_to_fahrenheit(tmp: f32) -> f32 {
    tmp * 9.0 / 5.0 + 32.0
}

fn fahrenheit_to_celsius(tmp: f32) -> f32 {
    (tmp - 32.0) * 5.0 / 9.0
}
