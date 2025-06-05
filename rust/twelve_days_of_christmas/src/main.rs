const DAYS_COUNT: [&str;12] = [
    "first",
    "second",
    "third",
    "fourth",
    "fifth",
    "sixth",
    "seventh",
    "eighth",
    "ninth",
    "tenth",
    "eleventh",
    "twelfth",
];

const LINES: [&str;12] = [
    "Twelve drummers drumming,",
    "Eleven pipers piping,",
    "Ten lords a-leaping,",
    "Nine ladies dancing,",
    "Eight maids a-milking,",
    "Seven swans a-swimming,",
    "Six geese a-laying,",
    "Five gold rings,",
    "Four calling birds,",
    "Three French hens,",
    "Two turtle doves,",
    "And a partridge in a pear tree!",
];

fn main() {
    for i in 0..12 {
        println!("On the {} day of Christmas,", DAYS_COUNT[i]);
        println!("my true love sent to me");
        if i == 0 {
            println!("A partridge in a pear tree.");
        } else {
            for j in (11-i)..12 {
                println!("{}", LINES[j]);
            }
        }

        println!("");
    }
}
