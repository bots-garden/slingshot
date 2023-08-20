use extism_pdk::*;
use serde::{Serialize, Deserialize};
use thiserror::Error;

extern "C" {
    fn hostPrint(ptr: u64) -> u64;
}

pub fn print(text: String) {
    let mut memory_text: Memory = extism_pdk::Memory::new(text.len());
    memory_text.store(text);
    unsafe { hostPrint(memory_text.offset) };
}

extern "C" {
    fn hostGetEnv(ptr: u64) -> u64;
}

pub fn get_env(name: String) -> String {
    // copy the name of the environment variable to the shared memory
    let mut variable_name: Memory = extism_pdk::Memory::new(name.len());
    variable_name.store(name);

    // call the host function
    let offset: u64 = unsafe { hostGetEnv(variable_name.offset) };

    // read the value of the result from the shared memory
    let variable_value: Memory = extism_pdk::Memory::find(offset).unwrap();

    // return the value
    return variable_value.to_string().unwrap()
}


#[derive(Serialize, Deserialize, Debug)]
struct NatsConfig {
    pub id: String,
    pub url: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct NatsMessage {
    pub id: String,
    pub subject: String,
    pub data: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct StringResult {
    pub success: String,
    pub failure: String,
}

#[derive(Error, Debug)]
pub enum NatsError {
    #[error("Nats Connection issue")]
    ConnectionFailure,
    #[error("Store issue")]
    MessageFailure,
    #[error("Not found")]
    NotFound,
}

extern "C" {
    fn hostInitNatsConnection(offset: u64) -> u64;
}

pub fn init_nats_connection(nats_connection_id: String, nats_url: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the NATS connection",
	//    "url": "URL of the NATS server"
	// }
    let args = NatsConfig {
        id: nats_connection_id,
        url: nats_url,
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostInitNatsConnection(memory_json_str.offset) };

    // Get result from the shared memory
	// The host function (hostInitNatsConnection) returns a JSON buffer:
	// {
    //   "success": "id of the connection",
	//   "failure": "error message if error, else empty"
	// }
    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
    
    let json_string:String = memory_result.to_string().unwrap();
    let result: StringResult = serde_json::from_str(&json_string).unwrap();

    if result.failure.is_empty()  {
        return Ok(result.success);
    } else {
        return Err(NatsError::ConnectionFailure.into());
    }
}

extern "C" {
    fn hostNatsPublish(offset: u64) -> u64;
}

pub fn nats_publish(nats_connection_id: String, subject: String, data: String) -> Result<String, Error> {
	// Prepare the arguments for the host function
	// with a JSON string:
	// {
	//    "id": "id of the NATS client",
	//    "subject": "name",
	//    "data": "Bob Morane"
	// }
    let args = NatsMessage {
        id: nats_connection_id, 
        subject: subject,
        data: data,
    };
    let json_str: String = serde_json::to_string(&args).unwrap();

    // Copy the string value to the shared memory
    let mut memory_json_str: Memory = extism_pdk::Memory::new(json_str.len());
    memory_json_str.store(json_str);

    // Call host function with the offset of the arguments
    let offset: u64 = unsafe { hostNatsPublish(memory_json_str.offset) };

    // Get result from the shared memory
	// The host function returns a JSON buffer:
	// {
    //   "success": "OK",
	//   "failure": "error message if error, else empty"
	// }
    let memory_result: Memory = extism_pdk::Memory::find(offset).unwrap();
    
    let json_string:String = memory_result.to_string().unwrap();
    let result: StringResult = serde_json::from_str(&json_string).unwrap();

    if result.failure.is_empty()  {
        return Ok(result.success);
    } else {
        return Err(NatsError::MessageFailure.into());
    }


}

#[plugin_fn]
pub fn publish(input: String) -> FnResult<u64> {

    let nats_url : String = get_env("NATS_URL".to_string());
    let nats_connection : Result<String, Error> = init_nats_connection("natsconn01".to_string(), nats_url);

    match nats_connection {
        Ok(value) => print("🦀 nats connection: ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }

    match nats_publish("natsconn01".to_string(), "news".to_string(), input.to_string()) {
        Ok(value)  => print("🦀 🙂 ".to_string() + &value),
        Err(error) => print("😡 error: ".to_string() + &error.to_string()),
    }
    
    Ok(0)
}
