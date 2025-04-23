# 📚 UxEdu: Scalable API for Modern Education Systems
UxEdu is a modular, distributed API designed to support a wide range of educational platforms. It provides ready-to-use templates, yet remains fully customizable to meet specific institutional or personal needs.

Built with gRPC and HTTP gateways <br/>
Designed for stability, scalability, and clarity <br/>
Easily integrable into existing systems 



# 📑 Table of Contents

- [📚 UxEdu: Scalable API for Modern Education Systems](#-uxedu-scalable-api-for-modern-education-systems)
   - [🏳️ How to run the project](#️-how-to-run-the-project)
   - [🔧 Services Overview](#-services-overview)
   - [🧩 What Each Service Does](#-what-each-service-does)
   - [⚙️ Tech-Stack](#️-tech-stack)
   - [👍Our advantages](#our-advantages)
- [🚀 Main API Gateway](#rocketmain-api-gateway)
   - [📁 File Service API](#-file-service-api)
      - [`GetFile`](#-getfile)
      - [`SetFile`](#-setfile)
      - [🧪 Example Usage (HTTP)](#-example-usage-http)
   - [📓 Text Service](#-text-service)
      - [1. Class Endpoints](#1-class-endpoints)
      - [2. Subject Endpoints](#2-subject-endpoints)
      - [3. Section Endpoints](#3-section-endpoints)
      - [4. Lesson Endpoints](#4-lesson-endpoints)
- [👤 User Service](#user-service)
   - [`GET /api/v0/ping`](#get-apiv0ping)
   - [`POST /api/v0/register`](#post-apiv0register)
   - [`POST /api/v0/login`](#post-apiv0login)
- [🎥 Video API Gateway](#video-api-gateway)
   - [🎥 Video Service](#video-service)
      - [`POST /video/v0/verify`](#post-videov0verify)
      - [`POST /video/v0`](#post-videov0)
      - [`GET /video/v0`](#get-videov0)



## 🏳️ How to run the project

I think you all have the Docker engine -> 
```shell
docker-compose up -d --build
```



### 🔧 Services Overview
UxEdu is composed of 4 gRPC services and 2 RESTful HTTP gateways, organized as follows:

```
                                     file-service [1]
main-api-gateway = '/api/v0' ──────> text-service [2]
                                     auth-service [3]
```

```
video-api-gateway = '/video/v0' ───> video-service [4]
```

### 🧩 What Each Service Does

```
Service	Responsibilities
File Service [1]	Upload and retrieve non-video files
Text Service [2]	Manage content: classes, tasks, subjects; handle comments and likes
Auth Service [3]	User authentication: registration and login
Video Service [4]	Stream or upload/download videos using chunks (supports both gRPC and HTTP)
```

## ⚙️ Tech-Stack

1. *Database* <br/>
   The main database of UxEdu is **PostgreSQL**. We're chosen postgres because of it speed, popularity and simplicity
2. *Cashing* <br/>
   We're using **Redis** to cash the items because of its fantastic speed, RAM-type data saving and also simplicity
3. *Object storage* <br/>
   Minio - our object storage. I have chosen it because of it popularity, low requirements and a large amount of documentation



## 👍Our advantages

1. Beautiful logging
2. Reasonable distribution of services, as well as their division into two gateways
3. Following the principles of clean architecture (CA)
4. Convenient commands in Makefile



# 🚀Main api gateway

The base path of main-api-gateway is
```
/api/v0/...
```


## 📁 File Service API


## Overview

The `File` service provides a simple interface for uploading and downloading files to and from MinIO
It exposes two HTTP endpoints under the `/api/v0/file` path: one for uploading (`POST`) and one for downloading (`GET`) files.


---


### ▸ `GetFile`

**Description:**  
Downloads a file from the specified bucket using the object key.

**HTTP Mapping:**  
GET /api/v0/file



#### Request: `GetFileRequest`

| Field         | Type   | Description                              |
|---------------|--------|------------------------------------------|
| `bucket_name` | string | Name of the bucket to fetch the file from |
| `object_key`  | string | Unique identifier (key) of the file      |

#### Response: `GetFileResponse`

| Field           | Type   | Description                                      |
|-----------------|--------|--------------------------------------------------|
| `content`       | bytes  | Binary content of the file                       |
| `filename`      | string | Original filename                                |
| `content_type`  | string | MIME type (e.g., `image/png`, `application/pdf`) |
| `size`          | int64  | Size of the file in bytes                        |
| `last_modified` | string | ISO 8601 timestamp of last modification          |

---

### ▸ `SetFile`

**Description:**  
Uploads a file to the specified bucket with the given object name.

**HTTP Mapping:**  
POST /api/v0/file

#### Request: `SetFileRequest`

| Field         | Type   | Description                                  |
|---------------|--------|----------------------------------------------|
| `bucket_name` | string | Name of the target bucket                    |
| `object_name` | string | Identifier to assign to the uploaded object  |
| `object`      | bytes  | Binary content of the file to upload         |

#### Response: `SetFileResponse`

- Empty on success. Use HTTP status codes to determine operation result.

## 🧪 Example Usage (HTTP)

### Upload File (POST)

```bash
curl -X POST http://localhost:8080/api/v0/file \
  -H "Content-Type: application/json" \
  --data-binary '{"bucket_name":"my-bucket", "object_name":"report.pdf", "object":"<base64-encoded-bytes>"}'
```

### Download File (GET)

```shell
curl "http://localhost:8080/api/v0/file?bucket_name=my-bucket&object_key=report.pdf"
```



<br/>
<br/>


## 📓Text Service

## Overview

This document provides a detailed and structured reference for the **TextService gRPC API**. The service is centered around educational content, organized hierarchically as **Class → Subject → Section → Lesson**

---

## 1. Class Endpoints

### Create Class

- **POST** `/class`

- **Request:** `CreateClassRequest`

- **Response:** `CreateClassResponse`


### Get Class

- **GET** `/class`

- **Request:** `GetClassRequest`

- **Response:** `GetClassResponse`


### Get All Classes

- **GET** `/classes`

- **Request:** `GetClassesRequest`

- **Response:** `GetClassesResponse`


### Add Subject to Class

- **POST** `/class-subject`

- **Request:** `AddSubjectInClassRequest`

- **Response:** `AddSubjectInClassResponse`


### Remove Subject from Class

- **DELETE** `/class-subject`

- **Request:** `RemoveSubjectFromClassRequest`

- **Response:** `RemoveSubjectFromClassResponse`


### Delete Class

- **DELETE** `/class`

- **Request:** `DeleteClassRequest`

- **Response:** `DeleteClassResponse`


---

## 2. Subject Endpoints

### Create Subject

- **POST** `/subject`

- **Request:** `CreateSubjectRequest`

- **Response:** `CreateSubjectResponse`


### Get Subject

- **GET** `/subject`

- **Request:** `GetSubjectRequest`

- **Response:** `GetSubjectResponse`


### Get All Subjects

- **GET** `/subjects`

- **Request:** `GetSubjectsRequest`

- **Response:** `GetSubjectsResponse`


### Add Section to Subject

- **POST** `/section-subject`

- **Request:** `AddSectionInSubjectRequest`

- **Response:** `AddSectionInSubjectResponse`


### Remove Section from Subject

- **DELETE** `/section-subject`

- **Request:** `RemoveSectionFromSubjectRequest`

- **Response:** `RemoveSectionFromSubjectResponse`


### Delete Subject

- **DELETE** `/subject`

- **Request:** `DeleteSubjectRequest`

- **Response:** `DeleteSubjectResponse`


---

## 3. Section Endpoints

### Create Section

- **POST** `/section`

- **Request:** `CreateSectionRequest`

- **Response:** `CreateSectionResponse`


### Get Section

- **GET** `/section`

- **Request:** `GetSectionRequest`

- **Response:** `GetSectionResponse`


### Get All Sections

- **GET** `/sections`

- **Request:** `GetSectionsRequest`

- **Response:** `GetSectionsResponse`


### Add Lesson to Section

- **POST** `/section-lesson`

- **Request:** `AddLessonInSectionRequest`

- **Response:** `AddLessonInSectionResponse`


### Remove Lesson from Section

- **DELETE** `/section-lesson`

- **Request:** `RemoveLessonFromSectionRequest`

- **Response:** `RemoveLessonFromSectionResponse`


### Delete Section

- **DELETE** `/section`

- **Request:** `DeleteSectionRequest`

- **Response:** `DeleteSectionResponse`


---

## 4. Lesson Endpoints

### Create Lesson

- **POST** `/lesson`

- **Request:** `CreateLessonRequest`

- **Response:** `CreateLessonResponse`


### Get Lesson

- **GET** `/lesson`

- **Request:** `GetLessonRequest`

- **Response:** `GetLessonResponse`


### Get All Lessons

- **GET** `/lessons`

- **Request:** `GetLessonsRequest`

- **Response:** `GetLessonsResponse`


### Add Media and Metadata to Lesson

#### Add Video

- **POST** `/lesson-video`

- **Request:** `AddVideoInLessonRequest`

- **Response:** `AddVideoInLessonResponse`


#### Remove Video

- **DELETE** `/lesson-video`

- **Request:** `RemoveVideoFromLessonRequest`

- **Response:** `RemoveVideoFromLessonResponse`


#### Add File

- **POST** `/lesson-file`

- **Request:** `AddFileInLessonRequest`

- **Response:** `AddFileInLessonResponse`


#### Remove File

- **DELETE** `/lesson-file`

- **Request:** `RemoveFileFromLessonRequest`

- **Response:** `RemoveFileFromLessonResponse`


#### Add Exercise

- **POST** `/lesson-exercise`

- **Request:** `AddExerciseInLessonRequest`

- **Response:** `AddExerciseInLessonResponse`


#### Remove Exercise

- **DELETE** `/lesson-exercise`

- **Request:** `RemoveExerciseFromLessonRequest`

- **Response:** `RemoveExerciseFromLessonResponse`


#### Add Comment

- **POST** `/lesson-comment`

- **Request:** `AddCommentInLessonRequest`

- **Response:** `AddCommentInLessonResponse`


#### Remove Comment

- **DELETE** `/lesson-comment`

- **Request:** `RemoveCommentFromLessonRequest`

- **Response:** `RemoveCommentFromLessonResponse`


#### Increase Rating

- **POST** `/lesson-rating`

- **Request:** `IncreaseRatingRequest`

- **Response:** `IncreaseRatingResponse`


#### Decrease Rating

- **DELETE** `/lesson-rating`

- **Request:** `DecreaseRatingRequest`

- **Response:** `DecreaseRatingResponse`


### Delete Lesson

- **DELETE** `/lesson`

- **Request:** `DeleteLessonRequest`

- **Response:** `DeleteLessonResponse`


---


<br/>
<br/>
<br/>

# 👤User Service

## General Information

The User Service is responsible for user registration and authentication. It supports a REST gateway using `google.api.http`.

---

## Endpoints

### `GET /api/v0/ping`

**Method:** `User.Get`

Checks the service's availability. Returns a message.

**Request:**

```proto
message Request {
    string message = 1;
}
```

**Response:**

```proto
message Reply {
    string message = 1;
}
```

---

### `POST /api/v0/register`

**Method:** `User.Register`

Registers a new user. Accepts all fields in the request body. Returns a UUID, token, and an `isAdmin` flag.

**Request:**

```proto
message RegisterRequest {
    string username = 1;
    string email = 2;
    string password = 3;
    string passwordConfirm = 4;
}
```

**Response:**

```proto
message RegisterResponse {
    string uuid = 1; // Unique user identifier
    string token = 2; // JWT token
    bool isAdmin = 3; // Admin flag
}
```

---

### `POST /api/v0/login`

**Method:** `User.Login`

Authenticates a user. Accepts email and password. Returns a UUID and token.

**Request:**

```proto
message LoginRequest {
    string email = 1;
    string password = 2;
}
```

**Response:**

```proto
message LoginResponse {
    string uuid = 1; // Unique user identifier
    string token = 2; // JWT token
}
```

---


# Video Api Gateway
The base endpoint prefix is
```
/video/v0
```


<br/><br/><br/>

## 🎥Video Service

## General Information

The Video Service manages video upload and retrieval using streams and chunks. It supports REST endpoints via `google.api.http` annotations and provides both streamed and standard HTTP interactions.

---

## Endpoints

### `POST /video/v0/verify`

**Method:** `Video.SetVideoChunk`

Initializes a video upload session by providing the video name. Returns the assigned video ID.

**Request:**

```proto
message SetVideoChunkRequest {
  string video_name = 1;
}
```

**Response:**

```proto
message SetVideoChunkResponse {
  string video_id = 1;
}
```

---

### `POST /video/v0`

**Method:** `Video.AddToVideoChunk`

Adds a new chunk of data to an existing video. The video is identified by name.

**Request:**

```proto
message AddToVideoChunkRequest {
  string video_name = 1;
  bytes chunk_data = 2;
}
```

**Response:**

```proto
message AddToVideoChunkResponse {
  string video_name = 1;
}
```

---

### `GET /video/v0`

**Method:** `Video.GetVideoChunk`

Retrieves video data chunk by video ID.

**Request:**

```proto
message GetVideoChunkRequest {
  string video_id = 1;
}
```

**Response:**

```proto
message GetVideoChunkResponse {
  bytes chunk_data = 1;
}
```

---

## Streaming Methods

### `Video.SetVideoStream`

Uploads a video stream directly.

**Request (stream):**

```proto
message SetVideoStreamRequest {
  string video_name = 1;
  bytes chunk_data = 2;
}
```

**Response:**

```proto
message SetVideoStreamResponse {
  string video_id = 1;
}
```

---

### `Video.GetVideoStream`

Retrieves a video as a stream of chunks.

**Request:**

```proto
message GetVideoStreamRequest {
  string video_id = 1;
}
```

**Response (stream):**

```proto
message GetVideoStreamResponse {
  bytes chunk_data = 1;
}
```