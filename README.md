# Calculator API

This API allows users to perform basic arithmetic operations like addition, subtraction, multiplication, and division, as well as summing an array of numbers.

## API Specification

### Base URL

### Endpoints

- **Endpoint:** `POST /add`  
  **Summary:** Adds two numbers.  
  Request Body:
  ```json
  {
    "num1": 10.5,
    "num2": 5
  }
  ```
  Result:
  ```json
  {
      "result": 15.5
  }
  ```

- **Endpoint:** `POST /subtract`  
  **Summary:** subtract two numbers.  
  Request Body:
  ```json
  {
    "num1": 10.5,
    "num2": 5
  }
  ```
  Result:
  ```json
  {
      "result": 5.5
  }
  ```

- **Endpoint:** `POST /multiply`  
  **Summary:** multiply two numbers.  
  Request Body:
  ```json
  {
    "num1": 10.5,
    "num2": 5
  }
  ```
  Result:
  ```json
  {
      "result": 52.5
  }
  ```

- **Endpoint:** `POST /divide`  
  **Summary:** divide two numbers.  
  Request Body:
  ```json
  {
    "num1": 10.5,
    "num2": 5
  }
  ```
  Result:
  ```json
  {
      "result": 2.1
  }
  ```
- **Endpoint:** `POST /sum`  
  **Summary:** sum of array of numbers.  
  Request Body:
  ```json
  {
    "nums": [10.13, 24.4, 19.23]
  }
  ```
  Result:
  ```json
  {
      "result": 53.76
  }
  ```
