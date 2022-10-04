use std::io;

fn main() {
    const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
    println!("The value of THREE_HOURS_IN_SECONDS is: {THREE_HOURS_IN_SECONDS}");

    // let x = 5; <- this doesn't work because it's not mutable
    let mut x = 5;
    println!("The value of x is: {x}");
    x = 6;
    println!("The value of x is: {x}");

    let guess = "42";
    {
        // this guess "shadows" the first guess, it is a number instead of a string
        let guess: u32 = "43".parse().expect("Not a number!");
        println!("The value of guess in the inner scope is: {guess}");
    }

    println!("The value of guess is: {guess}");

    let heart_eyed_cat: char = '😻';
    println!("emoji: {heart_eyed_cat}");

    let tupple: (i32, f64, u8) = (500, 6.4, 1);
    let five_hundred = tupple.0;
    let six_pt_four = tupple.1;
    let one = tupple.2;
    // could also do let (five_hundred, six_pt_four, one) = tupple
    println!("tupple: ({five_hundred}, {six_pt_four}, {one})");

    array_test()
}

fn array_test() {
    let a = [1, 2, 3, 4, 5];

    println!("Please enter an array index.");

    let mut index = String::new();

    io::stdin()
        .read_line(&mut index)
        .expect("Failed to read line");

    let index: usize = index
        .trim()
        .parse()
        .expect("Index entered was not a number");

    let element = a[index];

    println!("The value of the element at index {index} is: {element}");
}
