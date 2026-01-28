# ApiClient

A reusable JavaScript API client for making HTTP requests using the Fetch API. This class simplifies sending GET, POST, PUT, and DELETE requests with customizable headers and a base URL.

## Features

- Supports **GET**, **POST**, **PUT**, and **DELETE** methods.
- Configurable base URL and default headers.
- Automatic JSON parsing of responses.
- Error handling with descriptive messages.

## Installation

No external dependencies are required. Simply include the `ApiClient` class in your project.

```bash
# If using ES Modules, you can import it directly:
import { ApiClient } from './ApiClient.js';
```

## Usage

### Initialize the Client

```javascript
const api = new ApiClient("https://api.example.com", {
  Authorization: "Bearer YOUR_TOKEN",
});
```

### GET Request

```javascript
try {
  const data = await api.get("users");
  console.log(data);
} catch (error) {
  console.error(error);
}
```

### POST Request

```javascript
try {
  const data = await api.post("users", {
    name: "John Doe",
    email: "john@example.com",
  });
  console.log(data);
} catch (error) {
  console.error(error);
}
```

### PUT Request

```javascript
try {
  const data = await api.put("users/1", { name: "Jane Doe" });
  console.log(data);
} catch (error) {
  console.error(error);
}
```

### DELETE Request

```javascript
try {
  const data = await api.delete("users/1");
  console.log(data);
} catch (error) {
  console.error(error);
}
```

## Customization

- **baseURL**: Set the root URL for all API calls.
- **defaultHeaders**: Provide headers like `Authorization` or custom content types.

## Error Handling

If the response is not OK, the client throws an error with the status code and message.
