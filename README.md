# doodocs-rest

## How to run

Use the link: https://doodocs-rest.onrender.com

Or...

1. Clone repo:
```shell
git clone https://github.com/itelman/doodocs-rest
```

2. Open repo:
```shell
cd ./doodocs-rest
```

3. Run the repo:
```shell
go run .
```

## Endpoints

### 1. Listing all files in zip archive

- POST /api/archive/information
- Request:

```json
{
    "file_name": string,
    "archive_size": int,
	"total_size": int,
	"total_files": int,
	"files" : {
        "file_path": string,
	    "size": 3,
        "mime_type": string,
    }
}
```

- Response (200 OK):

```json
{
    "file_name": string,
    "archive_size": int,
    "total_size": int,
    "total_files": int,
    "files": [
        {
            "file_path": string,
            "size": int,
            "mime_type": string
        },
        {
            "file_path": string,
            "size": int,
            "mime_type": string
        },
        ...
        ...
        ...
    ]
}
```

### 2. Creating zip archive

- POST /api/archive/files
- Request:

```json
{
    "file_name": string,
    "archive_size": int,
	"total_size": int,
	"total_files": int,
	"files" : {
        "file_path": string,
	    "size": 3,
        "mime_type": string,
    }
}
```

- Response (200 OK):

```json
{
    "message": "OK"
}
```

### 3. Sending files by email

- POST /api/mail/file
- Request:

```json
{
   "file": string,
   "emails": string,
}
```

- Response (200 OK):

```json
{
    "message": "Email sent successfully"
}
```