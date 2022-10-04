fn main() {
    println!("Hello, world!");
    another_function(five());
    print_labeled_measurement(5, 'h');
    wild_let();
}

fn another_function(x: i32) {
    println!("The value of x is {x}");
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {value}{unit_label}");
}

fn wild_let() {
    let y = {
        let x = 3;
        x + 1
    }; // this block evaluates to 4

    println!("The value of y is: {y}");
}

fn five() -> i32 {
    5
}