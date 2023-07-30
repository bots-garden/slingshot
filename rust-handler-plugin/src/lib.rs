#![no_main]

use extism_pdk::*;
use serde::Serialize;

#[derive(Serialize)]
struct Output {
    pub success: String,
    pub failure: String,
}

#[plugin_fn]
pub fn handle(input: String) -> FnResult<Json<Output>> {

    let msg: String = "🦀 Hello ".to_string() + &input;

    let output = Output { success: msg, failure: "no error".to_string() };
    
    Ok(Json(output))
}

