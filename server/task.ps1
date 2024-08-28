[CmdletBinding()]
param (
  [Parameter(Mandatory=$true)]$action,
  [Parameter()]$postgres_uri="postgres://postgres:postgres@localhost:5432/chatty?sslmode=disable",
  [Parameter()]$port=6000
)
  
begin {
  $env:POSTGRES_URI=$postgres_uri
  $env:PORT=$port
  
}
  
process {
  if ($action -eq 'build') {
    Invoke-Expression "go build -o bin/chatty.exe ."
  }

  if ($action -eq 'run') {
    Invoke-Expression "./bin/chatty.exe"
  }
}
  
end {}
