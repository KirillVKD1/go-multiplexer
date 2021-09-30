### RUN 
go run main.go  
 
### DOCKER 
sudo docker build -t go-multiplexer . 
sudo docker run -p 8000:8000 go-multiplexer

### ENDPOINT  
    POST http://localhost:8000/http://localhost:8000/
 
### REQUEST BODY EXAMPLE 

{
    "urls": [
        "http://www.mocky.io/v2/5d99a2893100005d0097d991",
        "http://www.mocky.io/v2/5d99a2893100005d0097d991",
        "http://www.mocky.io/v2/5d99a2893100005d0097d991",
        "http://www.mocky.io/v2/5d99a2c6310000550097d992"    ]
} 

### RESPONSE OK EXAMPLE 
 
 {
    "Responses": [
        {
            "Url": "http://www.mocky.io/v2/5d99a2893100005d0097d991",
            "Status": 200,
            "Body": "{\n    id : 231231\n}"
        },
        {
            "Url": "http://www.mocky.io/v2/5d99a2893100005d0097d991",
            "Status": 200,
            "Body": "{\n    id : 231231\n}"
        },
        {
            "Url": "http://www.mocky.io/v2/5d99a2c6310000550097d992",
            "Status": 200,
            "Body": "{\n    request_id : 55, \n    status : 'new'\n}"
        },
        {
            "Url": "http://www.mocky.io/v2/5d99a2893100005d0097d991",
            "Status": 200,
            "Body": "{\n    id : 231231\n}"
        }
    ]
} 

### RESPONSE FAIL EXAMPLE 
 
Url http://www.mocky.io/v2/5d9c2d4931000055002fc3e5 failed with status: 500 
Body {
    status : 'i am a bad request'
}

### FAKE URLS

Fake API response:
http://www.mocky.io/v2/5d99a2893100005d0097d991
http://www.mocky.io/v2/5d99a2c6310000550097d992
Fake API Failed Response (500 Error)
http://www.mocky.io/v2/5d9c2d4931000055002fc3e5
  
