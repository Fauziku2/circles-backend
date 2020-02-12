# circles-backend

# Details
- Download Go
- open folder
- run 'go mod init github.com/Fauziku2/circles-backend'
- run "go run main.go"
- server will be running on localhost port: 8080


<h3>API List</h3> 

<strong>- returns all resumes</strong>
<p>GET /api/getResumes</p>

<strong>- return a resume by ID</strong>
<p>GET /api/getResumeById/{resume_id}</p>

<strong>- returns resumes whose names match the given name</strong>
<p>GET /api/getResumeByName/{name}</p>

<strong>- Upload a resume</strong>
<p>POST /api/uploadResumeDetails</p>
