use extism_pdk::*;

extern "C" {
    fn hostPrintln(ptr: u64) -> u64;
}

pub fn println(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrintln(memory_text.offset) };
}

#[plugin_fn]
pub fn message(input: String) -> FnResult<u64> {

    println("🦀 👋 message: ".to_string() + &input);
    
    Ok(0)
}
