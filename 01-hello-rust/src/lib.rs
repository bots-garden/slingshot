use std::collections::HashMap;

use extism_pdk::*;
use json::Value;
use serde::{Serialize, Deserialize};

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
struct Response {
    pub text_body: String,
    pub json_body: Value,
    pub headers : HashMap<String,String>,
    pub status_code: i64
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
struct Request {
    pub body: String,
    pub base_url: String,
    pub headers : HashMap<String,String>,
    pub method: String
}


#[plugin_fn]
pub fn handle(http_request_data: String) -> FnResult<Json<Response>> {

    let request : Request = serde_json::from_str(&http_request_data).unwrap();

    let message: String = "🩵 Hello ".to_string() + &request.body;

    let mut headers: HashMap<String, String> = HashMap::new();
    headers.insert("Content-Type".to_string(), "text/plain; charset=utf-8".to_string());
    headers.insert("X-Slingshot-version".to_string(), "0.0.0".to_string());

    let response : Response = Response { 
        text_body: message, 
        json_body: serde_json::from_str("{}")?, 
        headers , 
        status_code: 200 
    };
    
    Ok(Json(response))
}

