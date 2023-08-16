use extism_pdk::*;

#[plugin_fn]
pub fn hello(input: String) -> FnResult<String> {

    let output : String = "👋 Hello ".to_string() + &input;
    
    Ok(output)
}
