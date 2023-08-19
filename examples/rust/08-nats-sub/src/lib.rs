use extism_pdk::*;

extern "C" {
    fn hostPrint(ptr: u64) -> u64;
}

pub fn print(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrint(memory_text.offset) };
}

#[plugin_fn]
pub fn message(input: String) -> FnResult<u64> {

    print("🦀 👋 message: ".to_string() + &input);
    
    Ok(0)
}
