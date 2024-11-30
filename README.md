# Calculator API

This API allows users to perform basic arithmetic operations like addition, subtraction, multiplication, and division, as well as summing an array of numbers.

## API Specification

### Base URL

### Endpoints

- **Endpoint:** `POST /add`  
  **Summary:** Adds two numbers.  
  **Request Body:**
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
  **Request Body:**
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
