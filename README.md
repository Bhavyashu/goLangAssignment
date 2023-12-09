# Array Sorting Application

This is a simple Go application for sorting arrays. It provides two endpoints for sorting arrays sequentially and concurrently.

## Endpoints

### Process Single (`/process-single`)

Sorts arrays sequentially.

- **Endpoint:** `/process-single`
- **Method:** `POST`
- **Request Payload:**
  ```json
  {
    "to_sort": [
      [9, 7, 5, 3, 1],
      [8, 6, 4, 2, 0],
      [10, 20, 15, 25, 5]
    ]
  }
  ```
- **Response:**
  ```json
  {
    "sorted_arrays": [
      [1, 3, 5, 7, 9],
      [0, 2, 4, 6, 8],
      [5, 10, 15, 20, 25]
    ],
    "time_ns": 7583
  }
  ```

### Process Concurrent (`/process-concurrent`)

Sorts arrays concurrently.

- **Endpoint:** `/process-concurrent`
- **Method:** `POST`
- **Request Payload:**
  ```json
  {
    "to_sort": [
      [9, 7, 5, 3, 1],
      [8, 6, 4, 2, 0],
      [10, 20, 15, 25, 5]
    ]
  }
  ```
- **Response:**
  ```json
  {
    "sorted_arrays": [
      [1, 3, 5, 7, 9],
      [0, 2, 4, 6, 8],
      [5, 10, 15, 20, 25]
    ],
    "time_ns": 26125
  }
  ```


## Testing

To test the application, you can use the provided `input.json` file with a large dataset. Use tools like Postman to make HTTP POST requests to the following routes:

- `https://arraysorter.onrender.com/process-concurrent` for conccurrent sorting
- `https://arraysorter.onrender.com/process-single` for sequential sorting
Or else you can acess the postman collection below and just hit the api endpoints to get the time taken, if tested locally there's a significant diffrence in repsonse times of single processing and concurrency processing
but the projected hosted on free tier on render.com doesn't show much of a difference I'd request you to clone the repository and run localy to see 80% + diffrence in efficiency of sinle and concurrent sorting.

- Make sure you select the environment in the postman before hitting the endpoints.
- [Link to postman collection](https://www.postman.com/orange-equinox-912249/workspace/array-sorter/request/24863361-5428d9a1-400a-468d-95c2-ad909a9dc072)
