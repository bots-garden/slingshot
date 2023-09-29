use extism_pdk::*;

extern "C" {
    fn hostPrintln(ptr: u64) -> u64;
}

pub fn println(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrintln(memory_text.offset) };
}

extern "C" {
    fn hostLog(ptr: u64) -> u64;
}

pub fn log(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostLog(memory_text.offset) };
}

#[plugin_fn]
pub fn hello(input: String) -> FnResult<u64> {

    println("🦀 hello world 🌍 ".to_string() + &input);
    log("🙂 have a nice day 🏖️".to_string());
    
    Ok(0)
}
