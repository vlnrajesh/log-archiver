# log-archiver
[![GoDoc](https://godoc.org/github.com/vlnrajesh/log-archiver?status.svg)](http://godoc.org/github.com/vlnrajesh/log-archiver) [![CircleCI](https://circleci.com/gh/vlnrajesh/log-archiver.svg?style=svg)](https://circleci.com/gh/vlnrajesh/log-archiver) [![Go Report Card](https://goreportcard.com/badge/github.com/vlnrajesh/log-archiver)](https://goreportcard.com/report/github.com/vlnrajesh/log-archiver)

Log archiver for cloud platforms
## Problem
Organization rely on logs for understanding many business critical functions, application behaviour and user activities such as, operations
user performed, interpretability of application. These logs provide vital information about the application, its status on what the system is doing
and what has happened. These logs are often collected into files on a local disk with log rotation options imposed at individual application logs.
Logs also became important for network and system operations. As system truncate logs on periodic basis to have resources available for
applications, we require an additional storage location to preserve these logs for offline analysis and record of evidence.




Package:
1. set_logger: This method is used to facilitate console and file log handlers
2. file_matrix: Read Configuration file for important sections and finalize initial file matrix
3. apply_size_filter: Apply size filters to given file. Size of file calculated only in Bytes not in bits
4. apply_age_filter: Apply age filters to given file. Age of file calculated in seconds
5. apply_filter: Read each element of initial file list and apply size and age filters. If filters were success upload files to S3.
6. s3_upload_file: Given file will be uploaded to S3 as a multipart object. This method adds hostname plus current date of upload
        as additional directory path
